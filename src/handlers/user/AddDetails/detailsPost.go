package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "github.com/PragaL15/med_admin_backend/database"
    models "github.com/PragaL15/med_admin_backend/src/model"
)

// AddPatient handles the creation of a new patient record
func AddPatient(w http.ResponseWriter, r *http.Request) {
    // Enable CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // Handle OPTIONS request for CORS preflight
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    // Restrict to POST method only
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var patient models.Patient

    // Decode JSON body into patient struct
    err := json.NewDecoder(r.Body).Decode(&patient)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Set the timestamps
    currentTime := time.Now()

    // SQL query to insert new patient
    query := `
        INSERT INTO patient_id (p_name, p_number, p_email, p_status, p_address, p_mode, p_age, p_gender, createdat, updatedat)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
        RETURNING id, p_id, createdat, updatedat;
    `

    // Initialize variables for returned values
    var newID, newPID int
    err = database.DB.QueryRow(
        context.Background(),
        query,
        patient.Name, patient.Phone, patient.Email, patient.Status,
        patient.Address, patient.Mode, patient.Age, patient.Gender,
        currentTime, currentTime,
    ).Scan(&newID, &newPID, &patient.CreatedAt, &patient.UpdatedAt)

    if err != nil {
        http.Error(w, "Error inserting new patient", http.StatusInternalServerError)
        return
    }

    // Set the patient ID and PID with returned values
    patient.ID = newID
    patient.PID = newPID

    // Respond with the new patient data as JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(patient)
}
