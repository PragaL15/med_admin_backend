package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
	models "github.com/PragaL15/med_admin_backend/src/model"
)

func RecentOperation(db *gorm.DB) http.HandlerFunc {
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

		var admittedRecords []models.Admitted

		err := db.
			Table("admitted").
			Select("admitted.id, admitted.p_id, patient_id.p_name, admitted.p_health, admitted.p_operation, admitted.p_operation_date, admitted.p_operated_doctor, admitted.duration_admit, admitted.ward_no").
			Joins("JOIN patient_id ON admitted.p_id = patient_id.p_id").
			Find(&admittedRecords).Error

		if err != nil {
			log.Printf("Error executing query: %v", err) 
			http.Error(w, "Error fetching admitted records", http.StatusInternalServerError)
			return
		}

		if len(admittedRecords) == 0 {
			http.Error(w, "No admitted records found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(admittedRecords)
		if err != nil {
			log.Printf("Error encoding JSON response: %v", err)
			http.Error(w, "Error sending response", http.StatusInternalServerError)
		}
	}
}
