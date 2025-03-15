package services

import (
	"playground/models"
	"playground/repositories"
)

// RecipeService handles business logic for recipes
type RecipeService struct {
	repository repositories.RecipeRepository
}

// NewRecipeService creates a new recipe service with the given repository
func NewRecipeService(repository repositories.RecipeRepository) *RecipeService {
	return &RecipeService{
		repository: repository,
	}
}

// GetAllRecipes returns all recipes
func (s *RecipeService) GetAllRecipes() []models.Recipe {
	return s.repository.FindAll()
}

// GetRecipeByID returns a recipe by ID
func (s *RecipeService) GetRecipeByID(id string) (models.Recipe, error) {
	return s.repository.FindByID(id)
}

// CreateRecipe adds a new recipe
func (s *RecipeService) CreateRecipe(input models.RecipeInput) models.Recipe {
	return s.repository.Create(input)
}

// UpdateRecipe modifies an existing recipe
func (s *RecipeService) UpdateRecipe(id string, input models.RecipeInput) (models.Recipe, error) {
	return s.repository.Update(id, input)
}

// DeleteRecipe removes a recipe
func (s *RecipeService) DeleteRecipe(id string) error {
	return s.repository.Delete(id)
}

// FilterRecipesByTag returns recipes that have the specified tag
// This demonstrates a higher-order function that takes a predicate function
func (s *RecipeService) FilterRecipesByTag(tag string) []models.Recipe {
	allRecipes := s.repository.FindAll()
	return Filter(allRecipes, func(recipe models.Recipe) bool {
		return Contains(recipe.Tags, tag)
	})
}

// Filter is a higher-order function that filters a slice based on a predicate
func Filter[T any](items []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Contains checks if a slice contains a specific value
func Contains[T comparable](items []T, item T) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

// SortBy defines the criteria for sorting recipes
type SortBy string

// Sort criteria constants
const (
	SortByPrepTime  SortBy = "prepTime"
	SortByCookTime  SortBy = "cookTime"
	SortByTotalTime SortBy = "totalTime"
	SortByTitle     SortBy = "title"
	SortByServings  SortBy = "servings"
)

// SortRecipes returns recipes sorted by the specified criteria
func (s *RecipeService) SortRecipes(criteria SortBy, ascending bool) []models.Recipe {
	allRecipes := s.repository.FindAll()

	// Define a less function based on the sorting criteria
	less := func(i, j int) bool {
		switch criteria {
		case SortByPrepTime:
			if ascending {
				return allRecipes[i].PrepTime < allRecipes[j].PrepTime
			}
			return allRecipes[i].PrepTime > allRecipes[j].PrepTime
		case SortByCookTime:
			if ascending {
				return allRecipes[i].CookTime < allRecipes[j].CookTime
			}
			return allRecipes[i].CookTime > allRecipes[j].CookTime
		case SortByTotalTime:
			totalTimeI := allRecipes[i].PrepTime + allRecipes[i].CookTime
			totalTimeJ := allRecipes[j].PrepTime + allRecipes[j].CookTime
			if ascending {
				return totalTimeI < totalTimeJ
			}
			return totalTimeI > totalTimeJ
		case SortByTitle:
			if ascending {
				return allRecipes[i].Title < allRecipes[j].Title
			}
			return allRecipes[i].Title > allRecipes[j].Title
		case SortByServings:
			if ascending {
				return allRecipes[i].Servings < allRecipes[j].Servings
			}
			return allRecipes[i].Servings > allRecipes[j].Servings
		default:
			// Default to sorting by title
			if ascending {
				return allRecipes[i].Title < allRecipes[j].Title
			}
			return allRecipes[i].Title > allRecipes[j].Title
		}
	}

	// Create a copy of the slice to avoid modifying the original
	result := make([]models.Recipe, len(allRecipes))
	copy(result, allRecipes)

	// Sort the copy using the less function
	Sort(result, less)

	return result
}

// Sort is a generic function that sorts a slice using the provided less function
func Sort[T any](items []T, less func(i, j int) bool) {
	n := len(items)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if less(j, i) {
				items[i], items[j] = items[j], items[i]
			}
		}
	}
}
