package middleware

import (
    "context"
    "net/http"
    "strings"

    "github.com/PragaL15/med_admin_backend/src/utils"
)

// JWTAuthMiddleware checks the token's validity and extracts user information.
func JWTAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract the token from the Authorization header.
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization token is required", http.StatusUnauthorized)
            return
        }

        // Remove "Bearer " prefix if present.
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        // Parse the token using utils.ParseJWT.
        claims, err := utils.ParseJWT(tokenString)
        if err != nil {
            http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }

        // Retrieve user_id and role_id from token claims.
        userID, okUser := claims["user_id"].(string)
        roleID, okRole := claims["role_id"].(string)
        if !okUser || !okRole {
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
            return
        }

        // Set user_id and role_id in request context.
        ctx := context.WithValue(r.Context(), "user_id", userID)
        ctx = context.WithValue(ctx, "role_id", roleID)

        // Proceed to the next handler with updated context.
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
