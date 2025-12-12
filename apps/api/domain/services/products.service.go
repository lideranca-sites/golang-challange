package services

import (
	"example/apps/api/domain/dto"
	"example/apps/api/domain/entities"
	api_errors "example/apps/api/domain/errors"
	"example/apps/api/domain/ports"
	domain_repositories "example/apps/api/domain/repositories/products"
	domain_user_repositories "example/apps/api/domain/repositories/users"
)

type productUnifiedRepositories interface {
	domain_repositories.CreateProductRepository
	domain_repositories.GetProductsRepository
	domain_repositories.SaveProductRepository
	domain_repositories.DeleteProductRepository
}

type userRepositories interface {
	domain_user_repositories.FindUserRepository
	domain_user_repositories.SaveUserRepository
}

type ProductsServices struct {
	repository     productUnifiedRepositories
	userRepository userRepositories
	logger         ports.LoggerPort
}

func NewProductsServices(repository productUnifiedRepositories, userRepository userRepositories, logger ports.LoggerPort) ProductsServices {
	return ProductsServices{repository: repository, userRepository: userRepository, logger: logger}
}

func (ps *ProductsServices) GetAllProducts() ([]dto.ProductDTO, api_errors.ApiErrorPort) {
	products, err := ps.repository.GetAll()
	dtos := transformEntitiesToDto(products)

	return api_errors.ErrorHandler(dtos, err)
}

func (ps *ProductsServices) GetProductsByUserID(userId uint) ([]dto.ProductDTO, api_errors.ApiErrorPort) {
	products, err := ps.repository.GetByUserId(userId)
	dtos := transformEntitiesToDto(products)

	return api_errors.ErrorHandler(dtos, err)
}

func (ps *ProductsServices) CreateProduct(input dto.ProductDTO) (dto.ProductDTO, api_errors.ApiErrorPort) {
	user, err := ps.userRepository.FindByID(input.UserID)
	if err != nil {
		ps.logger.Error(err.Error())
		return api_errors.ErrorHandler(dto.ProductDTO{}, error(api_errors.New("user not found", 404)))
	}
	newProduct := entities.ProductEntiryFromDto(input)
	err = user.AddProduct(newProduct)
	if err != nil {
		ps.logger.Error(err.Error())
		return api_errors.ErrorHandler(dto.ProductDTO{}, error(api_errors.New(err.Error(), 400)))
	}

	err = ps.userRepository.SaveUser(&user)

	index := len(user.Products) - 1
	newProduct = user.Products[index]
	return api_errors.ErrorHandler(newProduct.ToDto(), err)
}

func (ps *ProductsServices) SaveProduct(input dto.ProductDTO) (dto.ProductDTO, api_errors.ApiErrorPort) {
	inputEntity := entities.ProductEntiryFromDto(input)
	p, e := ps.repository.SaveProduct(inputEntity)

	return api_errors.ErrorHandler(p.ToDto(), e)
}

func (ps *ProductsServices) DeleteProduct(productId uint) (any, api_errors.ApiErrorPort) {
	e := ps.repository.DeleteProductById(productId)
	return api_errors.ErrorHandler("", e)
}

func transformEntitiesToDto(input []entities.ProductEntity) []dto.ProductDTO {
	result := make([]dto.ProductDTO, len(input))
	for i, p := range input {
		result[i] = p.ToDto()
	}

	return result
}
