package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
	models "github.com/PragaL15/med_admin_backend/src/model"
)

func CreateAppointment(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, `{"status": false, "message": "Method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}

		var appointment models.AppointmentPost
		if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
			log.Println("Error decoding request body:", err)
			http.Error(w, `{"status": false, "message": "Invalid request body"}`, http.StatusBadRequest)
			return
		}

	
		if appointment.PID == 0 || appointment.DID == 0 || appointment.AppDate == "" || appointment.Time == "" {
			http.Error(w, `{"status": false, "message": "Missing required fields"}`, http.StatusBadRequest)
			return
		}

	
		parsedDate, err := time.Parse("02-01-2006", appointment.AppDate)
		if err != nil {
			log.Println("Error parsing app_date:", err)
			http.Error(w, `{"status": false, "message": "Invalid date format. Expected DD-MM-YYYY"}`, http.StatusBadRequest)
			return
		}

		parsedTime, err := time.Parse("15:04:05", appointment.Time)
		if err != nil {
			log.Println("Error parsing time:", err)
			http.Error(w, `{"status": false, "message": "Invalid time format. Expected HH:mm:ss"}`, http.StatusBadRequest)
			return
		}

		combinedDateTime := time.Date(
			parsedDate.Year(), parsedDate.Month(), parsedDate.Day(),
			parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), 0, time.UTC,
		)

		appointment.AppDate = combinedDateTime.Format("2006-01-02") 
		appointment.Time = combinedDateTime.Format("15:04:05")      

		if err := db.Create(&appointment).Error; err != nil {
			log.Printf("Error creating appointment: %v", err)
			http.Error(w, `{"status": false, "message": "Failed to create appointment"}`, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response := map[string]interface{}{
			"status":  true,
			"message": "Appointment created successfully",
			"data":    appointment,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Error encoding JSON response: %v", err)
			http.Error(w, `{"status": false, "message": "Error sending response"}`, http.StatusInternalServerError)
		}
	}
}
