package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/PragaL15/med_admin_backend/database"
	models "github.com/PragaL15/med_admin_backend/src/model"
	"github.com/gorilla/mux"
)

// GetRecords retrieves all records.
func GetRecords(w http.ResponseWriter, r *http.Request) {
    var records []models.Record
    query := `SELECT id, p_id, dr_id, createdat, updatedat, description FROM records`

    rows, err := database.DB.Query(context.Background(), query)
    if err != nil {
        http.Error(w, "Failed to fetch records", http.StatusInternalServerError)
        log.Println("Database query error:", err)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var record models.Record
        if err := rows.Scan(&record.ID, &record.PID, &record.DRID, &record.CreatedAt, &record.UpdatedAt, &record.Description); err != nil {
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
    query := `SELECT id, p_id, dr_id, createdat, updatedat, description FROM records WHERE id=$1`
    err = database.DB.QueryRow(context.Background(), query, id).Scan(&record.ID, &record.PID, &record.DRID, &record.CreatedAt, &record.UpdatedAt, &record.Description)
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

    query := `INSERT INTO records (p_id, dr_id, createdat, updatedat, description) VALUES ($1, $2, $3, $4, $5) RETURNING id`
    err := database.DB.QueryRow(context.Background(), query, record.PID, record.DRID, record.CreatedAt, record.UpdatedAt, record.Description).Scan(&record.ID)
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

    query := `UPDATE records SET p_id=$1, dr_id=$2, updatedat=$3, description=$4 WHERE id=$5`
    _, err = database.DB.Exec(context.Background(), query, record.PID, record.DRID, record.UpdatedAt, record.Description, id)
    if err != nil {
        http.Error(w, "Failed to update record", http.StatusInternalServerError)
        log.Println("Record update error:", err)
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

    query := `DELETE FROM records WHERE id=$1`
    _, err = database.DB.Exec(context.Background(), query, id)
    if err != nil {
        http.Error(w, "Failed to delete record", http.StatusInternalServerError)
        log.Println("Record deletion error:", err)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
