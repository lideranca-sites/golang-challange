package models

import (
	"example/apps/api/domain/dto"
	"example/apps/api/domain/entities"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model `gorm:"embedded"`
	ID         uint      `gorm:"primaryKey"`
	Name       string    `gorm:"type:varchar(100);not null"`
	Email      string    `gorm:"type:varchar(100);not null;uniqueIndex:unique_email"`
	Password   string    `gorm:"type:varchar(256);not null"`
	Products   []Product `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime:true"`
}

func (u User) ToEntity() entities.UserEntity {
	products := make([]entities.ProductEntity, len(u.Products))
	for i, p := range u.Products {
		products[i] = p.ToEntity()
	}

	return entities.UserEntity{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt.Time,
		Products:  products,
	}
}

func UserModelFromDto(input dto.UserDTO) User {
	products := make([]Product, len(input.Products))
	for i, p := range input.Products {
		products[i] = ProductModelFromDto(p)
	}

	return User{
		ID:       input.ID,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Products: products,
	}
}
