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
	ID        int       `json:"id"`
	PID       int       `json:"p_id"`
	PName     string    `json:"p_name"`
	PNumber   string    `json:"p_number"` // Assuming this is a string in the DB
	PEmail    string    `json:"p_email"`
	PStatus   string    `json:"p_status"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
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