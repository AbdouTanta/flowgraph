package auth

import (
	"flowgraph/config"
	"flowgraph/utils"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	authRepository AuthRepository
}

func NewAuthService(authRepository *AuthRepository) *AuthService {
	return &AuthService{
		authRepository: *authRepository,
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

func (s *AuthService) GetUserFromToken(tokenString string) (*User, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.SigningKey), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	log.Printf("Parsed token: %+v", token)

	// Extract user ID from token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["id"] == nil {
		return nil, fmt.Errorf("invalid token claims")
	}

	userId := claims["id"].(string)

	// Fetch the user from the database using the user ID
	user, err := s.authRepository.FindUserById(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by ID: %v", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}
