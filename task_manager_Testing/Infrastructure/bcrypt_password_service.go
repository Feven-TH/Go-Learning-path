package infrastructure

import (
    "golang.org/x/crypto/bcrypt"
)

type BcryptPasswordService struct{}

func NewBcryptPasswordService() *BcryptPasswordService {
    return &BcryptPasswordService{}
}

func (b *BcryptPasswordService) IHashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func (b *BcryptPasswordService) IComparePassword(hashed, plain string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}
