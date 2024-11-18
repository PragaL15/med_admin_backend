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
        if r.Header.Get("Content-Type") != "application/json" {
            http.Error(w, "Invalid Content-Type, expected application/json", http.StatusUnsupportedMediaType)
            return
        }

        var requestData struct {
            PID     uint   `json:"p_id"`
            Name    string `json:"name"`
            Phone   string `json:"phone"`
            Email   string `json:"email"`
            Status  string `json:"status"`
            Address string `json:"address"`
            Mode    string `json:"mode"`
            Age     int    `json:"age"`
            Gender  string `json:"gender"`
        }

        if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        if requestData.Name == "" || requestData.Phone == "" || requestData.Age <= 0 || requestData.Gender == "" {
            http.Error(w, "Missing or invalid required fields", http.StatusBadRequest)
            return
        }

        patient := models.Patient{
            PID:       requestData.PID,
            Name:      requestData.Name,
            Phone:     requestData.Phone,
            Email:     requestData.Email,
            Status:    requestData.Status,
            Address:   requestData.Address,
            Mode:      requestData.Mode,
            Age:       requestData.Age,
            Gender:    requestData.Gender,
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        }

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

