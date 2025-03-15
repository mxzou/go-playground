package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"playground/middleware"
	"playground/models"
	"playground/services"
)

// RatingHandler handles HTTP requests for recipe ratings
type RatingHandler struct {
	ratingService *services.RatingService
}

// NewRatingHandler creates a new rating handler with the given service
func NewRatingHandler(ratingService *services.RatingService) *RatingHandler {
	return &RatingHandler{
		ratingService: ratingService,
	}
}

// GetRatingsByRecipeID returns all ratings for a specific recipe
func (h *RatingHandler) GetRatingsByRecipeID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recipeID := vars["recipeId"]

	ratings := h.ratingService.GetRatingsByRecipeID(recipeID)
	respondWithJSON(w, http.StatusOK, ratings)
}

// GetAverageRatingForRecipe returns the average rating score for a recipe
func (h *RatingHandler) GetAverageRatingForRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recipeID := vars["recipeId"]

	average := h.ratingService.GetAverageRatingForRecipe(recipeID)
	respondWithJSON(w, http.StatusOK, map[string]float64{"average": average})
}

// CreateRating adds a new rating for a recipe
func (h *RatingHandler) CreateRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recipeID := vars["recipeId"]

	// Get user ID from context (set by auth middleware)
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var input models.RatingInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate rating score
	if input.Score < 1 || input.Score > 5 {
		respondWithError(w, http.StatusBadRequest, "Rating score must be between 1 and 5")
		return
	}

	rating, err := h.ratingService.CreateRating(recipeID, userID, input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, rating)
}

// UpdateRating modifies an existing rating
func (h *RatingHandler) UpdateRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Get user ID from context (set by auth middleware)
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var input models.RatingInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate rating score
	if input.Score < 1 || input.Score > 5 {
		respondWithError(w, http.StatusBadRequest, "Rating score must be between 1 and 5")
		return
	}

	rating, err := h.ratingService.UpdateRating(id, userID, input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, rating)
}

// DeleteRating removes a rating
func (h *RatingHandler) DeleteRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Get user ID from context (set by auth middleware)
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	err := h.ratingService.DeleteRating(id, userID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}