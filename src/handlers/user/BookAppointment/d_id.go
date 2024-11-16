package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
	models "github.com/PragaL15/med_admin_backend/src/model"
)

type DoctorPatientData struct {
	Doctors  []models.Doctor  `json:"doctors"`
	Patients []models.Patient `json:"patients"`
}
func GetDoctorsAndPatients(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var doctors []models.Doctor
		var patients []models.Patient

		if err := db.Table("doctor_id").Select("d_id", "d_name").Find(&doctors).Error; err != nil {
			log.Printf("Error retrieving doctors: %v", err)
			http.Error(w, "Failed to retrieve doctors", http.StatusInternalServerError)
			return
		}

		if err := db.Table("patient_id").Select("p_id", "p_name").Find(&patients).Error; err != nil {
			log.Printf("Error retrieving patients: %v", err)
			http.Error(w, "Failed to retrieve patients", http.StatusInternalServerError)
			return
		}

		response := DoctorPatientData{
			Doctors:  doctors,
			Patients: patients,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
