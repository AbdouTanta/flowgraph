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

type UserDTO struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (h *AuthRestController) Login(c *gin.Context) {
	var loginPayload User
	if err := c.ShouldBindJSON(&loginPayload); err != nil {
		log.Printf("Error binding user data: %v", err)
		c.JSON(400, gin.H{"error": "Invalid user data"})
		return
	}

	// Insert the new flow into the database
	user, token, err := h.authService.Login(loginPayload)
	if err != nil {
		log.Printf("Error logging user in: %v", err)
		c.JSON(500, gin.H{"error": "Failed to login user"})
		return
	}

	userDto := UserDTO{
		Id:       user.Id.Hex(),
		Username: user.Username,
		Email:    user.Email,
	}

	c.JSON(201, gin.H{"user": userDto, "token": token})
}

func (h *AuthRestController) Register(c *gin.Context) {
	var registerPayload User
	if err := c.ShouldBindJSON(&registerPayload); err != nil {
		log.Printf("Error binding user data: %v", err)
		c.JSON(400, gin.H{"error": "Invalid user data"})
		return
	}

	// Insert the new flow into the database
	user, token, err := h.authService.Register(registerPayload)
	if err != nil {
		log.Printf("Error registering user in: %v", err)
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	userDto := UserDTO{
		Id:       user.Id.Hex(),
		Username: user.Username,
		Email:    user.Email,
	}

	c.JSON(201, gin.H{"user": userDto, "token": token})
}
