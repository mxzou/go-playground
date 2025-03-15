package handlers

import (
	"encoding/json"
	"net/http"

	"playground/models"
	"playground/services"
)

// AuthHandler handles HTTP requests for authentication
type AuthHandler struct {
	userService *services.UserService
}

// NewAuthHandler creates a new auth handler with the given service
func NewAuthHandler(userService *services.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the login response body
type LoginResponse struct {
	Token string `json:"token"`
}

// Login authenticates a user and returns a JWT token
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Authenticate user
	token, err := h.userService.Authenticate(req.Username, req.Password)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	// Return token
	respondWithJSON(w, http.StatusOK, LoginResponse{Token: token})
}

// Register creates a new user
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input models.UserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Create user
	user, err := h.userService.CreateUser(input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Return user
	respondWithJSON(w, http.StatusCreated, user)
}