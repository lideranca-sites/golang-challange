package ports

type CryptoPort interface {
	GenerateHash(input string) ([]byte, error)
	IsSameContent(hashedInput string, input string) bool
}
