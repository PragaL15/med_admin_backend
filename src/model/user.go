package models

import (
	"time"
)

// Record represents a row in your database table.
type Record struct {
	ID          int       `db:"id" json:"id"`
	PID         int       `db:"p_id" json:"p_id"`
	DID         int       `db:"d_id" json:"d_id"`
	Date        time.Time `db:"date" json:"date"`
	CreatedAt   time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `db:"updatedAt" json:"updatedAt"`
	Description string    `db:"Description" json:"description"`
	Prescription string  `db:"Prescription" json:"prescription"`
}

func (Record) TableName() string {
	return "record" // Ensure the table name matches the database table
}

type Patient struct {
	ID        int       `json:"id"`
  PID uint `gorm:"column:p_id;primaryKey"`
	Name      string    `json:"p_name"`
	Phone     string    `json:"p_number"`
	Email     string    `json:"p_email"`
	Status    string    `json:"p_status"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	Address   string    `json:"p_address"`
	Mode      string    `json:"p_mode"`
	Age       int       `json:"p_age"`
	Gender    string    `json:"p_gender"`
}

func (Patient) TableName() string {
	return "patient_id" // Ensure the table name matches the database table
}

type Doctor struct {
	ID        int       `db:"id" json:"id"`
	DID uint `gorm:"column:d_id;primaryKey"`
	DName     string    `db:"d_name" json:"d_name"`
	DNumber   int64     `db:"d_number" json:"d_number"` 
	DEmail    string    `db:"d_email" json:"d_email"`
	DStatus   string    `db:"d_status" json:"d_status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
type DoctorWithPatients struct {
	DID   uint    `json:"d_id"`
	DName string `json:"d_name"`
	PIDs  []int  `json:"p_ids"`
}


func (Doctor) TableName() string {
	return "doctor_id" // Ensure the table name matches the database table
}

type Appointment struct {
	ID            int       `gorm:"primaryKey"`
	PID           int       `json:"p_id"`
	PName         string    `json:"p_name"`
	PNumber       string    `json:"p_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	AppDate       time.Time `json:"app_date"`
	PHealth       string    `json:"p_health"`
	DID           int       `json:"d_id"`
	Time          string    `json:"time"`
	ProblemHint   string    `json:"problem_hint"`
	AppoStatus    string    `json:"appo_status"`
}
type AppointmentPost struct {
	ID            int       `gorm:"primaryKey"`
	PID           int       `json:"p_id"`
	AppDate       time.Time `json:"app_date"`
	PHealth       string    `json:"p_health"`
	DID           int       `json:"d_id"`
	Time          string    `json:"time"`
	ProblemHint   string    `json:"problem_hint"`
	AppoStatus    string    `json:"appo_status"`
}

func (Appointment) TableName() string {
	return "appointments" // Ensure the table name matches the database table
}
func (AppointmentPost) TableName() string {
	return "appointments" // Ensure the table name
}

type Admitted struct {
	ID               int       `json:"id"`
	PID              int       `json:"p_id"`
	PName            string    `json:"p_name"`
	PHealth          string    `json:"p_health"`
	POperation       string    `json:"p_operation"`
	POperationDate   time.Time `json:"p_operation_date"`
	POperatedDoctor  string    `json:"p_operated_doctor"`
	DurationAdmit    string    `json:"duration_admit"`
	WardNo           string    `json:"ward_no"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (Admitted) TableName() string {
	return "admitted" // Ensure the table name matches the database table
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	UserID    int       `json:"user_id"`
	Status    int       `json:"status"`
	RoleID    int       `json:"role_id"`
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
}

func (User) TableName() string {
	return "user_table" // Ensure the table name matches the database table
}

type Route struct {
	RouteID  int    `json:"route_id"`
	RoutePath string `json:"route_path"`
	UserID   int    `json:"user_id"`
}

func (Route) TableName() string {
	return "routes" // Ensure the table name matches the database table
}

// Role represents the role record in the database.
type Role struct {
	RoleID   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}

func (Role) TableName() string {
	return "roles" // Ensure the table name matches the database table
}

