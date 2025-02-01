package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	 "gorm.io/gorm"
	models "github.com/PragaL15/med_admin_backend/src/model"
)

func CreateDoctor(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var doctor models.Doctor
		if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		doctor.CreatedAt = time.Now()
		doctor.UpdatedAt = time.Now()

		if err := db.Create(&doctor).Error; err != nil {
			log.Printf("Error creating doctor: %v", err)
			http.Error(w, "Failed to create doctor", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(doctor)
	}
}

func GetAllDoctors(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var doctors []models.Doctor

		if err := db.Find(&doctors).Error; err != nil {
			log.Printf("Error retrieving doctors: %v", err)
			http.Error(w, "Failed to retrieve doctors", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(doctors)
	}
}
func GetDoctorByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var doctor models.Doctor
		if err := db.Where("id = ?", id).First(&doctor).Error; err != nil {
			if err.Error() == "record not found" {
				http.Error(w, "Doctor not found", http.StatusNotFound)
			} else {
				log.Printf("Error retrieving doctor: %v", err)
				http.Error(w, "Failed to retrieve doctor", http.StatusInternalServerError)
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(doctor)
	}
}
func UpdateDoctor(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var doctor models.Doctor

		if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		doctor.UpdatedAt = time.Now()
		if err := db.Model(&doctor).Where("id = ?", id).Updates(map[string]interface{}{
			"d_id":       doctor.DID,
			"d_name":     doctor.DName,
			"d_number":   doctor.DNumber,
			"d_email":    doctor.DEmail,
			"d_status":   doctor.DStatus,
			"updated_at": doctor.UpdatedAt,
		}).Error; err != nil {
			log.Printf("Error updating doctor: %v", err)
			http.Error(w, "Failed to update doctor", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Doctor updated successfully"})
	}
}
func DeleteDoctor(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		result := db.Where("id = ?", id).Delete(&models.Doctor{})
		if result.Error != nil {
			log.Printf("Error deleting doctor: %v", result.Error)
			http.Error(w, "Failed to delete doctor", http.StatusInternalServerError)
			return
		}
		if result.RowsAffected == 0 {
			http.Error(w, "Doctor not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Doctor deleted successfully"})
	}
}
