package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/PragaL15/med_admin_backend/database"
	models "github.com/PragaL15/med_admin_backend/src/model"
	"github.com/gorilla/mux"
)

// CreateDoctor creates a new doctor
func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var doctor models.Doctor
	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the CreatedAt and UpdatedAt fields to the current time
	doctor.CreatedAt = time.Now()
	doctor.UpdatedAt = time.Now()

	query := `
		INSERT INTO doctor_id (d_id, d_name, d_number, d_email, d_status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`
	err := database.DB.QueryRow(context.Background(), query, doctor.DID, doctor.DName, doctor.DNumber, doctor.DEmail, doctor.DStatus, doctor.CreatedAt, doctor.UpdatedAt).Scan(&doctor.ID)
	if err != nil {
		log.Println("Error creating doctor:", err)
		http.Error(w, "Failed to create doctor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(doctor)
}

// GetAllDoctors retrieves all doctors
func GetAllDoctors(w http.ResponseWriter, r *http.Request) {
	var doctors []models.Doctor
	query := `SELECT * FROM doctor_id`
	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		log.Println("Error retrieving doctors:", err)
		http.Error(w, "Failed to retrieve doctors", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var doctor models.Doctor
		if err := rows.Scan(&doctor.ID, &doctor.DID, &doctor.DName, &doctor.DNumber, &doctor.DEmail, &doctor.DStatus, &doctor.CreatedAt, &doctor.UpdatedAt); err != nil {
			log.Println("Error scanning doctor:", err)
			http.Error(w, "Failed to scan doctor", http.StatusInternalServerError)
			return
		}
		doctors = append(doctors, doctor)
	}

	json.NewEncoder(w).Encode(doctors)
}

// GetDoctorByID retrieves a doctor by ID
func GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var doctor models.Doctor

	query := `SELECT * FROM doctor_id WHERE id = $1`
	err := database.DB.QueryRow(context.Background(), query, id).Scan(&doctor.ID, &doctor.DID, &doctor.DName, &doctor.DNumber, &doctor.DEmail, &doctor.DStatus, &doctor.CreatedAt, &doctor.UpdatedAt)
	if err != nil {
		log.Println("Error retrieving doctor:", err)
		http.Error(w, "Doctor not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(doctor)
}

// UpdateDoctor updates an existing doctor
func UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var doctor models.Doctor

	if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update the UpdatedAt field to the current time
	doctor.UpdatedAt = time.Now()

	query := `
		UPDATE doctor_id
		SET d_id = $1, d_name = $2, d_number = $3, d_email = $4, d_status = $5, updated_at = $6
		WHERE id = $7
	`
	_, err := database.DB.Exec(context.Background(), query, doctor.DID, doctor.DName, doctor.DNumber, doctor.DEmail, doctor.DStatus, doctor.UpdatedAt, id)
	if err != nil {
		log.Println("Error updating doctor:", err)
		http.Error(w, "Failed to update doctor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Doctor updated successfully"})
}

// DeleteDoctor deletes a doctor
func DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	query := `DELETE FROM doctor_id WHERE id = $1`
	_, err := database.DB.Exec(context.Background(), query, id)
	if err != nil {
		log.Println("Error deleting doctor:", err)
		http.Error(w, "Failed to delete doctor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Doctor deleted successfully"})
}
