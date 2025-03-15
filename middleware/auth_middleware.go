package middleware

import (
	"context"
	"net/http"
	"strings"

	"playground/services"
)

// ContextKey is a type for context keys
type ContextKey string

// Context keys
const (
	UserIDKey ContextKey = "userID"
	UserRoleKey ContextKey = "userRole"
)

// AuthMiddleware creates a middleware that validates JWT tokens
func AuthMiddleware(userService *services.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token from Authorization header
			authorization := r.Header.Get("Authorization")
			if authorization == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Check if the Authorization header has the Bearer prefix
			if !strings.HasPrefix(authorization, "Bearer ") {
				http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
				return
			}

			// Extract the token
			token := strings.TrimPrefix(authorization, "Bearer ")

			// Validate token
			claims, err := userService.ValidateToken(token)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Add user ID and role to context
			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, UserRoleKey, claims.Role)

			// Call next handler with updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RoleMiddleware creates a middleware that checks if the user has the required role
func RoleMiddleware(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get user role from context
			role, ok := r.Context().Value(UserRoleKey).(string)
			if !ok || role != requiredRole {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			// Call next handler
			next.ServeHTTP(w, r)
		})
	}
}