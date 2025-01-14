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

type ProductTestSuite struct {
	suite.Suite

	app *fiber.App

	db         *gorm.DB
	connection *sql.DB
	user       *models.User
	product    *models.Product
	token      string
}

func (suite *ProductTestSuite) ProductSetupTest() {
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
		Price:  1000,
		Quantity: 10,
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

	token, err := features.CreateJwtToken(features.CreateJwtTokenDTO{
		UserId: suite.user.ID,
	})

	assert.NoError(suite.T(), err)

	suite.token = token
}

func (suite *ProductTestSuite) TestCreateProduct() {
	new_product := &models.Product{
		Name:   "Product 2",
		Price:  1000,
		Quantity: 10,
		UserID: suite.user.ID,
	}

	body, err := json.Marshal(new_product)

	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/user/1/product", bytes.NewReader(body))

	assert.NoError(suite.T(), err)

	req.Header.Add("Authorization", "Bearer "+suite.token)

	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), fiber.StatusCreated, resp.StatusCode)

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.Contains(suite.T(), response, "id")
	assert.Contains(suite.T(), response, "name")
	assert.Contains(suite.T(), response, "price")
	assert.Contains(suite.T(), response, "quantity")
	assert.Contains(suite.T(), response, "user_id")

	assert.Equal(suite.T(), new_product.Name, response["name"])
	assert.Equal(suite.T(), new_product.Price, response["price"])
	assert.Equal(suite.T(), new_product.Quantity, response["quantity"])
	assert.Equal(suite.T(), new_product.UserID, int(response["user_id"].(float64)))
}

func (suite *ProductTestSuite) TestGetProduct() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/user/1/product/1", nil)

	assert.NoError(suite.T(), err)

	req.Header.Add("Authorization", "Bearer "+suite.token)

	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.Contains(suite.T(), response, "id")
	assert.Contains(suite.T(), response, "name")
	assert.Contains(suite.T(), response, "price")
	assert.Contains(suite.T(), response, "quantity")
	assert.Contains(suite.T(), response, "user_id")

	assert.Equal(suite.T(), suite.product.ID, int(response["id"].(float64)))
	assert.Equal(suite.T(), suite.product.Name, response["name"])
	assert.Equal(suite.T(), suite.product.Price, response["price"])
	assert.Equal(suite.T(), suite.product.Quantity, response["quantity"])
	assert.Equal(suite.T(), suite.product.UserID, int(response["user_id"].(float64)))
}

func (suite *ProductTestSuite) TestUpdateProduct() {
	new_product := &models.Product{
		ID:     1,
		Name:  suite.product.Name,
		Price: 2000,
		Quantity: 20,
	}

	body, err := json.Marshal(new_product)

	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/user/1/product/1", bytes.NewReader(body))

	assert.NoError(suite.T(), err)

	req.Header.Add("Authorization", "Bearer "+suite.token)

	req.Header.Set("Content-Type", "application/json")

	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.Contains(suite.T(), response, "id")
	assert.Equal(suite.T(), new_product.ID, int(response["id"].(float64)))
	assert.Contains(suite.T(), response, "name")
	assert.Equal(suite.T(), new_product.Name, response["name"])
	assert.Contains(suite.T(), response, "price")
	assert.Equal(suite.T(), new_product.Price, response["price"])
	assert.Contains(suite.T(), response, "quantity")
	assert.Equal(suite.T(), new_product.Quantity, response["quantity"])

	assert.Contains(suite.T(), response, "user_id")
	assert.Equal(suite.T(), suite.product.UserID, int(response["user_id"].(float64)))
}


func (suite *ProductTestSuite) TestDeleteProduct() {
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/user/1/product/1", nil)

	assert.NoError(suite.T(), err)

	req.Header.Add("Authorization", "Bearer "+suite.token)

	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), fiber.StatusNoContent, resp.StatusCode)

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.Empty(suite.T(), response)

	var product models.Product

	result := suite.db.First(&product, 1)

	assert.Error(suite.T(), result.Error)

	assert.Equal(suite.T(), gorm.ErrRecordNotFound, result.Error)

}

func (suite *ProductTestSuite) TestListProduct() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/user/1/product", nil)

	assert.NoError(suite.T(), err)

	req.Header.Add("Authorization", "Bearer "+suite.token)

	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)

	var response []map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.NotEmpty(suite.T(), response)

	assert.Contains(suite.T(), response[0], "id")
	assert.Contains(suite.T(), response[0], "name")
	assert.Contains(suite.T(), response[0], "price")
	assert.Contains(suite.T(), response[0], "quantity")
	assert.Contains(suite.T(), response[0], "user_id")

	assert.Equal(suite.T(), suite.product.ID, int(response[0]["id"].(float64)))
	assert.Equal(suite.T(), suite.product.Name, response[0]["name"])
	assert.Equal(suite.T(), suite.product.Price, response[0]["price"])
	assert.Equal(suite.T(), suite.product.Quantity, response[0]["quantity"])
	assert.Equal(suite.T(), suite.product.UserID, int(response[0]["user_id"].(float64)))
}

func (suite *ProductTestSuite) TestListAllProduct() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/product", nil)

	assert.NoError(suite.T(), err)

	resp, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)

	var response []map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.NotEmpty(suite.T(), response)

	assert.Contains(suite.T(), response[0], "id")
	assert.Contains(suite.T(), response[0], "name")
	assert.Contains(suite.T(), response[0], "price")
	assert.Contains(suite.T(), response[0], "quantity")
	assert.Contains(suite.T(), response[0], "user_id")

	assert.Equal(suite.T(), suite.product.ID, int(response[0]["id"].(float64)))
	assert.Equal(suite.T(), suite.product.Name, response[0]["name"])
	assert.Equal(suite.T(), suite.product.Price, response[0]["price"])
	assert.Equal(suite.T(), suite.product.Quantity, response[0]["quantity"])
	assert.Equal(suite.T(), suite.product.UserID, int(response[0]["user_id"].(float64)))
}


func (suite *ProductTestSuite) ProductTearDownSuite() {
	suite.db.Migrator().DropTable(&models.User{}, &models.Product{})
	suite.connection.Close()
}

func ProductTest(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}
