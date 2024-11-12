package middleware

import (
    "context"
    "net/http"
    "strconv"
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

        // Retrieve user_id from token claims.
        var userID string
        if uid, ok := claims["user_id"].(float64); ok {
            userID = strconv.Itoa(int(uid))
        } else if uidStr, ok := claims["user_id"].(string); ok {
            userID = uidStr
        } else {
            http.Error(w, "Invalid token claims: user_id", http.StatusUnauthorized)
            return
        }

        // Set user_id in request context.
        ctx := context.WithValue(r.Context(), "user_id", userID)

        // Proceed to the next handler with updated context.
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
