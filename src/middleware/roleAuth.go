package middleware

import (
    "net/http"
)

// RoleAuthMiddleware checks if a user has one of the required roles to access a route.
func RoleAuthMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Retrieve role_id from the request context
            roleID, ok := r.Context().Value("role_id").(string)
            if !ok {
                http.Error(w, "Access denied: Role not found", http.StatusForbidden)
                return
            }

            // Check if roleID is in the allowedRoles list
            isAuthorized := false
            for _, role := range allowedRoles {
                if roleID == role {
                    isAuthorized = true
                    break
                }
            }

            // If roleID does not match any allowed role, deny access
            if !isAuthorized {
                http.Error(w, "Access denied: Unauthorized role", http.StatusForbidden)
                return
            }

            // Proceed to the next handler if authorized
            next.ServeHTTP(w, r)
        })
    }
}
