package handlers

import (
	"net/http"
	"playground/services"
	"strings"
)

// SortHandler handles HTTP requests for sorting recipes
type SortHandler struct {
	recipeService *services.RecipeService
}

// NewSortHandler creates a new sort handler with the given service
func NewSortHandler(recipeService *services.RecipeService) *SortHandler {
	return &SortHandler{
		recipeService: recipeService,
	}
}

// SortRecipes returns recipes sorted by the specified criteria
func (h *SortHandler) SortRecipes(w http.ResponseWriter, r *http.Request) {
	// Get sort criteria from query parameters
	criteriaParam := r.URL.Query().Get("criteria")
	orderParam := r.URL.Query().Get("order")
	
	// Default values
	criteria := services.SortByTitle
	ascending := true
	
	// Parse criteria parameter
	switch strings.ToLower(criteriaParam) {
	case "preptime":
		criteria = services.SortByPrepTime
	case "cooktime":
		criteria = services.SortByCookTime
	case "totaltime":
		criteria = services.SortByTotalTime
	case "title":
		criteria = services.SortByTitle
	case "servings":
		criteria = services.SortByServings
	}
	
	// Parse order parameter
	if strings.ToLower(orderParam) == "desc" || strings.ToLower(orderParam) == "descending" {
		ascending = false
	}
	
	// Get sorted recipes
	recipes := h.recipeService.SortRecipes(criteria, ascending)
	
	// Return sorted recipes
	respondWithJSON(w, http.StatusOK, recipes)
}