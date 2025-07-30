package middleware_test

import (
    "net/http"
    "net/http/httptest"
    "task_manager_Testing/Domain/services"
    "task_manager_Testing/Infrastructure/middleware"
    "task_manager_Testing/mocks"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "github.com/stretchr/testify/assert"
)

func setupRouter(tokenService services.TokenService, requiredRole string) *gin.Engine {
    gin.SetMode(gin.TestMode)
    r := gin.New()

    r.Use(middleware.AuthMiddleware(tokenService))

    if requiredRole != "" {
        r.Use(middleware.RoleRequired(requiredRole))
    }

    r.GET("/secure", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
    })

    return r
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
    mockTokenSvc := new(mocks.TokenService)
    mockClaims := jwt.MapClaims{
        "sub":  "user123",
        "role": "admin",
    }

    mockTokenSvc.On("IVerifyToken", "valid-token").Return(mockClaims, nil)
    r := setupRouter(mockTokenSvc, "admin")

    req := httptest.NewRequest("GET", "/secure", nil)
    req.Header.Set("Authorization", "Bearer valid-token")
    resp := httptest.NewRecorder()

    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
    assert.Contains(t, resp.Body.String(), "Access granted")
}

func TestAuthMiddleware_MissingHeader(t *testing.T) {
    r := setupRouter(new(mocks.TokenService), "admin")
    req := httptest.NewRequest("GET", "/secure", nil)
    resp := httptest.NewRecorder()

    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusUnauthorized, resp.Code)
    assert.Contains(t, resp.Body.String(), "Missing Authorization header")
}

func TestAuthMiddleware_InvalidFormat(t *testing.T) {
    r := setupRouter(new(mocks.TokenService), "admin")
    req := httptest.NewRequest("GET", "/secure", nil)
    req.Header.Set("Authorization", "InvalidFormat")
    resp := httptest.NewRecorder()

    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusUnauthorized, resp.Code)
    assert.Contains(t, resp.Body.String(), "Invalid Authorization header format")
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
    mockTokenSvc := new(mocks.TokenService)
    mockTokenSvc.On("IVerifyToken", "bad-token").Return(nil, assert.AnError)

    r := setupRouter(mockTokenSvc, "admin")
    req := httptest.NewRequest("GET", "/secure", nil)
    req.Header.Set("Authorization", "Bearer bad-token")
    resp := httptest.NewRecorder()

    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusUnauthorized, resp.Code)
    assert.Contains(t, resp.Body.String(), "Invalid token")
}

func TestRoleRequired_FailsForWrongRole(t *testing.T) {
    mockTokenSvc := new(mocks.TokenService)
    mockClaims := jwt.MapClaims{
        "sub":  "user123",
        "role": "user", 
    }

    mockTokenSvc.On("IVerifyToken", "valid-token").Return(mockClaims, nil)

    r := setupRouter(mockTokenSvc, "admin")

    req := httptest.NewRequest("GET", "/secure", nil)
    req.Header.Set("Authorization", "Bearer valid-token")
    resp := httptest.NewRecorder()

    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusForbidden, resp.Code)
    assert.Contains(t, resp.Body.String(), "admin access required")
}
