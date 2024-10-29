package models

import (
	"time"
)

// Record represents a row in your database table.
type Record struct {
	ID          int       `db:"id"`          // Unique identifier for the record
	PID         string    `db:"p_id"`        // Patient ID
	DRID        string    `db:"dr_id"`       // Doctor ID
	CreatedAt   time.Time `db:"createdat"`   // Timestamp for when the record was created
	UpdatedAt   time.Time `db:"updatedat"`   // Timestamp for when the record was last updated
	Description string    `db:"description"`  // Description of the record
}
