package domain_repositories

import (
	"example/apps/api/domain/entities"
)

type SaveUserRepository interface {
	SaveUser(input *entities.UserEntity) error
}
