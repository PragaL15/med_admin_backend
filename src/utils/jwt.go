package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// Secret key used to sign the JWT token (replace with your actual secret key)
var secretKey = []byte("your-secret-key") // Use a more secure key in production

// GenerateJWT generates a new JWT token for the user, containing the user ID.
func GenerateJWT(userID int) (string, error) {
	// Create a new JWT token with user_id and expiration claim
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Expiry time of 1 day
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}

	return tokenString, nil
}

// DecodeJWTTokenAndGetUserID decodes a JWT token and extracts the user ID.
func DecodeJWTTokenAndGetUserID(tokenString string) (int, error) {
	// Parse the JWT token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the token's signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	// Handle errors from parsing the token
	if err != nil {
		return 0, fmt.Errorf("error parsing token: %v", err)
	}

	// Check if the token is valid and extract the claims
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		// Extract the user_id from the claims
		userID, ok := claims["user_id"].(float64) // user_id stored as float64 in JWT claims
		if !ok {
			return 0, fmt.Errorf("user_id not found in token")
		}
		return int(userID), nil
	}

	return 0, fmt.Errorf("invalid token")
}
