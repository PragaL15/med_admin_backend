package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
	models "github.com/PragaL15/med_admin_backend/src/model"
)

// DoctorPatientData represents the combined doctor and patient data for the response
type DoctorPatientData struct {
	Doctors  []models.Doctor  `json:"doctors"`
	Patients []models.Patient `json:"patients"`
}

// GetDoctorsAndPatients retrieves all d_id, d_name from doctor_id table and all p_id, p_name from patient_id table
func GetDoctorsAndPatients(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var doctors []models.Doctor
		var patients []models.Patient

		// Retrieve all d_id and d_name from doctor_id table
		if err := db.Table("doctor_id").Select("d_id", "d_name").Find(&doctors).Error; err != nil {
			log.Printf("Error retrieving doctors: %v", err)
			http.Error(w, "Failed to retrieve doctors", http.StatusInternalServerError)
			return
		}

		// Retrieve all p_id and p_name from patient_id table
		if err := db.Table("patient_id").Select("p_id", "p_name").Find(&patients).Error; err != nil {
			log.Printf("Error retrieving patients: %v", err)
			http.Error(w, "Failed to retrieve patients", http.StatusInternalServerError)
			return
		}

		// Combine the data into a single response structure
		response := DoctorPatientData{
			Doctors:  doctors,
			Patients: patients,
		}

		// Encode and send the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
