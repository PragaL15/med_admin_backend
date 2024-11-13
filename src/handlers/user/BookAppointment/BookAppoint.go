package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"gorm.io/gorm"
	models "github.com/PragaL15/med_admin_backend/src/model"
)

// CreateAppointment handles creating new appointment records
func CreateAppointment(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle OPTIONS request for CORS preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Restrict to POST method only
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the JSON request body into an Appointment struct
		var appointment models.AppointmentPost
		if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
			log.Println("Error decoding request body:", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate required fields
		if appointment.PID == 0 || appointment.DID == 0 || appointment.AppDate.IsZero() {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		// Insert the appointment into the database
		if err := db.Create(&appointment).Error; err != nil {
			log.Printf("Error creating appointment: %v", err)
			http.Error(w, "Failed to create appointment", http.StatusInternalServerError)
			return
		}

		// Send a success response with the created appointment
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(appointment); err != nil {
			log.Printf("Error encoding JSON response: %v", err)
			http.Error(w, "Error sending response", http.StatusInternalServerError)
		}
	}
}
