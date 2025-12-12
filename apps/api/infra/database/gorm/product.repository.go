package gorm_repositories

import (
	"errors"
	"example/apps/api/domain/dto"
	"example/apps/api/domain/entities"
	api_errors "example/apps/api/domain/errors"
	"example/libs/database"
	"example/libs/database/models"

	"gorm.io/gorm"
)

type ProductGormRepository struct {
	client *gorm.DB
}

func NewProductGormRepository() (ProductGormRepository, error) {
	client := database.DB
	if client == nil {
		return ProductGormRepository{}, errors.New("failed to connect database")
	}

	return ProductGormRepository{client: client}, nil
}

func (pr ProductGormRepository) CreateProduct(input dto.ProductDTO) (entities.ProductEntity, error) {
	model := models.ProductModelFromDto(input)
	result := pr.client.Create(&model)

	if result.RowsAffected < 1 {
		return entities.ProductEntity{}, api_errors.New("Failed to create new product", 400)
	}

	return model.ToEntity(), nil
}

func (pr ProductGormRepository) GetAll() ([]entities.ProductEntity, error) {
	var products []models.Product
	result := pr.client.Find(&products)
	if result.RowsAffected < 1 {
		return []entities.ProductEntity{}, api_errors.New("No products found", 404)
	}
	entities := transformModelsToEntity(products)

	return entities, nil
}

func (pr ProductGormRepository) GetByUserId(id uint) ([]entities.ProductEntity, error) {
	var products []models.Product
	result := pr.client.Where("user_id = ?", id).Find(&products)
	if result.RowsAffected < 1 {
		return []entities.ProductEntity{}, api_errors.New("No products found for this user", 404)
	}
	entities := transformModelsToEntity(products)

	return entities, nil
}

func (pr ProductGormRepository) SaveProduct(input entities.ProductEntity) (entities.ProductEntity, error) {
	model := models.ProductModelFromDto(input.ToDto())
	result := pr.client.Save(&model)
	if result.RowsAffected < 1 {
		return entities.ProductEntity{}, result.Error
	}
	return model.ToEntity(), nil
}

func (pr ProductGormRepository) DeleteProductById(id uint) error {
	result := pr.client.Delete(&entities.ProductEntity{}, id)
	if result.RowsAffected < 1 {
		return result.Error
	}

	return nil
}

func transformModelsToEntity(a []models.Product) []entities.ProductEntity {
	entities := make([]entities.ProductEntity, len(a))
	for i, p := range a {
		entities[i] = p.ToEntity()
	}

	return entities
}
