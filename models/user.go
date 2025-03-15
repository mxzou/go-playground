package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Password is not exposed in JSON
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// UserInput represents the data needed to create or update a user
type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}

// NewUser creates a new User with the given input and generated ID
func NewUser(id string, input UserInput) (User, error) {
	now := time.Now()
	
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	
	// Default role if not provided
	role := input.Role
	if role == "" {
		role = "user"
	}
	
	return User{
		ID:        id,
		Username:  input.Username,
		Email:     input.Email,
		Password:  string(hashedPassword),
		Role:      role,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// CheckPassword verifies if the provided password matches the stored hash
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// UpdateUser creates a new User with updated fields but preserves the original ID and creation time
func UpdateUser(original User, input UserInput) (User, error) {
	updated := original
	updated.Username = input.Username
	updated.Email = input.Email
	updated.UpdatedAt = time.Now()
	
	// Only update role if provided
	if input.Role != "" {
		updated.Role = input.Role
	}
	
	// Only update password if provided
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return User{}, err
		}
		updated.Password = string(hashedPassword)
	}
	
	return updated, nil
}