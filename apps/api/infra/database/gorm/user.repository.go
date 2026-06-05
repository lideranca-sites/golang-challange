package gorm_repositories

import (
	"errors"
	"example/apps/api/domain/entities"
	api_errors "example/apps/api/domain/errors"
	"example/libs/database"
	"example/libs/database/models"

	"gorm.io/gorm"
)

type UserGormRepository struct {
	client *gorm.DB
}

func NewUserGormRepository() (UserGormRepository, error) {
	client := database.DB
	if client == nil {
		return UserGormRepository{}, errors.New("failed to connecto to database")
	}

	return UserGormRepository{client: client}, nil
}

func (ur UserGormRepository) CreateUser(input entities.UserEntity) (entities.UserEntity, error) {
	user := models.UserModelFromDto(input.ToDto())
	result := ur.client.Create(&user)

	if result.RowsAffected < 1 {
		return entities.UserEntity{}, api_errors.New("failed to create new user", 400)
	}

	return user.ToEntity(), nil
}

func (ur UserGormRepository) FindByID(id uint) (entities.UserEntity, error) {
	var user models.User
	result := ur.client.Preload("Products").First(&user, id)

	if result.Error != nil {
		return entities.UserEntity{}, api_errors.New("Not found", 404)
	}

	return user.ToEntity(), nil
}

func (ur UserGormRepository) FindOneByEmail(email string) (entities.UserEntity, error) {
	var user models.User
	result := ur.client.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return entities.UserEntity{}, api_errors.New("User not found", 404)
	}

	return user.ToEntity(), nil
}

func (ur UserGormRepository) SaveUser(input *entities.UserEntity) error {
	inputModel := models.UserModelFromDto(input.ToDto())
	result := ur.client.Preload("Products").Save(inputModel)
	if result.RowsAffected < 1 {
		return api_errors.New(result.Error.Error(), 400)
	}
	*input = inputModel.ToEntity()
	return nil
}
