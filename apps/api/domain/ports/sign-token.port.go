package ports

type SignTokenPort interface {
	GenerateToken(id uint) (string, error)
}
