package services

import (
	"errors"
	"example/apps/api/domain/dto"
	"example/apps/api/domain/entities"
	api_errors "example/apps/api/domain/errors"
	"example/apps/api/domain/ports"
	domain_repositories "example/apps/api/domain/repositories/users"
	"fmt"
)

type userRepository interface {
	domain_repositories.CreateUserRepository
	domain_repositories.FindUserRepository
}

type SignUserServices struct {
	repository userRepository
	crypto     ports.CryptoPort
	token      ports.SignTokenPort
}

func NewSignUserServices(repo userRepository, crypto ports.CryptoPort, token ports.SignTokenPort) SignUserServices {
	return SignUserServices{repository: repo, crypto: crypto, token: token}
}

func (su SignUserServices) SignIn(input dto.UserDTO) (string, api_errors.ApiErrorPort) {
	user, err := su.repository.FindOneByEmail(input.Email)
	if err != nil {
		return api_errors.ErrorHandler("", err)
	}

	comparison := su.crypto.IsSameContent(user.Password, input.Password)
	if !comparison {
		return api_errors.ErrorHandler("", errors.New("invalid credentials"))
	}

	return api_errors.ErrorHandler(su.token.GenerateToken(user.ID))
}

func (su SignUserServices) SignUp(input dto.UserDTO) (string, api_errors.ApiErrorPort) {
	passwordHash, err := su.crypto.GenerateHash(input.Password)
	if err != nil {
		fmt.Println(err.Error())
		return api_errors.ErrorHandler("", err)
	}

	input.Password = string(passwordHash)
	inputEntity := entities.UserEntityFromDto(input)
	user, err := su.repository.CreateUser(inputEntity)
	if err != nil {
		fmt.Println(err.Error())
		return api_errors.ErrorHandler("", err)
	}

	return api_errors.ErrorHandler(su.token.GenerateToken(user.ID))
}

func (su SignUserServices) FindMe(id uint) (dto.UserDTO, api_errors.ApiErrorPort) {
	user, err := su.repository.FindByID(id)
	return api_errors.ErrorHandler(user.ToDto(), err)
}
