package middleware

import (
	"github.com/gin-gonic/gin"
	"golang-review-phone/utils"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		claims, err := utils.VerifyToken(token)
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		requiredRoles := []string{"user", "admin"}
		userRole, ok := claims["role"].(string)
		if !ok || !contains(requiredRoles, userRole) {
			utils.RespondWithError(c, http.StatusForbidden, "Insufficient permissions")
			c.Abort()
			return
		}

		c.Next()
	}
}

func contains(slices []string, target string) bool {
	for _, s := range slices {
		if s == target {
			return true
		}
	}
	return false
}
