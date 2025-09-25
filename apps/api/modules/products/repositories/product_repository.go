package repositories

import "example/libs/database/models"

type ProductRepository interface {
	Create(product *models.Product) (*models.Product, error)

	FindAll() ([]models.Product, error)

	FindByUserID(userID int) ([]models.Product, error)

	FindByID(id int) (*models.Product, error)

	Update(product *models.Product) (*models.Product, error)

	Delete(id int) error
}
