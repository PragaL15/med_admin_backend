package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
	models "github.com/PragaL15/med_admin_backend/src/model"
    "gorm.io/gorm"

)

// CreatePatient inserts a new patient record into the database.
func CreatePatient(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var patient models.Patient
		if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Set CreatedAt and UpdatedAt to the current time
		patient.CreatedAt = time.Now()
		patient.UpdatedAt = time.Now()

		// Create the patient record using GORM
		if err := db.Create(&patient).Error; err != nil {
			log.Println("Error creating patient:", err)
			http.Error(w, "Failed to create patient", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(patient)
	}
}

// GetAllPatients retrieves all patient records from the database.
func GetAllPatients(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var patients []models.Patient

		// Retrieve all patient records using GORM
		if err := db.Find(&patients).Error; err != nil {
			log.Println("Error retrieving patients:", err)
			http.Error(w, "Failed to retrieve patients", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(patients)
	}
}

// GetPatientByID retrieves a single patient by ID from the database.
func GetPatientByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid patient ID", http.StatusBadRequest)
			return
		}

		var patient models.Patient
		// Retrieve a patient by ID using GORM
		if err := db.First(&patient, id).Error; err != nil {
			if err.Error() == "record not found" {
				http.Error(w, "Patient not found", http.StatusNotFound)
			} else {
				log.Println("Error retrieving patient:", err)
				http.Error(w, "Failed to retrieve patient", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(patient)
	}
}

// UpdatePatient updates an existing patient record in the database.
func UpdatePatient(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid patient ID", http.StatusBadRequest)
			return
		}

		var patient models.Patient
		if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Set UpdatedAt to the current time
		patient.UpdatedAt = time.Now()

		// Update the patient record using GORM
		if err := db.Model(&patient).Where("id = ?", id).Updates(map[string]interface{}{
			"p_id":       patient.PID,
			"p_name":     patient.Name,
			"p_number":   patient.Phone,
			"p_email":    patient.Email,
			"p_status":   patient.Status,
			"updatedat":  patient.UpdatedAt,
			"p_address":  patient.Address,
			"p_mode":     patient.Mode,
			"p_age":      patient.Age,
			"p_gender":   patient.Gender,
		}).Error; err != nil {
			log.Println("Error updating patient:", err)
			http.Error(w, "Failed to update patient", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Patient updated successfully"})
	}
}

// DeletePatient deletes a patient record from the database.
func DeletePatient(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid patient ID", http.StatusBadRequest)
			return
		}

		// Delete the patient record using GORM
		result := db.Where("id = ?", id).Delete(&models.Patient{})
		if result.Error != nil {
			log.Println("Error deleting patient:", result.Error)
			http.Error(w, "Failed to delete patient", http.StatusInternalServerError)
			return
		}

		// Check if any rows were deleted
		if result.RowsAffected == 0 {
			http.Error(w, "Patient not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Patient deleted successfully"})
	}
}