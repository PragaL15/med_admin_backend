package models

import (
	"time"
)

// Record represents a row in your database table.
type Record struct {
	ID          int       `db:"id" json:"id"`
	PID         int       `db:"p_id" json:"p_id"`
	DID         int       `db:"d_id" json:"d_id"`
	CreatedAt   time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `db:"updatedAt" json:"updatedAt"`
	Description string    `db:"Description" json:"description"`
	Precription string `db:"Precription" json:"precription"`
}

type Patient struct {
	ID      int    `db:"id" json:"id"`
	PID     int    `db:"p_id" json:"p_id"`
	PName   string `db:"p_name" json:"p_name"`
	PNumber int64  `db:"p_number" json:"p_number"` // BIGINT in PostgreSQL, matches Go's int64
	PEmail  string `db:"p_email" json:"p_email"`
	PStatus string `db:"p_status" json:"p_status"`
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
