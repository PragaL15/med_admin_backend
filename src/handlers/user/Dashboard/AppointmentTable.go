package handlers

import (
	"encoding/json"
	"net/http"
	
	models "github.com/PragaL15/med_admin_backend/src/model"
	"gorm.io/gorm"
)

// GetAppointments handles fetching all appointment records
func GetAppointments(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle OPTIONS request for CORS preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Restrict to GET method only
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var appointments []models.Appointment

		// Use GORM to fetch appointments and associated patient details
		err := db.
			Preload("Patient"). // Preload the associated patient data (optional, if you have a Patient relation in the model)
			Find(&appointments).Error

		if err != nil {
			http.Error(w, "Error fetching appointments", http.StatusInternalServerError)
			return
		}

		// Format the time fields (if needed) and prepare the response
		for i := range appointments {
			// Format dbTime to 12-hour format (if dbTime is stored as time)
			appointments[i].Time = appointments[i].AppDate.Format("03:04 PM")
		}

		// Respond with the fetched appointment data as JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(appointments)
	}
}
