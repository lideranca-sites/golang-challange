package mocks

import (
	"example/libs/database/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) Create(product *models.Product) (*models.Product, error) {
	args := m.Called(product)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *ProductRepositoryMock) FindAll() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *ProductRepositoryMock) FindByUserID(userID int) ([]models.Product, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *ProductRepositoryMock) FindByID(id int) (*models.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *ProductRepositoryMock) Update(product *models.Product) (*models.Product, error) {
	args := m.Called(product)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *ProductRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
