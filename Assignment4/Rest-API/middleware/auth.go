package middleware

import (
	"net/http"
	"rest-api/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logrus.Warn("Authorization header missing")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			logrus.Warnf("Invalid Authorization header format: %s", authHeader)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}

		tokenString := parts[1]
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			logrus.Warnf("Invalid token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		userID, okUserID := claims["user_id"].(float64)
		role, okRole := claims["role"].(string)

		if !okUserID || !okRole {
			logrus.Error("Token claims missing user_id, role, username, or type assertion failed")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		c.Set("userID", uint(userID))
		c.Set("role", role)

		if len(allowedRoles) > 0 {
			roleMatch := false
			for _, ar := range allowedRoles {
				if role == ar {
					roleMatch = true
					break
				}
			}
			if !roleMatch {
				logrus.Warnf("UserID: %d with role '%s' tried to access a resource restricted to roles: %v", uint(userID), role, allowedRoles)
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
				return
			}
		}

		c.Next()
	}
}
