package infrastructure

import (
    "golang.org/x/crypto/bcrypt"
)

type BcryptPasswordService struct{}

func NewBcryptPasswordService() *BcryptPasswordService {
    return &BcryptPasswordService{}
}

func (b *BcryptPasswordService) HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func (b *BcryptPasswordService) ComparePassword(hashed, plain string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}
