package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/PragaL15/med_admin_backend/src/utils" // Assuming GenerateJWT and DecodeJWT are in utils
	"gorm.io/gorm"
)

const (
	ErrUnauthorized   = "Unauthorized"
	ErrInternalServer = "Internal Server Error"
	ErrForbidden      = "Forbidden"
	ErrNotAuthorized  = "Not Authorized"
)

// RoleBasedAccessMiddleware ensures the user has the proper role to access the route
func RoleBasedAccessMiddleware(db *gorm.DB) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Extract the user_id from the JWT token in the Authorization header
			userID, err := getUserIDFromJWT(r)
			if err != nil {
				http.Error(w, ErrUnauthorized, http.StatusUnauthorized)
				return
			}

			// Print user_id to confirm it's extracted correctly
			log.Printf("User ID from JWT: %d", userID)

			// 2. Get the route path the user is trying to access
			routePath := r.URL.Path

			// 3. Get the role_id based on the user_id and route_path
			roleID, err := getUserRoleFromAPI(db, userID, routePath)
			if err != nil {
				log.Printf("Error getting user role from API for user_id: %d and route: %s - %v", userID, routePath, err)
				http.Error(w, ErrInternalServer, http.StatusInternalServerError)
				return
			}

			// 4. Check if the user has permission for the route
			if !hasPermission(db, roleID, routePath) {
				http.Error(w, ErrNotAuthorized, http.StatusForbidden)
				return
			}

			// 5. If all checks pass, proceed to the next handler
			ctx := context.WithValue(r.Context(), "userID", userID)
			ctx = context.WithValue(ctx, "roleID", roleID)
			r = r.WithContext(ctx)

			// Continue processing the request by passing it to the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// Helper function to extract userID from JWT token in the Authorization header
func getUserIDFromJWT(r *http.Request) (int, error) {
	// Extract the token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Printf("Authorization header is missing")
		return 0, fmt.Errorf("authorization header is missing")
	}

	// The Authorization header should be in the format: "Bearer <token>"
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		log.Printf("Invalid Authorization header format")
		return 0, fmt.Errorf("invalid authorization header format")
	}

	// Decode the token to extract the user_id
	userID, err := utils.DecodeJWTTokenAndGetUserID(tokenParts[1])
	if err != nil {
		log.Printf("Error decoding token: %v", err)
		return 0, fmt.Errorf("error decoding token")
	}

	return userID, nil
}

// Helper function to get the role_id based on user_id and route_path
func getUserRoleFromAPI(db *gorm.DB, userID int, routePath string) (int, error) {
	var permission struct {
		RoleID int `gorm:"column:role_id"`
	}

	// Query the api_permissions table to get the role_id for the user and route
	err := db.Table("api_permissions").
		Select("api_permissions.role_id").
		Joins("JOIN user_roles ON user_roles.role_id = api_permissions.role_id").
		Joins("JOIN user_table ON user_table.user_id = user_roles.user_id").
		Where("user_table.user_id = ? AND api_permissions.route_path = ?", userID, routePath).
		First(&permission).Error

	if err != nil {
		log.Printf("Failed query: SELECT role_id FROM api_permissions JOIN user_roles ON user_roles.role_id = api_permissions.role_id JOIN user_table ON user_table.user_id = user_roles.user_id WHERE user_table.user_id = %d AND api_permissions.route_path = '%s'. Error: %v", userID, routePath, err)
		return 0, err
	}
	return permission.RoleID, nil
}

// Helper function to check if the user has permission to access the route
func hasPermission(db *gorm.DB, roleID int, routePath string) bool {
	var permissionCount int64

	// Query the api_permissions table to check if the route exists for this role
	err := db.Table("api_permissions").
		Where("route_path = ? AND role_id = ?", routePath, roleID).
		Count(&permissionCount).Error

	if err != nil {
		log.Printf("Failed to check permission for route '%s' with role_id %d. Error: %v", routePath, roleID, err)
		return false
	}

	// If the count is greater than 0, the user has permission for the route
	return permissionCount > 0
}
