package services

import (
	"playground/models"
	"strings"
)

// SearchService handles search functionality for recipes
type SearchService struct {
	recipeService *RecipeService
}

// NewSearchService creates a new search service with the given recipe service
func NewSearchService(recipeService *RecipeService) *SearchService {
	return &SearchService{
		recipeService: recipeService,
	}
}

// SearchByIngredient returns recipes that contain the specified ingredient
func (s *SearchService) SearchByIngredient(ingredient string) []models.Recipe {
	allRecipes := s.recipeService.GetAllRecipes()
	return Filter(allRecipes, func(recipe models.Recipe) bool {
		for _, ing := range recipe.Ingredients {
			if strings.Contains(strings.ToLower(ing), strings.ToLower(ingredient)) {
				return true
			}
		}
		return false
	})
}

// SearchByTag returns recipes that have the specified tag
func (s *SearchService) SearchByTag(tag string) []models.Recipe {
	return s.recipeService.FilterRecipesByTag(tag)
}

// SearchByTitle returns recipes that contain the specified title
func (s *SearchService) SearchByTitle(title string) []models.Recipe {
	allRecipes := s.recipeService.GetAllRecipes()
	return Filter(allRecipes, func(recipe models.Recipe) bool {
		return strings.Contains(strings.ToLower(recipe.Title), strings.ToLower(title))
	})
}

// GetPaginatedRecipes returns a paginated list of recipes
func (s *SearchService) GetPaginatedRecipes(page, pageSize int) []models.Recipe {
	allRecipes := s.recipeService.GetAllRecipes()
	
	// Calculate start and end indices
	start := (page - 1) * pageSize
	end := start + pageSize
	
	// Check bounds
	if start >= len(allRecipes) {
		return []models.Recipe{}
	}
	if end > len(allRecipes) {
		end = len(allRecipes)
	}
	
	return allRecipes[start:end]
}