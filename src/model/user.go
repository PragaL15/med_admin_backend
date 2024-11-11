package models

import (
	"time"
)

// Record represents a row in your database table.
type Record struct {
	ID          int       `db:"id" json:"id"`
	PID         int       `db:"p_id" json:"p_id"`
	DID         int       `db:"d_id" json:"d_id"`
	Date time.Time `db:"date" json:"date"`
	CreatedAt   time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `db:"updatedAt" json:"updatedAt"`
	Description string    `db:"Description" json:"description"`
	Prescription string `db:"Prescription" json:"prescription"`
}

type Patient struct {
	ID        int       `json:"id,omitempty"`          // Maps to the 'id' column in the database
	PID       int       `json:"p_id,omitempty"`        // Maps to the 'p_id' column in the database
	Name      string    `json:"p_name"`                // Maps to the 'p_name' column
	Phone     string    `json:"p_number"`              // Maps to the 'p_number' column
	Email     string    `json:"p_email"`               // Maps to the 'p_email' column
	Status    string    `json:"p_status"`              // Maps to the 'p_status' column
	Address   string    `json:"p_address"`             // Maps to the 'p_address' column
	Mode      string    `json:"p_mode"`                // Maps to the 'p_mode' column
	Age       int       `json:"p_age"`                 // Maps to the 'p_age' column
	Gender    string    `json:"p_gender"`              // Maps to the 'p_gender' column
	CreatedAt time.Time `json:"createdat,omitempty"`   // Maps to the 'createdat' column
	UpdatedAt time.Time `json:"updatedat,omitempty"`   // Maps to the 'updatedat' column
}


type Doctor struct {
	ID        int       `db:"id" json:"id"`
	DID       int       `db:"d_id" json:"d_id"` 
	DName     string    `db:"d_name" json:"d_name"`
	DNumber   int64     `db:"d_number" json:"d_number"` // Adjust type if needed
	DEmail    string    `db:"d_email" json:"d_email"`
	DStatus   string    `db:"d_status" json:"d_status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"` // Assuming you added this
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"` // Assuming you added this
}


type Appointment struct {
	ID           int       `json:"id"`
	PID          int       `json:"p_id"`
	PName        string    `json:"p_name"`
	PNumber      string    `json:"p_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	AppDate      time.Time `json:"app_date"`
	PHealth      string    `json:"p_health"`
	DID          int       `json:"d_id"`
	Time         string    `json:"time"`
	ProblemHint  string    `json:"problem_hint"`
	AppoStatus   string    `json:"appo_status"`
}

type Admitted struct {
	ID               int       `json:"id"`
	PID              int       `json:"p_id"`
	PName            string    `json:"p_name"` // This is the patient name from patient_id table
	PHealth          string    `json:"p_health"`
	POperation       string    `json:"p_operation"`
	POperationDate   time.Time `json:"p_operation_date"`
	POperatedDoctor  string    `json:"p_operated_doctor"`
	DurationAdmit    string    `json:"duration_admit"`
	WardNo           string    `json:"ward_no"`
	CreatedAt        time.Time `json:"created_at"`  // Ensure these fields exist if used
	UpdatedAt        time.Time `json:"updated_at"`
	
}
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	UserID    int       `json:"user_id"`   // User ID (may be external or internal)
	Status    int       `json:"status"`    // 1 = Active, 0 = Inactive
	RoleID    int       `json:"role_id"`   // Role ID linking to roles table
	RoleName  string    `json:"role_name"` // Role name retrieved via JOIN with roles table
	CreatedAt time.Time `json:"created_at"`
}

// Role represents the role record in the database.
type Role struct {
	RoleID   int    `json:"role_id"`   // Unique identifier for the role
	RoleName string `json:"role_name"` // The name of the role (e.g., "admin", "user", etc.)
}

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
	RoleID  int    `json:"role_id,omitempty"` // RoleID of the user
	RoleName string `json:"role_name,omitempty"` // RoleName of the user
}

// CreateUserRequest represents the expected JSON payload for creating a new user.
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"` // Role assigned to the new user
	Status   int    `json:"status"`  // User status (1 = Active, 0 = Inactive)
}

// CreateUserResponse represents the response after successfully creating a user.
type CreateUserResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	User    User   `json:"user"`
}