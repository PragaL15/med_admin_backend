package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/PragaL15/med_admin_backend/database"
	models "github.com/PragaL15/med_admin_backend/src/model"
	"github.com/PragaL15/med_admin_backend/src/utils"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"context"
)

// LoginRequest represents the expected JSON payload for login.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the response after a successful or failed login attempt.
type LoginResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Token   string `json:"token,omitempty"`  // Token only included on successful login
	UserID  int    `json:"user_id,omitempty"` // UserID only included on successful login
	RoleID  int    `json:"role_id,omitempty"` // RoleID of the user
	RoleName string `json:"role_name,omitempty"` // RoleName of the user
}

// Login authenticates a user and issues a JWT if valid.
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"message":"Invalid request payload","status":false}`, http.StatusBadRequest)
		return
	}

	// Retrieve user data by username using GORM
	var user models.User
	err := database.DB.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		http.Error(w, `{"message":"Invalid username or password","status":false}`, http.StatusUnauthorized)
		return
	}

	// Check if the user's account is active
	if user.Status != 1 {
		http.Error(w, `{"message":"Account is inactive","status":false}`, http.StatusUnauthorized)
		return
	}

	// Verify the password using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		http.Error(w, `{"message":"Invalid username or password","status":false}`, http.StatusUnauthorized)
		return
	}

	// Generate a JWT for the authenticated user
	tokenString, err := utils.GenerateJWT(user.UserID)
	if err != nil {
		http.Error(w, `{"message":"Could not generate token","status":false}`, http.StatusInternalServerError)
		return
	}

	// Decode the token to get the user_id (this is the key part)
	decodedUserID, err := utils.DecodeJWTTokenAndGetUserID(tokenString)
	if err != nil {
		http.Error(w, `{"message":"Error decoding token","status":false}`, http.StatusInternalServerError)
		return
	}

	// Print the user_id to the terminal (this is what you want)
	fmt.Printf("Decoded user_id: %d\n", decodedUserID)

	// Store user_id in the request context for use in middlewares and other handlers
	ctx := context.WithValue(r.Context(), "userID", decodedUserID)
	r = r.WithContext(ctx)

	// Respond with the token, user ID, and role information
	response := LoginResponse{
		Message: "Login successful",
		Status:  true,
		Token:   tokenString,
		UserID:  user.UserID,
		RoleID:  user.RoleID,
		RoleName: user.RoleName,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
