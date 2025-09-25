package services

import (
	"errors"
	"example/apps/api/modules/products/repositories"
	"example/libs/database/models"
	"strconv"
)

type ProductService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(name string, price float64, quantity int, userID int) (*models.Product, error) {
	product := &models.Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
		UserID:   userID,
	}
	return s.repo.Create(product)
}

func (s *ProductService) GetProducts(userID string) ([]models.Product, error) {
	if userID != "" {
		id, err := strconv.Atoi(userID)
		if err != nil {
			return nil, errors.New("invalid user ID")
		}
		return s.repo.FindByUserID(id)
	}
	return s.repo.FindAll()
}

func (s *ProductService) UpdateProduct(id int, name string, price float64, quantity int) (*models.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	product.Name = name
	product.Price = price
	product.Quantity = quantity

	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}
