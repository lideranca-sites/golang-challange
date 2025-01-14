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

type TestSuiteProduct struct {
	suite.Suite

	app *fiber.App

	db         *gorm.DB
	connection *sql.DB
	user       *models.User
	product    *models.Product
	token      string
}

func (suite *TestSuiteProduct) SetupTest() {
	var err error

	suite.app = server.Setup()

	suite.db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
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
		ID:       1,
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

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.Contains(suite.T(), response, "message")
	assert.Contains(suite.T(), response, "product")

	assert.Equal(suite.T(), "Product created successfully", response["message"])

	assert.Contains(suite.T(), response["product"], "id")
	assert.Contains(suite.T(), response["product"], "name")
	assert.Contains(suite.T(), response["product"], "price")
	assert.Contains(suite.T(), response["product"], "quantity")
	assert.Contains(suite.T(), response["product"], "user_id")

	assert.Equal(suite.T(), new_product.Name, response["product"].(map[string]interface{})["name"])
	assert.Equal(suite.T(), new_product.Price, int(response["product"].(map[string]interface{})["price"].(float64)))
	assert.Equal(suite.T(), new_product.Quantity, int(response["product"].(map[string]interface{})["quantity"].(float64)))
	assert.Equal(suite.T(), suite.user.ID, int(response["product"].(map[string]interface{})["user_id"].(float64)))
}

func (suite *TestSuiteProduct) TestGetProductsByUser() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/products?user_id=1", nil)

	assert.NoError(suite.T(), err)

	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.Contains(suite.T(), response, "products")

	products := response["products"].([]interface{})
	assert.NotEmpty(suite.T(), products)

	assert.Contains(suite.T(), products[0], "id")
	assert.Contains(suite.T(), products[0], "name")
	assert.Contains(suite.T(), products[0], "price")
	assert.Contains(suite.T(), products[0], "quantity")
	assert.Contains(suite.T(), products[0], "user_id")

	assert.Equal(suite.T(), suite.product.ID, int(products[0].(map[string]interface{})["id"].(float64)))
	assert.Equal(suite.T(), suite.product.Name, products[0].(map[string]interface{})["name"])
	assert.Equal(suite.T(), suite.product.Price, int(products[0].(map[string]interface{})["price"].(float64)))
	assert.Equal(suite.T(), suite.product.Quantity, int(products[0].(map[string]interface{})["quantity"].(float64)))
	assert.Equal(suite.T(), suite.product.UserID, int(products[0].(map[string]interface{})["user_id"].(float64)))
}

func (suite *TestSuiteProduct) TestGetProducts() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/products", nil)

	assert.NoError(suite.T(), err)

	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.Contains(suite.T(), response, "products")

	products := response["products"].([]interface{})
	assert.NotEmpty(suite.T(), products)

	assert.Contains(suite.T(), products[0], "id")
	assert.Contains(suite.T(), products[0], "name")
	assert.Contains(suite.T(), products[0], "price")
	assert.Contains(suite.T(), products[0], "quantity")
	assert.Contains(suite.T(), products[0], "user_id")

	assert.Equal(suite.T(), suite.product.ID, int(products[0].(map[string]interface{})["id"].(float64)))
	assert.Equal(suite.T(), suite.product.Name, products[0].(map[string]interface{})["name"])
	assert.Equal(suite.T(), suite.product.Price, int(products[0].(map[string]interface{})["price"].(float64)))
	assert.Equal(suite.T(), suite.product.Quantity, int(products[0].(map[string]interface{})["quantity"].(float64)))
	assert.Equal(suite.T(), suite.product.UserID, int(products[0].(map[string]interface{})["user_id"].(float64)))
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

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.Contains(suite.T(), response, "message")
	assert.Contains(suite.T(), response, "product")

	assert.Equal(suite.T(), "Product updated successfully", response["message"])
	
	assert.Contains(suite.T(), response["product"], "id")
	assert.Contains(suite.T(), response["product"], "name")
	assert.Contains(suite.T(), response["product"], "price")
	assert.Contains(suite.T(), response["product"], "quantity")

	assert.Equal(suite.T(), suite.product.ID, int(response["product"].(map[string]interface{})["id"].(float64)))
	assert.Equal(suite.T(), new_product.Name, response["product"].(map[string]interface{})["name"])
	assert.Equal(suite.T(), new_product.Price, int(response["product"].(map[string]interface{})["price"].(float64)))
	assert.Equal(suite.T(), new_product.Quantity, int(response["product"].(map[string]interface{})["quantity"].(float64)))
}

func (suite *TestSuiteProduct) TestDeleteProduct() {
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/products/1", nil)

	assert.NoError(suite.T(), err)

	req.Header.Add("Authorization", "Bearer "+suite.token)

	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), fiber.StatusNoContent, resp.StatusCode)
}

func (suite *TestSuiteProduct) TearDownSuite() {
	suite.db.Migrator().DropTable(&models.User{}, &models.Product{})
	suite.connection.Close()
}

func TestProduct(t *testing.T) {
	suite.Run(t, new(TestSuiteProduct))
}
