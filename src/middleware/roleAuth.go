package middleware

import (
    "net/http"
)

// UserAuthMiddleware checks if a user has one of the required user IDs to access a route.
func RoleAuthMiddleware(allowedUserIDs ...string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Retrieve user_id from the request context
            userID, ok := r.Context().Value("user_id").(string)
            if !ok {
                http.Error(w, "Access denied: User ID not found", http.StatusForbidden)
                return
            }

            // Check if userID is in the allowedUserIDs list
            isAuthorized := false
            for _, id := range allowedUserIDs {
                if userID == id {
                    isAuthorized = true
                    break
                }
            }

            // If userID does not match any allowed ID, deny access
            if !isAuthorized {
                http.Error(w, "Access denied: Unauthorized user", http.StatusForbidden)
                return
            }

            // Proceed to the next handler if authorized
            next.ServeHTTP(w, r)
        })
    }
}
