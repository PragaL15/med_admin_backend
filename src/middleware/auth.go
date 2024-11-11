package middleware

import (
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt"
    "github.com/PragaL15/med_admin_backend/utils"
    "strings"
    "time"
)

// JWTAuthMiddleware checks the token's validity and extracts user information
func JWTAuthMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        token := c.Get("Authorization")
        if token == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Authorization token is required",
            })
        }

        // Parse JWT token
        claims, err := utils.ParseJWT(strings.TrimPrefix(token, "Bearer "))
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Invalid or expired token",
            })
        }

        // Store user ID and role from the JWT claims
        c.Locals("user_id", claims["user_id"])
        c.Locals("role_id", claims["role_id"])

        return c.Next()
    }
}