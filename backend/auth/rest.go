package auth

import (
	"flowgraph/utils"

	"github.com/gin-gonic/gin"
)

type AuthRestController struct {
	authService *AuthService
	responder   *utils.HTTPResponder
}

func NewAuthRestController(authService *AuthService, httpResponder *utils.HTTPResponder) *AuthRestController {
	return &AuthRestController{
		authService: authService,
		responder:   httpResponder,
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
		h.responder.BadRequest(c, "Invalid user data", err)
		return
	}

	// Insert the new flow into the database
	user, token, err := h.authService.Login(loginPayload)
	if err != nil {
		h.responder.InternalError(c, "Failed to login user", err)
		return
	}

	userDto := UserDTO{
		Id:       user.Id.Hex(),
		Username: user.Username,
		Email:    user.Email,
	}

	h.responder.SuccessWithMessage(c, 201, "Login successful", gin.H{"user": userDto, "token": token})
}

func (h *AuthRestController) Register(c *gin.Context) {
	var registerPayload User
	if err := c.ShouldBindJSON(&registerPayload); err != nil {
		h.responder.BadRequest(c, "Invalid user data", err)
		return
	}

	// Insert the new flow into the database
	user, token, err := h.authService.Register(registerPayload)
	if err != nil {
		h.responder.InternalError(c, "Failed to register user", err)
		return
	}

	userDto := UserDTO{
		Id:       user.Id.Hex(),
		Username: user.Username,
		Email:    user.Email,
	}

	h.responder.SuccessWithMessage(c, 201, "Registration successful", gin.H{"user": userDto, "token": token})
}
