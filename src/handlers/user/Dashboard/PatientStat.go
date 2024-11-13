package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"gorm.io/gorm"
)

// PatientStatusRecord holds the structure of each record returned for graphing
type PatientStatusRecord struct {
	PatientID int    `json:"p_id"`
	Month     string `json:"month"`
	Status    string `json:"p_status"`
}

// GetPatientStatusForGraph retrieves patient status data for graphing by month
func GetPatientStatusForGraph(db *gorm.DB) http.HandlerFunc {
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

		var records []PatientStatusRecord

		// Use GORM to query the record and join with patient_id for the patient status
		err := db.
			Table("record").
			Select("record.p_id, record.date, patient_id.p_status").
			Joins("JOIN patient_id ON record.p_id = patient_id.p_id").
			Find(&records).Error

		if err != nil {
			http.Error(w, "Error fetching patient status data", http.StatusInternalServerError)
			return
		}

		// Format the date to display only the month
		for i := range records {
			parsedDate, err := time.Parse("2006-01-02", records[i].Month)
			if err == nil {
				records[i].Month = parsedDate.Format("January")
			}
		}

		// Respond with the patient status records as JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(records)
	}
}
