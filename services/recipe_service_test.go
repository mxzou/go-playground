package services

import (
	"testing"

	"playground/models"
	"playground/repositories"
)

// TestCreateRecipe tests the CreateRecipe function of RecipeService
func TestCreateRecipe(t *testing.T) {
	// Create a repository
	repo := repositories.NewInMemoryRecipeRepository()

	// Create the service with the repository
	service := NewRecipeService(repo)

	// Test creating a recipe
	recipeInput := models.RecipeInput{
		Title:        "Test Recipe",
		Description:  "A test recipe for unit testing",
		Ingredients:  []string{"ingredient1", "ingredient2"},
		Instructions: []string{"step1", "step2"},
		PrepTime:     10,
		CookTime:     20,
		Servings:     4,
		Tags:         []string{"test", "unit-test"},
	}

	recipe := service.CreateRecipe(recipeInput)

	// Check recipe properties
	if recipe.Title != recipeInput.Title {
		t.Errorf("Expected title %s, but got %s", recipeInput.Title, recipe.Title)
	}

	if recipe.Description != recipeInput.Description {
		t.Errorf("Expected description %s, but got %s", recipeInput.Description, recipe.Description)
	}

	if len(recipe.Ingredients) != len(recipeInput.Ingredients) {
		t.Errorf("Expected %d ingredients, but got %d", len(recipeInput.Ingredients), len(recipe.Ingredients))
	}

	if len(recipe.Tags) != len(recipeInput.Tags) {
		t.Errorf("Expected %d tags, but got %d", len(recipeInput.Tags), len(recipe.Tags))
	}
}

// TestFilterRecipesByTag tests the FilterRecipesByTag function of RecipeService
func TestFilterRecipesByTag(t *testing.T) {
	// Create a repository
	repo := repositories.NewInMemoryRecipeRepository()

	// Create the service with the repository
	service := NewRecipeService(repo)

	// Create test recipes with different tags
	recipe1 := models.RecipeInput{
		Title: "Recipe 1",
		Tags:  []string{"vegetarian", "quick"},
	}

	recipe2 := models.RecipeInput{
		Title: "Recipe 2",
		Tags:  []string{"meat", "dinner"},
	}

	recipe3 := models.RecipeInput{
		Title: "Recipe 3",
		Tags:  []string{"vegetarian", "dinner"},
	}

	// Add recipes to repository
	service.CreateRecipe(recipe1)
	service.CreateRecipe(recipe2)
	service.CreateRecipe(recipe3)

	// Test filtering by tag
	tests := []struct {
		tag           string
		expectedCount int
	}{
		{"vegetarian", 2},
		{"meat", 1},
		{"dinner", 2},
		{"quick", 1},
		{"nonexistent", 0},
	}

	for _, tc := range tests {
		t.Run(tc.tag, func(t *testing.T) {
			results := service.FilterRecipesByTag(tc.tag)

			if len(results) != tc.expectedCount {
				t.Errorf("Expected %d recipes with tag '%s', but got %d",
					tc.expectedCount, tc.tag, len(results))
			}
		})
	}
}

// TestFilter tests the Filter utility function
func TestFilter(t *testing.T) {
	// Test filtering integers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filter even numbers
	evenNumbers := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})

	expectedEven := []int{2, 4, 6, 8, 10}

	if len(evenNumbers) != len(expectedEven) {
		t.Errorf("Expected %d even numbers, but got %d", len(expectedEven), len(evenNumbers))
	}

	// Test filtering strings
	words := []string{"apple", "banana", "cherry", "date", "elderberry"}

	// Filter words starting with 'a' or 'b'
	abWords := Filter(words, func(word string) bool {
		return word[0] == 'a' || word[0] == 'b'
	})

	expectedAB := []string{"apple", "banana"}

	if len(abWords) != len(expectedAB) {
		t.Errorf("Expected %d words starting with 'a' or 'b', but got %d",
			len(expectedAB), len(abWords))
	}
}

// TestSortRecipes tests the SortRecipes function of RecipeService
func TestSortRecipes(t *testing.T) {
	// Create a repository
	repo := repositories.NewInMemoryRecipeRepository()

	// Create the service with the repository
	service := NewRecipeService(repo)

	// Create test recipes with different properties
	recipe1 := models.RecipeInput{
		Title:    "Pasta Carbonara",
		PrepTime: 15,
		CookTime: 20,
		Servings: 4,
	}

	recipe2 := models.RecipeInput{
		Title:    "Quick Salad",
		PrepTime: 10,
		CookTime: 0,
		Servings: 2,
	}

	recipe3 := models.RecipeInput{
		Title:    "Beef Stew",
		PrepTime: 30,
		CookTime: 120,
		Servings: 6,
	}

	// Add recipes to repository
	service.CreateRecipe(recipe1)
	service.CreateRecipe(recipe2)
	service.CreateRecipe(recipe3)

	// Test sorting by different criteria
	tests := []struct {
		name      string
		criteria  SortBy
		ascending bool
		expected  []string // Expected order of recipe titles
	}{
		{"PrepTime ascending", SortByPrepTime, true, []string{"Quick Salad", "Pasta Carbonara", "Beef Stew"}},
		{"PrepTime descending", SortByPrepTime, false, []string{"Beef Stew", "Pasta Carbonara", "Quick Salad"}},
		{"CookTime ascending", SortByCookTime, true, []string{"Quick Salad", "Pasta Carbonara", "Beef Stew"}},
		{"CookTime descending", SortByCookTime, false, []string{"Beef Stew", "Pasta Carbonara", "Quick Salad"}},
		{"TotalTime ascending", SortByTotalTime, true, []string{"Quick Salad", "Pasta Carbonara", "Beef Stew"}},
		{"TotalTime descending", SortByTotalTime, false, []string{"Beef Stew", "Pasta Carbonara", "Quick Salad"}},
		{"Title ascending", SortByTitle, true, []string{"Beef Stew", "Pasta Carbonara", "Quick Salad"}},
		{"Title descending", SortByTitle, false, []string{"Quick Salad", "Pasta Carbonara", "Beef Stew"}},
		{"Servings ascending", SortByServings, true, []string{"Quick Salad", "Pasta Carbonara", "Beef Stew"}},
		{"Servings descending", SortByServings, false, []string{"Beef Stew", "Pasta Carbonara", "Quick Salad"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			results := service.SortRecipes(tc.criteria, tc.ascending)

			if len(results) != 3 {
				t.Fatalf("Expected 3 recipes, but got %d", len(results))
			}

			// Check if the order matches expected
			for i, expectedTitle := range tc.expected {
				if results[i].Title != expectedTitle {
					t.Errorf("Expected recipe at position %d to be '%s', but got '%s'",
						i, expectedTitle, results[i].Title)
				}
			}
		})
	}
}
