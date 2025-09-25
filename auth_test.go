package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"example/apps/api/infra/server"
	"example/apps/api/modules/auth/features"
	"example/apps/api/modules/products/handlers"
	"example/apps/api/modules/products/repositories"
	"example/apps/api/modules/products/services"
	"example/libs/database/models"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TestSuiteAuth struct {
	suite.Suite
	app        *fiber.App
	db         *gorm.DB
	connection *sql.DB
	user       *models.User
}

func (suite *TestSuiteAuth) SetupSuite() {
	var err error

	suite.db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(suite.T(), err)

	suite.connection, err = suite.db.DB()
	assert.NoError(suite.T(), err)

	suite.db.AutoMigrate(&models.User{}, &models.Product{})

	productRepository := repositories.NewProductGormRepository(suite.db)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	suite.app = server.Setup(suite.db, productHandler)

	suite.user = &models.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@doe.com",
		Password: "123456",
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(suite.user.Password), bcrypt.DefaultCost)
	assert.NoError(suite.T(), err)

	result := suite.db.Create(&models.User{
		Name:     suite.user.Name,
		Email:    suite.user.Email,
		Password: string(hash),
	})
	assert.NoError(suite.T(), result.Error)
}

func (suite *TestSuiteAuth) TearDownSuite() {
	suite.db.Migrator().DropTable(&models.User{}, &models.Product{})
	suite.connection.Close()
}

func (suite *TestSuiteAuth) TestSignUp() {
	new_user := &models.User{
		Name:     "Jane Doe",
		Email:    "jane@doe.com",
		Password: "123456",
	}
	body := &features.SignUpBodyDTO{
		Name:     &new_user.Name,
		Email:    &new_user.Email,
		Password: &new_user.Password,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(suite.T(), err)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/sign-up", bytes.NewReader(bodyBytes))
	assert.NoError(suite.T(), err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := suite.app.Test(req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusCreated, resp.StatusCode)
}

func (suite *TestSuiteAuth) TestSignIn() {
	body := &features.SignInBodyDTO{
		Email:    &suite.user.Email,
		Password: &suite.user.Password,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(suite.T(), err)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/sign-in", bytes.NewReader(bodyBytes))
	assert.NoError(suite.T(), err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := suite.app.Test(req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)
}

func (suite *TestSuiteAuth) TestMe() {
	body := &features.SignInBodyDTO{
		Email:    &suite.user.Email,
		Password: &suite.user.Password,
	}
	bodyBytes, err := json.Marshal(body)
	assert.NoError(suite.T(), err)
	auth_req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/sign-in", bytes.NewReader(bodyBytes))
	assert.NoError(suite.T(), err)
	auth_req.Header.Set("Content-Type", "application/json")
	auth_resp, err := suite.app.Test(auth_req)
	assert.NoError(suite.T(), err)
	var auth_response map[string]interface{}
	err = json.NewDecoder(auth_resp.Body).Decode(&auth_response)
	assert.NoError(suite.T(), err)
	token := auth_response["access_token"].(string)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/auth/me", nil)
	assert.NoError(suite.T(), err)
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := suite.app.Test(req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)
}

func TestAuth(t *testing.T) {
	suite.Run(t, new(TestSuiteAuth))
}
