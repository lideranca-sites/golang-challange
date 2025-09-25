package services

import (
	"errors"
	"example/apps/api/modules/products/repositories/mocks"
	"example/libs/database/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductService_CreateProduct(t *testing.T) {
	mockRepo := new(mocks.ProductRepositoryMock)
	productService := NewProductService(mockRepo)

	productToCreate := &models.Product{
		Name:     "Test Product",
		Price:    10.0,
		Quantity: 100,
		UserID:   1,
	}

	mockRepo.On("Create", productToCreate).Return(productToCreate, nil)

	createdProduct, err := productService.CreateProduct("Test Product", 10.0, 100, 1)

	assert.NoError(t, err)
	assert.NotNil(t, createdProduct)
	assert.Equal(t, "Test Product", createdProduct.Name)
	mockRepo.AssertExpectations(t)
}

func TestProductService_CreateProduct_Error(t *testing.T) {
	mockRepo := new(mocks.ProductRepositoryMock)
	productService := NewProductService(mockRepo)

	productToCreate := &models.Product{
		Name:     "Test Product",
		Price:    10.0,
		Quantity: 100,
		UserID:   1,
	}

	expectedError := errors.New("database error")
	mockRepo.On("Create", productToCreate).Return(nil, expectedError)

	createdProduct, err := productService.CreateProduct("Test Product", 10.0, 100, 1)

	assert.Error(t, err)
	assert.Nil(t, createdProduct)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
