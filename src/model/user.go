package models

import (
	"time"
)

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

type Patient struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`               
	PID       uint      `gorm:"column:p_id;autoIncrement;not null;uniqueIndex" json:"p_id"`     
	Name      string    `gorm:"column:p_name;not null" json:"name"`               
	Phone     string    `gorm:"column:p_number;not null" json:"number"`             
	Email     string    `gorm:"column:p_email;not null" json:"email"`              
	Status    string    `gorm:"column:p_status" json:"status"`                     
	Address   string    `gorm:"column:p_address" json:"address"`                   
	Mode      string    `gorm:"column:p_mode" json:"mode"`                        
	Age       int       `gorm:"column:p_age;not null" json:"age"`                  
	Gender    string    `gorm:"column:p_gender;not null" json:"gender"`
	DOB       time.Time   `gorm:"type:date"column:dob;not null" json:"dob"`       
	Occupation string `gorm:"column:occupation;not null" json:"occupation"`     
	Language   string `gorm:"column:lang_spoken;not null" json:"lang_spoken"`
	CreatedAt time.Time `gorm:"column:createdat;autoCreateTime" json:"createdAt"` 
	UpdatedAt time.Time `gorm:"column:updatedat;autoUpdateTime" json:"updatedAt"` 
}
func (Patient) TableName() string {
	return "patient_id"
}

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
type DoctorWithPatients struct {
	DID   uint   `json:"d_id"`
	DName string `json:"d_name"`
	PIDs  []int  `json:"p_ids"`
}
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
	UserID    int       `gorm:"column:user_id;uniqueIndex;not null" json:"user_id"`
	Status    int       `gorm:"column:status;default:1" json:"status,omitempty"`    
	RoleID    int       `gorm:"column:role_id" json:"role_id,omitempty"`         
	RoleName  string    `gorm:"column:role_name" json:"role_name,omitempty"`      
	DID       int       `gorm:"column:d_id" json:"d_id,omitempty"`               
	PID       int       `gorm:"column:p_id" json:"p_id,omitempty"`             
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
func (User) TableName() string {
	return "user_table"
}
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
type Role struct {
	RoleID   int    `gorm:"primaryKey;autoIncrement" json:"role_id"`
	RoleName string `gorm:"column:role_name;not null" json:"role_name"`
}
func (Role) TableName() string {
	return "roles"
}
