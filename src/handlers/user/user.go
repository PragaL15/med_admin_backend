package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PragaL15/med_admin_backend/database"
	models "github.com/PragaL15/med_admin_backend/src/model"
	"github.com/gorilla/mux"
)

// GetRecords retrieves all records.
func GetRecords(w http.ResponseWriter, r *http.Request) {
	var records []models.Record
	query := `SELECT id, p_id, d_id, date, createdat, updatedat, description, prescription FROM record`

	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		http.Error(w, "Failed to fetch records", http.StatusInternalServerError)
		log.Println("Database query error:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var record models.Record
		if err := rows.Scan(&record.ID, &record.PID, &record.DID, &record.Date, &record.CreatedAt, &record.UpdatedAt, &record.Description, &record.Prescription); err != nil {
			http.Error(w, "Failed to parse record data", http.StatusInternalServerError)
			log.Println("Row scan error:", err)
			return
		}
		records = append(records, record)
	}
	if rows.Err() != nil {
		http.Error(w, "Error in rows result set", http.StatusInternalServerError)
		log.Println("Rows iteration error:", rows.Err())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

// GetRecordByID retrieves a record by ID.
func GetRecordByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var record models.Record
	query := `SELECT id, p_id, d_id, date, createdat, updatedat, description, prescription FROM record WHERE id=$1`
	err = database.DB.QueryRow(context.Background(), query, id).Scan(&record.ID, &record.PID, &record.DID, &record.Date, &record.CreatedAt, &record.UpdatedAt, &record.Description, &record.Prescription)
	if err != nil {
		http.Error(w, "Record not found", http.StatusNotFound)
		log.Println("Record fetch error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

// CreateRecord creates a new record.
func CreateRecord(w http.ResponseWriter, r *http.Request) {
	var record models.Record
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO record (p_id, d_id, date, createdat, updatedat, description, prescription) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := database.DB.QueryRow(context.Background(), query, record.PID, record.DID, record.Date, time.Now(), time.Now(), record.Description, record.Prescription).Scan(&record.ID)
	if err != nil {
		http.Error(w, "Failed to create record", http.StatusInternalServerError)
		log.Println("Record creation error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

// UpdateRecord updates an existing record by ID.
func UpdateRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var record models.Record
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `UPDATE record SET p_id=$1, d_id=$2, date=$3, updatedat=$4, description=$5, prescription=$6 WHERE id=$7`
	_, err = database.DB.Exec(context.Background(), query, record.PID, record.DID, record.Date, time.Now(), record.Description, record.Prescription, id)
	if err != nil {
		http.Error(w, "Failed to update record", http.StatusInternalServerError)
		log.Println("Record update error:", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdateDescriptionByPID updates only the description field for a specific p_id.
func UpdateDescriptionByPID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, err := strconv.Atoi(vars["p_id"])
	if err != nil {
		http.Error(w, "Invalid Patient ID", http.StatusBadRequest)
		return
	}

	var data struct {
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `UPDATE record SET description=$1, updatedat=$2 WHERE p_id=$3`
	_, err = database.DB.Exec(context.Background(), query, data.Description, time.Now(), pID)
	if err != nil {
		http.Error(w, "Failed to update description", http.StatusInternalServerError)
		log.Println("Description update error:", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdatePrescription updates only the prescription for multiple IDs.
func UpdatePrescription(w http.ResponseWriter, r *http.Request) {
	type UpdateData struct {
		IDs          []int  `json:"ids"`
		Prescription string `json:"prescription"`
	}
	var data UpdateData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `UPDATE record SET prescription=$1, updatedat=$2 WHERE id=ANY($3)`
	_, err := database.DB.Exec(context.Background(), query, data.Prescription, time.Now(), data.IDs)
	if err != nil {
		http.Error(w, "Failed to update prescription", http.StatusInternalServerError)
		log.Println("Prescription update error:", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteRecord deletes a record by ID.
func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM record WHERE id=$1`
	_, err = database.DB.Exec(context.Background(), query, id)
	if err != nil {
		http.Error(w, "Failed to delete record", http.StatusInternalServerError)
		log.Println("Record deletion error:", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
