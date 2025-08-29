package auth

import (
	"flowgraph/config"
	"flowgraph/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	authRepository AuthRepository
}

func NewAuthService(authRepository AuthRepository) *AuthService {
	return &AuthService{
		authRepository: authRepository,
	}
}

func (s *AuthService) Login(loginPayload User) (*User, string, error) {
	user, err := s.authRepository.Login(loginPayload.Email, loginPayload.Password)
	if err != nil {
		return nil, "", fmt.Errorf("failed to login user: %v", err)
	}

	// Construct and return the User object and the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"email":    user.Email,
	})

	tokenString, err := token.SignedString([]byte(config.Config.SigningKey))
	if err != nil {
		return nil, "", err
	}

	return user, tokenString, nil
}

func (s *AuthService) Register(userPayload User) (*User, string, error) {
	// TODO - Validate the input
	// Check if the user already exists with the same email
	existingUser, err := s.authRepository.FindUserByEmail(userPayload.Email)
	if err != nil {
		return nil, "", fmt.Errorf("failed to check existing user: %v", err)
	}
	if existingUser != nil {
		return nil, "", fmt.Errorf("user with email %s already exists", userPayload.Email)
	}

	// Check if the user already exists with the same username
	existingUser, err = s.authRepository.FindUserByUsername(userPayload.Username)
	if err != nil {
		return nil, "", fmt.Errorf("failed to check existing user: %v", err)
	}
	if existingUser != nil {
		return nil, "", fmt.Errorf("user with username %s already exists", userPayload.Username)
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(userPayload.Password)
	if err != nil {
		return nil, "", fmt.Errorf("failed to hash password: %v", err)
	}

	user, err := s.authRepository.Register(userPayload.Email, userPayload.Username, hashedPassword)
	if err != nil {
		return nil, "", fmt.Errorf("failed to register user: %v", err)
	}

	// Construct and return the User object and the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id.Hex(),
		"username": user.Username,
		"email":    user.Email,
	})

	tokenString, err := token.SignedString([]byte(config.Config.SigningKey))
	if err != nil {
		return nil, "", err
	}

	return user, tokenString, nil
}
