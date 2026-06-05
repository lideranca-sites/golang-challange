package domain_repositories

import (
	"example/apps/api/domain/entities"
)

type GetProductsRepository interface {
	GetAll() ([]entities.ProductEntity, error)
	GetByUserId(id uint) ([]entities.ProductEntity, error)
}
