package auth

import (
	"log"

	"github.com/gin-gonic/gin"
)

type AuthRestController struct {
	authService *AuthService
}

func NewAuthRestController(authService *AuthService) *AuthRestController {
	return &AuthRestController{
		authService: authService,
	}
}

func (h *AuthRestController) Login(c *gin.Context) {
	var loginPayload User
	if err := c.ShouldBindJSON(&loginPayload); err != nil {
		log.Printf("Error binding user data: %v", err)
		c.JSON(400, gin.H{"error": "Invalid user data"})
		return
	}

	// Insert the new flow into the database
	user, err := h.authService.Login(loginPayload)
	if err != nil {
		log.Printf("Error logging user in: %v", err)
		c.JSON(500, gin.H{"error": "Failed to login user"})
		return
	}

	c.JSON(201, gin.H{"user": user})
}
