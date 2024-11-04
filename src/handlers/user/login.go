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
	query := `SELECT id, username, password, status FROM user_table WHERE username = $1`
	err := database.DB.QueryRow(context.Background(), query, req.Username).Scan(&user.ID, &user.Username, &user.Password, &user.Status)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(LoginResponse{Message: "Invalid username or password", Status: false})
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

	// Send a success response if login is successful
	response := LoginResponse{
		Message: "Login successful",
		Status:  true,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HashPassword hashes a password before storing it.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CreateUser creates a new user in the user_table.
func CreateUser(username, password string) (models.User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return models.User{}, err
	}

	query := `INSERT INTO user_table (username, password, status, createdat) VALUES ($1, $2, $3, $4) RETURNING id`
	var user models.User
	err = database.DB.QueryRow(context.Background(), query, username, hashedPassword, true, time.Now()).Scan(&user.ID)
	if err != nil {
		return models.User{}, err
	}

	user.Username = username
	user.Status = true
	user.Password = hashedPassword // Store hashed password in the User struct
	return user, nil
}
