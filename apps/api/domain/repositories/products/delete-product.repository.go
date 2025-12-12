package domain_repositories

type DeleteProductRepository interface {
	DeleteProductById(id uint) error
}
