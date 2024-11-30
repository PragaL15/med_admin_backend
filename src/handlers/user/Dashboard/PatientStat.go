package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"gorm.io/gorm"
)

type PatientStatusRecord struct {
	PatientID int       `json:"p_id"`
	Month     string    `json:"month"`    // Change to string to hold the month name
	Status    string    `json:"p_status"`
}

func GetPatientStatusForGraph(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var records []PatientStatusRecord

		// Query the record table and join with patient_id to get the status
		err := db.
			Table("record").
			Select("record.p_id, record.date, patient_id.p_status").
			Joins("JOIN patient_id ON record.p_id = patient_id.p_id").
			Find(&records).Error

		if err != nil {
			http.Error(w, "Error fetching patient status data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		for i := range records {
			parsedDate, err := time.Parse("2006-01-02", records[i].Month)
			if err == nil {
				records[i].Month = parsedDate.Format("January")
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(records)
	}
}
