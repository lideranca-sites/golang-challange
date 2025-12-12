package utils

import "golang.org/x/crypto/bcrypt"

type CryptoUtils struct{}

func NewCryptoUtils() CryptoUtils {
	return CryptoUtils{}
}

func (this *CryptoUtils) GenerateHash(input string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
}

func (this *CryptoUtils) IsSameContent(hashedInput string, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedInput), []byte(input))
	return err == nil
}
