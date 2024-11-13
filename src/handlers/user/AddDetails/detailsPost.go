package handlers

import (
    "encoding/json"
    "net/http"
    "time"

    models "github.com/PragaL15/med_admin_backend/src/model"
    "gorm.io/gorm"
)

// AddPatient handles the creation of a new patient record
func AddPatient(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
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

        // Set patient creation and update time
        patient.CreatedAt = currentTime
        patient.UpdatedAt = currentTime

        // Insert the new patient record using GORM
        if err := db.Create(&patient).Error; err != nil {
            http.Error(w, "Error inserting new patient", http.StatusInternalServerError)
            return
        }

        // Respond with the new patient data as JSON
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "message": "Patient created successfully",
            "patient": patient,
        })
    }
}
