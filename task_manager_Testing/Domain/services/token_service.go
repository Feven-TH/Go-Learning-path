package services

import (
	"task_manager_Testing/Domain/response"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	IGenerateAccessToken(userID , role string) (*response.TokenResponse, error)
	IVerifyToken(tokenString string) (jwt.MapClaims, error)
}
