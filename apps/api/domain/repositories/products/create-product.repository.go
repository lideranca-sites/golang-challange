package domain_repositories

import (
	"example/apps/api/domain/dto"
	"example/apps/api/domain/entities"
)

type CreateProductRepository interface {
	CreateProduct(input dto.ProductDTO) (entities.ProductEntity, error)
}
