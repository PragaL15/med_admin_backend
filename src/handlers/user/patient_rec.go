package handlers

import (
    "context"
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "time"

    "github.com/PragaL15/med_admin_backend/database"
    "github.com/gorilla/mux"
    models "github.com/PragaL15/med_admin_backend/src/model"
)

// CreatePatient inserts a new patient record into the database.
func CreatePatient(w http.ResponseWriter, r *http.Request) {
    var patient models.Patient
    if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Set CreatedAt and UpdatedAt to the current time
    patient.CreatedAt = time.Now()
    patient.UpdatedAt = time.Now()

    query := `
        INSERT INTO patient_id (p_id, p_name, p_number, p_email, p_status, createdat, updatedat)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id
    `

    ctx := context.Background()
    err := database.DB.QueryRow(ctx, query, patient.PID, patient.Name, patient.Phone, patient.Email, patient.Status, patient.CreatedAt, patient.UpdatedAt).Scan(&patient.PID)
    if err != nil {
        log.Println("Error creating patient:", err)
        http.Error(w, "Failed to create patient", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(patient)
}

// GetAllPatients retrieves all patient records from the database.
func GetAllPatients(w http.ResponseWriter, r *http.Request) {
    var patients []models.Patient
    query := `SELECT id, p_id, p_name, p_number, p_email, p_status, createdat, updatedat FROM patient_id`

    ctx := context.Background()
    rows, err := database.DB.Query(ctx, query)
    if err != nil {
        log.Println("Error retrieving patients:", err)
        http.Error(w, "Failed to retrieve patients", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var patient models.Patient
        if err := rows.Scan(&patient.PID, &patient.PID, &patient.Name, &patient.Phone, &patient.Email, &patient.Status, &patient.CreatedAt, &patient.UpdatedAt); err != nil {
            log.Println("Error scanning patient:", err)
            http.Error(w, "Failed to retrieve patients", http.StatusInternalServerError)
            return
        }
        patients = append(patients, patient)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(patients)
}

// GetPatientByID retrieves a single patient by ID from the database.
func GetPatientByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid patient ID", http.StatusBadRequest)
        return
    }

    var patient models.Patient
    query := `SELECT id, p_id, p_name, p_number, p_email, p_status, createdat, updatedat FROM patient_id WHERE id = $1`
    ctx := context.Background()
    err = database.DB.QueryRow(ctx, query, id).Scan(&patient.PID, &patient.PID, &patient.Name, &patient.Phone, &patient.Email, &patient.Status, &patient.CreatedAt, &patient.UpdatedAt)
    if err == sql.ErrNoRows {
        http.Error(w, "Patient not found", http.StatusNotFound)
        return
    } else if err != nil {
        log.Println("Error retrieving patient:", err)
        http.Error(w, "Failed to retrieve patient", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(patient)
}

// UpdatePatient updates an existing patient record in the database.
func UpdatePatient(w http.ResponseWriter, r *http.Request) {
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

    query := `
        UPDATE patient_id
        SET p_id = $1, p_name = $2, p_number = $3, p_email = $4, p_status = $5, updatedat = $6
        WHERE id = $7
    `
    ctx := context.Background()
    _, err = database.DB.Exec(ctx, query, patient.PID, patient.Name, patient.Phone, patient.Email, patient.Status, patient.UpdatedAt, id)
    if err != nil {
        log.Println("Error updating patient:", err)
        http.Error(w, "Failed to update patient", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Patient updated successfully"})
}

// DeletePatient deletes a patient record from the database.
func DeletePatient(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid patient ID", http.StatusBadRequest)
        return
    }

    query := `DELETE FROM patient_id WHERE id = $1`
    ctx := context.Background()
    _, err = database.DB.Exec(ctx, query, id)
    if err != nil {
        log.Println("Error deleting patient:", err)
        http.Error(w, "Failed to delete patient", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Patient deleted successfully"})
}
