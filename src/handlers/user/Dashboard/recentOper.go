package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
	models "github.com/PragaL15/med_admin_backend/src/model"
)

// RecentOperation handles fetching recent operation records with patient data
func RecentOperation(db *gorm.DB) http.HandlerFunc {
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

		// Declare a slice to hold the recent operation records
		var admittedRecords []models.Admitted

		// Use GORM to join the 'admitted' table with 'patient_id' and fetch required fields
		err := db.
			Table("admitted").
			Select("admitted.id, admitted.p_id, patient_id.p_name, admitted.p_health, admitted.p_operation, admitted.p_operation_date, admitted.p_operated_doctor, admitted.duration_admit, admitted.ward_no").
			Joins("JOIN patient_id ON admitted.p_id = patient_id.p_id").
			Find(&admittedRecords).Error

		if err != nil {
			log.Printf("Error executing query: %v", err) // Log the error details
			http.Error(w, "Error fetching admitted records", http.StatusInternalServerError)
			return
		}

		// Check if no records were found
		if len(admittedRecords) == 0 {
			http.Error(w, "No admitted records found", http.StatusNotFound)
			return
		}

		// Send the fetched records as a JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(admittedRecords)
		if err != nil {
			log.Printf("Error encoding JSON response: %v", err) // Log the error details
			http.Error(w, "Error sending response", http.StatusInternalServerError)
		}
	}
}
