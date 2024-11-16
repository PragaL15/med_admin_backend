package handlers

import (
    "encoding/json"
    "net/http"
    "time"

    models "github.com/PragaL15/med_admin_backend/src/model"
    "gorm.io/gorm"
)

func AddPatient(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        var patient models.Patient

        err := json.NewDecoder(r.Body).Decode(&patient)
        if err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }
        currentTime := time.Now()
        patient.CreatedAt = currentTime
        patient.UpdatedAt = currentTime

        if err := db.Create(&patient).Error; err != nil {
            http.Error(w, "Error inserting new patient", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "message": "Patient created successfully",
            "patient": patient,
        })
    }
}
