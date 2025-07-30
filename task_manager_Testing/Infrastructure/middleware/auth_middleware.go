package middleware

import (
	"net/http"
	"strings"
	"task_manager_Testing/Domain/services"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenService services.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := splitToken[1]

		claims, err := tokenService.IVerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims["sub"])
		c.Set("role", claims["role"])
		c.Next()
	}
}


func RoleRequired(requiredRole string) gin.HandlerFunc {
    return func(c *gin.Context) {
        roleVal, exists := c.Get("role")
        role, ok := roleVal.(string)

        if !exists || !ok || role != requiredRole {
            c.JSON(http.StatusForbidden, gin.H{"error": requiredRole + " access required"})
            c.Abort()
            return
        }
        c.Next()
    }
}
