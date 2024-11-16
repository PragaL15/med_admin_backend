package handlers

import (
	"encoding/json"
	"net/http"
	
	models "github.com/PragaL15/med_admin_backend/src/model"
	"gorm.io/gorm"
)

func GetAppointments(db *gorm.DB) http.HandlerFunc {
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

		var appointments []models.Appointment
		err := db.
			Preload("Patient"). 
			Find(&appointments).Error

		if err != nil {
			http.Error(w, "Error fetching appointments", http.StatusInternalServerError)
			return
		}

		for i := range appointments {
			appointments[i].Time = appointments[i].AppDate.Format("03:04 PM")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(appointments)
	}
}
