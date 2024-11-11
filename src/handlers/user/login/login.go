package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/PragaL15/med_admin_backend/database"
	models "github.com/PragaL15/med_admin_backend/src/model"
	"golang.org/x/crypto/bcrypt"
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
	UserID  int    `json:"user_id,omitempty"` // UserID only included on successful login
}

// Login authenticates a user based on username and password.
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(LoginResponse{Message: "Invalid request payload", Status: false})
		return
	}

	// Query to retrieve user data by username from user_table
	var user models.User
	query := `SELECT id, username, password, user_id, status FROM user_table WHERE username = $1`
	err := database.DB.QueryRow(context.Background(), query, req.Username).Scan(&user.ID, &user.Username, &user.Password, &user.UserID, &user.Status)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(LoginResponse{Message: "Invalid username or password", Status: false})
		return
	}

	// Check if the user is active (status = 1)
	if user.Status != 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(LoginResponse{Message: "Account is inactive", Status: false})
		return
	}

	// Check if the password matches the hashed password in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(LoginResponse{Message: "Invalid username or password", Status: false})
		return
	}

	// Return successful login response with user_id
	response := LoginResponse{
		Message: "Login successful",
		Status:  true,
		UserID:  user.UserID,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HashPassword hashes a plain password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CreateUser creates a new user in the user_table.
func CreateUser(username, password string, userID int, status int) (models.User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return models.User{}, err
	}

	// Insert the new user into the user_table
	query := `INSERT INTO user_table (username, password, user_id, status, createdat) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var user models.User
	err = database.DB.QueryRow(context.Background(), query, username, hashedPassword, userID, status, time.Now()).Scan(&user.ID)
	if err != nil {
		return models.User{}, err
	}

	// Populate the returned user struct
	user.Username = username
	user.Password = hashedPassword
	user.UserID = userID
	user.Status = status
	return user, nil
}
