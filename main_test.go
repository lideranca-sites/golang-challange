package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"testing"

	"example/apps/api/infra/setup"
	"example/apps/api/modules/user/features"
	"example/libs/database"
	"example/libs/database/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TestSuite struct {
	suite.Suite

	app *fiber.App

	db         *gorm.DB
	connection *sql.DB
	user       *models.User
	product    *models.Product
}

func (suite *TestSuite) SetupSuite() {
	var err error

	suite.app = setup.Setup()

	suite.db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	database.DB = suite.db
	assert.NoError(suite.T(), err)

	suite.connection, err = suite.db.DB()
	assert.NoError(suite.T(), err)

	suite.db.AutoMigrate(&models.User{}, &models.Product{})

	suite.user = &models.User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@doe.com",
	}

	suite.product = &models.Product{
		Name:   "Product 1",
		UserID: suite.user.ID,
	}
}

func (suite *TestSuite) TestShouldCreateUser() {
	data := &features.CreateUserBodyDTO{
		Name:  &suite.user.Name,
		Email: &suite.user.Email,
	}

	payload, err := json.Marshal(data)

	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(payload))

	assert.NoError(suite.T(), err)

	req.Header.Set("Content-Type", "application/json")

	res, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), http.StatusCreated, res.StatusCode)
}

func (suite *TestSuite) TestShouldNotCreate() {
	data := &features.CreateUserBodyDTO{
		Name: &suite.user.Name,
	}

	payload, err := json.Marshal(data)

	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(payload))

	assert.NoError(suite.T(), err)

	req.Header.Set("Content-Type", "application/json")

	res, err := suite.app.Test(req)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), http.StatusBadRequest, res.StatusCode)

	var response map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&response)

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), "The Email field is required", response["message"])

}

func (suite *TestSuite) TearDownSuite() {
	suite.db.Migrator().DropTable(&models.User{}, &models.Product{})
	suite.connection.Close()
}

func TestSomething(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
