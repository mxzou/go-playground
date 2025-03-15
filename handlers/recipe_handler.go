package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"playground/models"
	"playground/services"
)

// RecipeHandler handles HTTP requests for recipes
type RecipeHandler struct {
	service *services.RecipeService
}

// NewRecipeHandler creates a new recipe handler with the given service
func NewRecipeHandler(service *services.RecipeService) *RecipeHandler {
	return &RecipeHandler{
		service: service,
	}
}

// GetAllRecipes returns all recipes as JSON
func (h *RecipeHandler) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	recipes := h.service.GetAllRecipes()
	respondWithJSON(w, http.StatusOK, recipes)
}

// GetRecipeByID returns a recipe by ID as JSON
func (h *RecipeHandler) GetRecipeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	recipe, err := h.service.GetRecipeByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Recipe not found")
		return
	}

	respondWithJSON(w, http.StatusOK, recipe)
}

// CreateRecipe creates a new recipe from JSON request body
func (h *RecipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var input models.RecipeInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	recipe := h.service.CreateRecipe(input)
	respondWithJSON(w, http.StatusCreated, recipe)
}

// UpdateRecipe updates an existing recipe from JSON request body
func (h *RecipeHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var input models.RecipeInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	recipe, err := h.service.UpdateRecipe(id, input)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Recipe not found")
		return
	}

	respondWithJSON(w, http.StatusOK, recipe)
}

// DeleteRecipe deletes a recipe by ID
func (h *RecipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.service.DeleteRecipe(id); err != nil {
		respondWithError(w, http.StatusNotFound, "Recipe not found")
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}

// Helper function to respond with JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Helper function to respond with an error
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}