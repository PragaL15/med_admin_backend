package utils

import (
    "errors"
    "time"
    "github.com/golang-jwt/jwt/v4"
)

// Secret key used to sign JWT tokens. Keep this secure, ideally stored in environment variables.
var jwtSecret = []byte("your_secret_key") // Change this to a strong secret

// GenerateJWT generates a new JWT for a given user ID.
func GenerateJWT(userID int) (string, error) {
    // Set token expiration time (e.g., 1 hour)
    expirationTime := time.Now().Add(time.Hour * 1).Unix()

    // Define claims for the JWT
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     expirationTime,
    }

    // Create a new token and sign it using the HMAC method with the secret key
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// ParseJWT parses and validates a JWT token string, returning the claims if valid.
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        return nil, errors.New("invalid or expired token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, errors.New("could not parse claims")
    }
    return claims, nil
}
