package services

type PasswordService interface {
    HashPassword(password string) (string, error)
    ComparePassword(hashed, plain string) error
}
