package infrastructure

import (
	"errors"
	// "os"
	"task_manager_Refactored/Domain/response"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenService struct {
    secret string
}

func NewJwtTokenService(secret string) *JwtTokenService {
    return &JwtTokenService{secret: secret}
}


func (j *JwtTokenService) GenerateAccessToken(userID, role string) (*response.TokenResponse, error) {
    claims := jwt.MapClaims{
        "sub":  userID,
        "role": role,
        "exp":  time.Now().Add(time.Hour * 1).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(j.secret))
    if err != nil {
        return nil,err
    }
    return &response.TokenResponse{
        AccessToken: signedToken,
    }, nil
}

func (j *JwtTokenService) VerifyToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(j.secret), nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, errors.New("invalid token")
}
