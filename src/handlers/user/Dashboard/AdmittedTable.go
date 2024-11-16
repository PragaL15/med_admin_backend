package handlers

import (
	"encoding/json"
	"net/http"
	models "github.com/PragaL15/med_admin_backend/src/model"
	"gorm.io/gorm"
)

func GetAdmitted(db *gorm.DB) http.HandlerFunc {
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
			Preload("Patient"). 
			Joins("JOIN patients ON admitted.p_id = patients.p_id"). 
			Select("admitted.id, admitted.p_id, patients.p_name, admitted.p_health, admitted.p_operation, admitted.p_operation_date, admitted.p_operated_doctor, admitted.duration_admit, admitted.ward_no").
			Find(&admittedRecords).Error

		if err != nil {
			http.Error(w, "Error fetching admitted records", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(admittedRecords)
	}
}
