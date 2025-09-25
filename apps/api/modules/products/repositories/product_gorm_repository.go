package repositories

import (
	"example/libs/database/models"

	"gorm.io/gorm"
)

type productGormRepository struct {
	db *gorm.DB
}

func NewProductGormRepository(db *gorm.DB) ProductRepository {
	return &productGormRepository{db: db}
}

func (r *productGormRepository) Create(product *models.Product) (*models.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productGormRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productGormRepository) FindByUserID(userID int) ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Where("user_id = ?", userID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productGormRepository) FindByID(id int) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productGormRepository) Update(product *models.Product) (*models.Product, error) {
	if err := r.db.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productGormRepository) Delete(id int) error {
	result := r.db.Delete(&models.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
