package handlers

import (
	"net/http"
	"playground/services"
	"strconv"
)

// SearchHandler handles HTTP requests for searching recipes
type SearchHandler struct {
	searchService *services.SearchService
}

// NewSearchHandler creates a new search handler with the given service
func NewSearchHandler(searchService *services.SearchService) *SearchHandler {
	return &SearchHandler{
		searchService: searchService,
	}
}

// SearchByIngredient returns recipes containing the specified ingredient
func (h *SearchHandler) SearchByIngredient(w http.ResponseWriter, r *http.Request) {
	ingredient := r.URL.Query().Get("q")
	if ingredient == "" {
		respondWithError(w, http.StatusBadRequest, "Missing query parameter")
		return
	}

	recipes := h.searchService.SearchByIngredient(ingredient)
	respondWithJSON(w, http.StatusOK, recipes)
}

// SearchByTag returns recipes with the specified tag
func (h *SearchHandler) SearchByTag(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("q")
	if tag == "" {
		respondWithError(w, http.StatusBadRequest, "Missing query parameter")
		return
	}

	recipes := h.searchService.SearchByTag(tag)
	respondWithJSON(w, http.StatusOK, recipes)
}

// SearchByTitle returns recipes containing the specified title
func (h *SearchHandler) SearchByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("q")
	if title == "" {
		respondWithError(w, http.StatusBadRequest, "Missing query parameter")
		return
	}

	recipes := h.searchService.SearchByTitle(title)
	respondWithJSON(w, http.StatusOK, recipes)
}

// GetPaginatedRecipes returns a paginated list of recipes
func (h *SearchHandler) GetPaginatedRecipes(w http.ResponseWriter, r *http.Request) {
	// Default values
	page := 1
	pageSize := 10

	// Parse page parameter
	pageParam := r.URL.Query().Get("page")
	if pageParam != "" {
		pageVal, err := strconv.Atoi(pageParam)
		if err != nil || pageVal < 1 {
			respondWithError(w, http.StatusBadRequest, "Invalid page parameter")
			return
		}
		page = pageVal
	}

	// Parse pageSize parameter
	pageSizeParam := r.URL.Query().Get("pageSize")
	if pageSizeParam != "" {
		pageSizeVal, err := strconv.Atoi(pageSizeParam)
		if err != nil || pageSizeVal < 1 || pageSizeVal > 100 {
			respondWithError(w, http.StatusBadRequest, "Invalid pageSize parameter")
			return
		}
		pageSize = pageSizeVal
	}

	recipes := h.searchService.GetPaginatedRecipes(page, pageSize)
	respondWithJSON(w, http.StatusOK, recipes)
}
