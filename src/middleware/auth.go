package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/PragaL15/med_admin_backend/src/utils" 
	"gorm.io/gorm"
)

const (
	ErrUnauthorized   = "Unauthorized"
	ErrInternalServer = "Internal Server Error"
	ErrForbidden      = "Forbidden"
	ErrNotAuthorized  = "Not Authorized"
)

func RoleBasedAccessMiddleware(db *gorm.DB) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, err := getUserIDFromJWT(r)
			if err != nil {
				http.Error(w, ErrUnauthorized, http.StatusUnauthorized)
				return
			}
			log.Printf("User ID from JWT: %d", userID)

			routePath := r.URL.Path

			roleID, err := getUserRoleFromAPI(db, userID, routePath)
			if err != nil {
				log.Printf("Error getting user role from API for user_id: %d and route: %s - %v", userID, routePath, err)
				http.Error(w, ErrInternalServer, http.StatusInternalServerError)
				return
			}
			if !hasPermission(db, roleID, routePath) {
				http.Error(w, ErrNotAuthorized, http.StatusForbidden)
				return
			}
			ctx := context.WithValue(r.Context(), "userID", userID)
			ctx = context.WithValue(ctx, "roleID", roleID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func getUserIDFromJWT(r *http.Request) (int, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Printf("Authorization header is missing")
		return 0, fmt.Errorf("authorization header is missing")
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		log.Printf("Invalid Authorization header format")
		return 0, fmt.Errorf("invalid authorization header format")
	}

	userID, err := utils.DecodeJWTTokenAndGetUserID(tokenParts[1])
	if err != nil {
		log.Printf("Error decoding token: %v", err)
		return 0, fmt.Errorf("error decoding token")
	}

	return userID, nil
}

func getUserRoleFromAPI(db *gorm.DB, userID int, routePath string) (int, error) {
	var permission struct {
		RoleID int `gorm:"column:role_id"`
	}

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

func hasPermission(db *gorm.DB, roleID int, routePath string) bool {
	var permissionCount int64

	err := db.Table("api_permissions").
		Where("route_path = ? AND role_id = ?", routePath, roleID).
		Count(&permissionCount).Error

	if err != nil {
		log.Printf("Failed to check permission for route '%s' with role_id %d. Error: %v", routePath, roleID, err)
		return false
	}

	return permissionCount > 0
}