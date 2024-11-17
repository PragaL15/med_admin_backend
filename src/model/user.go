package models

import (
	"time"
)

// Record represents a row in your database table.
type Record struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	PID         int       `gorm:"column:p_id;not null" json:"p_id"`
	DID         int       `gorm:"column:d_id;not null" json:"d_id"`
	Date        time.Time `gorm:"column:date;not null" json:"date"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	Description string    `gorm:"column:description" json:"description"`
	Prescription string   `gorm:"column:prescription" json:"prescription"`
}

func (Record) TableName() string {
	return "record"
}

// Patient represents a patient record.
type Patient struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`                // Unique Primary Key
	PID       uint      `gorm:"column:p_id;not null;uniqueIndex" json:"pid"`       // Unique Patient ID
	Name      string    `gorm:"column:p_name;not null" json:"name"`                // Patient Name
	Phone     string    `gorm:"column:p_number;not null" json:"phone"`             // Phone Number
	Email     string    `gorm:"column:p_email;not null" json:"email"`              // Email Address
	Status    string    `gorm:"column:p_status" json:"status"`                     // Patient Status
	Address   string    `gorm:"column:p_address" json:"address"`                   // Address
	Mode      string    `gorm:"column:p_mode" json:"mode"`                         // Mode of Payment or Consultation
	Age       int       `gorm:"column:p_age;not null" json:"age"`                  // Age
	Gender    string    `gorm:"column:p_gender;not null" json:"gender"`            // Gender
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"` // Created At Timestamp
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"` // Updated At Timestamp
}

func (Patient) TableName() string {
	return "patient_id"
}

// Doctor represents a doctor record.
type Doctor struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	DID       uint      `gorm:"column:d_id;not null;uniqueIndex" json:"d_id"`
	DName     string    `gorm:"column:d_name" json:"d_name"`
	DNumber   int64     `gorm:"column:d_number" json:"d_number"`
	DEmail    string    `gorm:"column:d_email" json:"d_email"`
	DStatus   string    `gorm:"column:d_status" json:"d_status"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Doctor) TableName() string {
	return "doctor_id"
}

// DoctorWithPatients represents a doctor with assigned patients.
type DoctorWithPatients struct {
	DID   uint   `json:"d_id"`
	DName string `json:"d_name"`
	PIDs  []int  `json:"p_ids"`
}

// Appointment represents an appointment record.
type Appointment struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	PID           int       `gorm:"column:p_id;not null" json:"p_id"`
	PName         string    `gorm:"column:p_name" json:"p_name"`
	PNumber       string    `gorm:"column:p_number" json:"p_number"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	AppDate       time.Time `gorm:"column:app_date;not null" json:"app_date"`
	PHealth       string    `gorm:"column:p_health" json:"p_health"`
	DID           int       `gorm:"column:d_id;not null" json:"d_id"`
	Time          string    `gorm:"column:time" json:"time"`
	ProblemHint   string    `gorm:"column:problem_hint" json:"problem_hint"`
	AppoStatus    string    `gorm:"column:appo_status" json:"appo_status"`
}

func (Appointment) TableName() string {
	return "appointments"
}

// AppointmentPost represents an appointment record for creating new appointments.
type AppointmentPost struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	PID         int       `gorm:"column:p_id;not null" json:"p_id"`
	AppDate     string    `gorm:"column:app_date;not null" json:"app_date"` 
	PHealth     string    `gorm:"column:p_health" json:"p_health"`
	DID         int       `gorm:"column:d_id;not null" json:"d_id"`
	Time        string    `gorm:"column:time;not null" json:"time"`       
	ProblemHint string    `gorm:"column:problem_hint" json:"problem_hint"`
	AppoStatus  string    `gorm:"column:appo_status" json:"appo_status"`
}


func (AppointmentPost) TableName() string {
	return "appointments"
}

// Admitted represents a record of admitted patients.
type Admitted struct {
	ID               int       `gorm:"primaryKey;autoIncrement" json:"id"`
	PID              int       `gorm:"column:p_id;not null" json:"p_id"`
	PName            string    `gorm:"column:p_name" json:"p_name"`
	PHealth          string    `gorm:"column:p_health" json:"p_health"`
	POperation       string    `gorm:"column:p_operation" json:"p_operation"`
	POperationDate   time.Time `gorm:"column:p_operation_date" json:"p_operation_date"`
	POperatedDoctor  string    `gorm:"column:p_operated_doctor" json:"p_operated_doctor"`
	DurationAdmit    string    `gorm:"column:duration_admit" json:"duration_admit"`
	WardNo           string    `gorm:"column:ward_no" json:"ward_no"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Admitted) TableName() string {
	return "admitted"
}

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"column:username;not null;unique" json:"username"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	UserID    int       `gorm:"column:user_id;uniqueIndex;not null" json:"user_id"` // Unique user identifier
	Status    int       `gorm:"column:status;default:1" json:"status,omitempty"`    // Default active status
	RoleID    int       `gorm:"column:role_id" json:"role_id,omitempty"`           // Role foreign key
	RoleName  string    `gorm:"column:role_name" json:"role_name,omitempty"`       // Role description
	DID       int       `gorm:"column:d_id" json:"d_id,omitempty"`                 // Doctor ID
	PID       int       `gorm:"column:p_id" json:"p_id,omitempty"`                 // Patient ID
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}



func (User) TableName() string {
	return "user_table"
}

// Route represents an API route.
type Route struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"` 
	Route     string    `gorm:"type:varchar(255);not null"` 
	RoleID    int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"` 
	UpdatedAt time.Time `gorm:"default:current_timestamp on update current_timestamp"`
}

func (Route) TableName() string {
	return "routes"
}

// Role represents a role record.
type Role struct {
	RoleID   int    `gorm:"primaryKey;autoIncrement" json:"role_id"`
	RoleName string `gorm:"column:role_name;not null" json:"role_name"`
}

func (Role) TableName() string {
	return "roles"
}
