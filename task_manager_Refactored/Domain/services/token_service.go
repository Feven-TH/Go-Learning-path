package services

import (
	"task_manager_Refactored/Domain/response"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateAccessToken(userID , role string) (*response.TokenResponse, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
}
