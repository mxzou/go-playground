package services

import (
	"testing"

	"playground/models"
	"playground/repositories"
)

// TestAuthenticate tests the Authenticate function of UserService
func TestAuthenticate(t *testing.T) {
	// Create a repository with a test user
	repo := repositories.NewInMemoryUserRepository()
	userInput := models.UserInput{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
		Role:     "user",
	}

	// Create the user
	repo.Create(userInput)

	// Create the service with the repository
	service := NewUserService(repo)

	// Test cases
	tests := []struct {
		name          string
		username      string
		password      string
		shouldSucceed bool
	}{
		{"Valid credentials", "testuser", "password123", true},
		{"Invalid username", "wronguser", "password123", false},
		{"Invalid password", "testuser", "wrongpassword", false},
	}

	// Run test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			token, err := service.Authenticate(tc.username, tc.password)

			if tc.shouldSucceed {
				// Should succeed
				if err != nil {
					t.Errorf("Expected authentication to succeed, but got error: %v", err)
				}
				if token == "" {
					t.Error("Expected token to be non-empty")
				}
			} else {
				// Should fail
				if err == nil {
					t.Error("Expected authentication to fail, but it succeeded")
				}
				if token != "" {
					t.Error("Expected token to be empty")
				}
			}
		})
	}
}

// TestCreateUser tests the CreateUser function of UserService
func TestCreateUser(t *testing.T) {
	// Create a repository
	repo := repositories.NewInMemoryUserRepository()

	// Create the service with the repository
	service := NewUserService(repo)

	// Test creating a user
	userInput := models.UserInput{
		Username: "newuser",
		Email:    "new@example.com",
		Password: "newpassword",
		Role:     "user",
	}

	user, err := service.CreateUser(userInput)

	// Check for errors
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	// Check user properties
	if user.Username != userInput.Username {
		t.Errorf("Expected username %s, but got %s", userInput.Username, user.Username)
	}

	if user.Email != userInput.Email {
		t.Errorf("Expected email %s, but got %s", userInput.Email, user.Email)
	}

	if user.Role != userInput.Role {
		t.Errorf("Expected role %s, but got %s", userInput.Role, user.Role)
	}

	// Test creating a duplicate user (should fail)
	_, err = service.CreateUser(userInput)
	if err == nil {
		t.Error("Expected error when creating duplicate user, but got none")
	}
}
