package domain_repositories

import (
	"example/apps/api/domain/entities"
)

type FindUserRepository interface {
	FindByID(id uint) (entities.UserEntity, error)
	FindOneByEmail(input string) (entities.UserEntity, error)
}
