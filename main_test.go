package main

import (
	"bytes"
	"crud/modules/database"
	"crud/modules/database/models"
	"crud/modules/user"
	"crud/modules/user/features"
	"database/sql"
	"encoding/json"
	"net/http"
	"testing"

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

	suite.app = fiber.New()
	user.SetupRoutes(suite.app)

	suite.db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	database.DB = suite.db
	assert.NoError(suite.T(), err)

	suite.connection, err = suite.db.DB()
	assert.NoError(suite.T(), err)

	suite.db.AutoMigrate(&models.User{}, &models.Product{})

	suite.user = &models.User{
		ID: 1, 
		Name: "John Doe", 
		Email: "john@doe.com",
	}

	suite.product = &models.Product{
		Name:   "Product 1",
		UserID: suite.user.ID,
	}
}

func (suite *TestSuite) TestShouldCreateUser() {
	data := &features.CreateUserDTO{
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
	data := &features.CreateUserDTO{
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
}

func (suite *TestSuite) TearDownSuite() {
	suite.db.Migrator().DropTable(&models.User{}, &models.Product{})
	suite.connection.Close()
}

func TestSomething(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
