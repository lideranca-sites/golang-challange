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

type TestSuiteProduct struct {
	suite.Suite
	app        *fiber.App
	db         *gorm.DB
	connection *sql.DB
	user       *models.User
	product    *models.Product
	token      string
}

func (suite *TestSuiteProduct) SetupTest() {
	var err error

	suite.db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
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

	suite.product = &models.Product{
		Name:     "Product 1",
		Price:    1000,
		Quantity: 10,
		UserID:   suite.user.ID,
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

	token, err := features.CreateJwtToken(features.CreateJwtTokenDTO{
		UserId: suite.user.ID,
	})
	assert.NoError(suite.T(), err)
	suite.token = token
}

func (suite *TestSuiteProduct) TearDownTest() {
	suite.db.Migrator().DropTable(&models.User{}, &models.Product{})
	suite.connection.Close()
}

func (suite *TestSuiteProduct) TestCreateProduct() {
	new_product := &models.Product{
		Name:     "Product 2",
		Price:    1000,
		Quantity: 10,
	}
	body, err := json.Marshal(new_product)
	assert.NoError(suite.T(), err)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewReader(body))
	assert.NoError(suite.T(), err)
	req.Header.Add("Authorization", "Bearer "+suite.token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := suite.app.Test(req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusCreated, resp.StatusCode)
}

func (suite *TestSuiteProduct) TestGetProducts() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/products", nil)
	assert.NoError(suite.T(), err)
	resp, err := suite.app.Test(req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)
}

func (suite *TestSuiteProduct) TestUpdateProduct() {
	new_product := &models.Product{
		Name:     suite.product.Name,
		Price:    2000,
		Quantity: 20,
	}
	body, err := json.Marshal(new_product)
	assert.NoError(suite.T(), err)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/products/1", bytes.NewReader(body))
	assert.NoError(suite.T(), err)
	req.Header.Add("Authorization", "Bearer "+suite.token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := suite.app.Test(req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)
}

func (suite *TestSuiteProduct) TestDeleteProduct() {
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/products/1", nil)
	assert.NoError(suite.T(), err)
	req.Header.Add("Authorization", "Bearer "+suite.token)
	resp, err := suite.app.Test(req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusNoContent, resp.StatusCode)
}

func TestProduct(t *testing.T) {
	suite.Run(t, new(TestSuiteProduct))
}
