package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"playground/models"
	"playground/repositories"
)

// JWT secret key - in a real application, this would be stored securely
var jwtSecret = []byte("your-secret-key")

// TokenClaims represents the JWT claims
type TokenClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

// UserService handles business logic for users
type UserService struct {
	repository repositories.UserRepository
}

// NewUserService creates a new user service with the given repository
func NewUserService(repository repositories.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

// GetAllUsers returns all users
func (s *UserService) GetAllUsers() []models.User {
	return s.repository.FindAll()
}

// GetUserByID returns a user by ID
func (s *UserService) GetUserByID(id string) (models.User, error) {
	return s.repository.FindByID(id)
}

// CreateUser adds a new user
func (s *UserService) CreateUser(input models.UserInput) (models.User, error) {
	return s.repository.Create(input)
}

// UpdateUser modifies an existing user
func (s *UserService) UpdateUser(id string, input models.UserInput) (models.User, error) {
	return s.repository.Update(id, input)
}

// DeleteUser removes a user
func (s *UserService) DeleteUser(id string) error {
	return s.repository.Delete(id)
}

// Authenticate verifies user credentials and returns a JWT token
func (s *UserService) Authenticate(username, password string) (string, error) {
	// Find user by username
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Check password
	if !user.CheckPassword(password) {
		return "", errors.New("invalid credentials")
	}

	// Create token claims
	claims := TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: user.ID,
		Role:   user.Role,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the claims
func (s *UserService) ValidateToken(tokenString string) (*TokenClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Validate token
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Get claims
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}