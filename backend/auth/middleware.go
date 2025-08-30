package auth

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

// Expects the Authorization header to be set with a user bearer token
func AuthMiddleware(authController *AuthRestController) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for bearer authorizationHeader in Authorization header
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// Extract token from "Bearer <token>" format
		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authorizationHeader, bearerPrefix) {
			c.JSON(401, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		// Find existing user from jwt token
		tokenString := strings.TrimPrefix(authorizationHeader, bearerPrefix)
		user, err := authController.authService.GetUserFromToken(tokenString)
		if err != nil {
			log.Printf("Error finding user from token: %v", err)
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		if user == nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// Set the user id in context and continue
		c.Set("user_id", user.Id.Hex())
		c.Next()
	}
}
