package domain_repositories

import (
	"example/apps/api/domain/entities"
)

type CreateUserRepository interface {
	CreateUser(input entities.UserEntity) (entities.UserEntity, error)
}
