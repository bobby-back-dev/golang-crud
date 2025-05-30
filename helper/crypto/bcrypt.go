package crypto

import "golang.org/x/crypto/bcrypt"

type Hash struct {
}

func NewHash() *Hash {
	return &Hash{}
}

func (h *Hash) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (h *Hash) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
