package services

type PasswordService interface {
    IHashPassword(password string) (string, error)
    IComparePassword(hashed, plain string) error
}
