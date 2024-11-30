package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)
var secretKey = []byte("your-secret-key") 
func GenerateJWT(userID int) (string, error) {
claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}

	return tokenString, nil
}
func DecodeJWTTokenAndGetUserID(tokenString string) (int, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return 0, fmt.Errorf("error parsing token: %v", err)
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userID, ok := claims["user_id"].(float64) 
		if !ok {
			return 0, fmt.Errorf("user_id not found in token")
		}
		return int(userID), nil
	}
	return 0, fmt.Errorf("invalid token")
}
