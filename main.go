package main

import (
	"log"
	"net/http"

	"playground/handlers"
	"playground/middleware"
	"playground/repositories"
	"playground/services"

	"github.com/gorilla/mux"
)

func main() {
	// Create repositories
	userRepo := repositories.NewInMemoryUserRepository()
	recipeRepo := repositories.NewInMemoryRecipeRepository()
	ratingRepo := repositories.NewInMemoryRatingRepository()

	// Create services
	userService := services.NewUserService(userRepo)
	recipeService := services.NewRecipeService(recipeRepo)
	ratingService := services.NewRatingService(ratingRepo, recipeService)
	searchService := services.NewSearchService(recipeService)

	// Create handlers
	authHandler := handlers.NewAuthHandler(userService)
	recipeHandler := handlers.NewRecipeHandler(recipeService)
	ratingHandler := handlers.NewRatingHandler(ratingService)
	searchHandler := handlers.NewSearchHandler(searchService)
	sortHandler := handlers.NewSortHandler(recipeService)

	// Create router
	router := mux.NewRouter()

	// Apply global middleware
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RecoveryMiddleware)

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Auth routes
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", authHandler.Register).Methods("POST")
	auth.HandleFunc("/login", authHandler.Login).Methods("POST")

	// Recipe routes
	recipes := api.PathPrefix("/recipes").Subrouter()
	recipes.HandleFunc("", recipeHandler.GetAllRecipes).Methods("GET")
	recipes.HandleFunc("/", recipeHandler.GetAllRecipes).Methods("GET")
	recipes.HandleFunc("/{id}", recipeHandler.GetRecipeByID).Methods("GET")
	recipes.HandleFunc("", recipeHandler.CreateRecipe).Methods("POST")
	recipes.HandleFunc("/", recipeHandler.CreateRecipe).Methods("POST")
	recipes.HandleFunc("/{id}", recipeHandler.UpdateRecipe).Methods("PUT")
	recipes.HandleFunc("/{id}", recipeHandler.DeleteRecipe).Methods("DELETE")

	// Protected recipe routes (require authentication)
	protectedRecipes := api.PathPrefix("/recipes").Subrouter()
	protectedRecipes.Use(middleware.AuthMiddleware(userService))
	protectedRecipes.HandleFunc("", recipeHandler.CreateRecipe).Methods("POST")
	protectedRecipes.HandleFunc("/", recipeHandler.CreateRecipe).Methods("POST")
	protectedRecipes.HandleFunc("/{id}", recipeHandler.UpdateRecipe).Methods("PUT")
	protectedRecipes.HandleFunc("/{id}", recipeHandler.DeleteRecipe).Methods("DELETE")

	// Rating routes
	ratings := api.PathPrefix("/recipes/{id}/ratings").Subrouter()
	ratings.HandleFunc("", ratingHandler.GetRatingsByRecipeID).Methods("GET")
	ratings.HandleFunc("/average", ratingHandler.GetAverageRatingForRecipe).Methods("GET")

	// Protected rating routes (require authentication)
	protectedRatings := api.PathPrefix("/recipes/{id}/ratings").Subrouter()
	protectedRatings.Use(middleware.AuthMiddleware(userService))
	protectedRatings.HandleFunc("", ratingHandler.CreateRating).Methods("POST")
	protectedRatings.HandleFunc("/{ratingId}", ratingHandler.UpdateRating).Methods("PUT")
	protectedRatings.HandleFunc("/{ratingId}", ratingHandler.DeleteRating).Methods("DELETE")

	// Search routes
	search := api.PathPrefix("/search").Subrouter()
	search.HandleFunc("/ingredient", searchHandler.SearchByIngredient).Methods("GET")
	search.HandleFunc("/tag", searchHandler.SearchByTag).Methods("GET")
	search.HandleFunc("/title", searchHandler.SearchByTitle).Methods("GET")
	search.HandleFunc("/paginated", searchHandler.GetPaginatedRecipes).Methods("GET")

	// Sort routes
	sort := api.PathPrefix("/sort").Subrouter()
	sort.HandleFunc("/recipes", sortHandler.SortRecipes).Methods("GET")

	// Admin routes (require admin role)
	admin := api.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AuthMiddleware(userService))
	admin.Use(middleware.RoleMiddleware("admin"))
	// Add admin routes here if needed

	// Start server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
