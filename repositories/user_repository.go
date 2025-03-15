package repositories

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"playground/models"
)

// UserRepository defines the interface for user storage operations
type UserRepository interface {
	FindAll() []models.User
	FindByID(id string) (models.User, error)
	FindByUsername(username string) (models.User, error)
	FindByEmail(email string) (models.User, error)
	Create(input models.UserInput) (models.User, error)
	Update(id string, input models.UserInput) (models.User, error)
	Delete(id string) error
}

// InMemoryUserRepository implements UserRepository with in-memory storage
type InMemoryUserRepository struct {
	users map[string]models.User
	mutex sync.RWMutex
}

// NewInMemoryUserRepository creates a new in-memory user repository
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]models.User),
	}
}

// FindAll returns all users
func (r *InMemoryUserRepository) FindAll() []models.User {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	result := make([]models.User, 0, len(r.users))
	for _, user := range r.users {
		result = append(result, user)
	}
	return result
}

// FindByID returns a user by ID
func (r *InMemoryUserRepository) FindByID(id string) (models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

// FindByUsername returns a user by username
func (r *InMemoryUserRepository) FindByUsername(username string) (models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

// FindByEmail returns a user by email
func (r *InMemoryUserRepository) FindByEmail(email string) (models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

// Create adds a new user
func (r *InMemoryUserRepository) Create(input models.UserInput) (models.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Check if username already exists
	for _, user := range r.users {
		if user.Username == input.Username {
			return models.User{}, errors.New("username already exists")
		}
		if user.Email == input.Email {
			return models.User{}, errors.New("email already exists")
		}
	}

	id := uuid.New().String()
	user, err := models.NewUser(id, input)
	if err != nil {
		return models.User{}, err
	}
	
	// Store a copy of the user (immutable pattern)
	r.users[id] = user
	
	return user, nil
}

// Update modifies an existing user
func (r *InMemoryUserRepository) Update(id string, input models.UserInput) (models.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	original, exists := r.users[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}

	// Check if username already exists (if changing username)
	if input.Username != original.Username {
		for _, user := range r.users {
			if user.ID != id && user.Username == input.Username {
				return models.User{}, errors.New("username already exists")
			}
		}
	}

	// Check if email already exists (if changing email)
	if input.Email != original.Email {
		for _, user := range r.users {
			if user.ID != id && user.Email == input.Email {
				return models.User{}, errors.New("email already exists")
			}
		}
	}

	// Create a new user with updated fields (immutable pattern)
	updated, err := models.UpdateUser(original, input)
	if err != nil {
		return models.User{}, err
	}
	
	// Store the updated user
	r.users[id] = updated
	
	return updated, nil
}

// Delete removes a user
func (r *InMemoryUserRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, exists := r.users[id]
	if !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}