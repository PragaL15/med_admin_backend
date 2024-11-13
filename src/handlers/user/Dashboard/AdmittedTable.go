package handlers

import (
	"encoding/json"
	"net/http"
	models "github.com/PragaL15/med_admin_backend/src/model"
	"gorm.io/gorm"
)

// GetAdmitted handles fetching all records from the admitted table
func GetAdmitted(db *gorm.DB) http.HandlerFunc {
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

		// Declare a variable to store admitted records
		var admittedRecords []models.Admitted

		// Use GORM to fetch admitted records with related patient data using a JOIN query
		err := db.
			Preload("Patient"). // Preload related Patient data if needed
			Joins("JOIN patients ON admitted.p_id = patients.p_id"). // Corrected JOIN query
			Select("admitted.id, admitted.p_id, patients.p_name, admitted.p_health, admitted.p_operation, admitted.p_operation_date, admitted.p_operated_doctor, admitted.duration_admit, admitted.ward_no").
			Find(&admittedRecords).Error

		if err != nil {
			http.Error(w, "Error fetching admitted records", http.StatusInternalServerError)
			return
		}

		// Send the records as JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(admittedRecords)
	}
}
