package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"testing"

	"example/apps/api/infra/server"
	"example/apps/api/modules/auth/features"
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AuthTestSuite struct {
	suite.Suite

	app *fiber.App

	db         *gorm.DB
	connection *sql.DB
	user       *models.User
	product    *models.Product
}

func (suite *AuthTestSuite) SetupSuite() {
	var err error

	suite.app = server.Setup()

	suite.db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	database.DB = suite.db
	assert.NoError(suite.T(), err)

	suite.connection, err = suite.db.DB()
	assert.NoError(suite.T(), err)

	suite.db.AutoMigrate(&models.User{}, &models.Product{})

	suite.user = &models.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@doe.com",
		Password: "123456",
	}

	suite.product = &models.Product{
		ID:     1,
		Name:   "Product 1",
		UserID: suite.user.ID,
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(suite.user.Password), bcrypt.DefaultCost)
	assert.NoError(suite.T(), err)

	result := suite.db.Create(&models.User{
		Name:     suite.user.Name,
		Email:    suite.user.Email,
		Password: string(hash),
	})
	assert.NoError(suite.T(), result.Error)

	result = suite.db.Create(suite.product)
	assert.NoError(suite.T(), result.Error)
}

func (suite *AuthTestSuite) TestSignUp() {
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

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.Contains(suite.T(), response, "access_token")

	token := response["access_token"].(string)

	assert.NotEmpty(suite.T(), token)

}

func (suite *AuthTestSuite) TestSignIn() {
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

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.Contains(suite.T(), response, "access_token")

	token := response["access_token"].(string)

	assert.NotEmpty(suite.T(), token)
}

func (suite *AuthTestSuite) TestMe() {
	var token string

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

	assert.Equal(suite.T(), fiber.StatusOK, auth_resp.StatusCode)

	var auth_response map[string]interface{}

	err = json.NewDecoder(auth_resp.Body).Decode(&auth_response)

	assert.NoError(suite.T(), err)

	assert.Contains(suite.T(), auth_response, "access_token")

	token = auth_response["access_token"].(string)

	assert.NotEmpty(suite.T(), token)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/auth/me", nil)

	assert.NoError(suite.T(), err)

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.Contains(suite.T(), response, "user")

	user := response["user"].(map[string]interface{})

	assert.Contains(suite.T(), user, "id")
	assert.Contains(suite.T(), user, "name")
	assert.Contains(suite.T(), user, "email")

	assert.Equal(suite.T(), suite.user.ID, int(user["id"].(float64)))
	assert.Equal(suite.T(), suite.user.Name, user["name"])
	assert.Equal(suite.T(), suite.user.Email, user["email"])
}

func (suite *AuthTestSuite) TearDownSuite() {
	suite.db.Migrator().DropTable(&models.User{}, &models.Product{})
	suite.connection.Close()
}

func AuthTest(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
