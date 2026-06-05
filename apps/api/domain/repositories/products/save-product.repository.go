package domain_repositories

import (
	"example/apps/api/domain/entities"
)

type SaveProductRepository interface {
	SaveProduct(input entities.ProductEntity) (entities.ProductEntity, error)
}
