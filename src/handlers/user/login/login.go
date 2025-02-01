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
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Message  string `json:"message"`
	Status   bool   `json:"status"`
	Token    string `json:"token,omitempty"`  
	UserID   int    `json:"user_id,omitempty"` 
	RoleID   int    `json:"role_id,omitempty"` 
	RoleName string `json:"role_name,omitempty"` 
}
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"message":"Invalid request payload","status":false}`, http.StatusBadRequest)
		return
	}
	var user models.User
	err := database.DB.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		http.Error(w, `{"message":"Invalid username or password","status":false}`, http.StatusUnauthorized)
		return
	}

	if user.Status != 1 {
		http.Error(w, `{"message":"Account is inactive","status":false}`, http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		http.Error(w, `{"message":"Invalid username or password","status":false}`, http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.GenerateJWT(user.UserID)
	if err != nil {
		http.Error(w, `{"message":"Could not generate token","status":false}`, http.StatusInternalServerError)
		return
	}

	fmt.Printf("Generated Token: %s\n", tokenString)
	decodedUserID, err := utils.DecodeJWTTokenAndGetUserID(tokenString)
	if err != nil {
		http.Error(w, `{"message":"Error decoding token","status":false}`, http.StatusInternalServerError)
		return
	}

	fmt.Printf("Decoded user_id: %d\n", decodedUserID)

	ctx := context.WithValue(r.Context(), "userID", decodedUserID)
	r = r.WithContext(ctx)

	response := LoginResponse{
		Message:  "Login successful",
		Status:   true,
		Token:    tokenString,
		UserID:   user.UserID,
		RoleID:   user.RoleID,
		RoleName: user.RoleName,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
