package addDetailsHandlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	models "github.com/PragaL15/med_admin_backend/src/model"
	"gorm.io/gorm"
)

func AddPatient(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
            PID      uint `json:"pid"`
			Name       string    `json:"name"`
			Phone      string    `json:"number"`
			Email      string    `json:"email"`
			Address    string    `json:"address"`
			Age        int       `json:"age"`
			Gender     string    `json:"gender"`
			Occupation string    `json:"occupation"`
			Language   string    `json:"lang_spoken"`
			DOB        string    `json:"dob"`
			CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
			UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
		}

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&input); err != nil {
			log.Println("Error decoding request body:", err)
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		parsedDOB, err := time.Parse("2006-01-02", input.DOB)
		if err != nil {
			log.Println("Error parsing DOB:", err)
			http.Error(w, "Invalid date format. Use YYYY-MM-DD.", http.StatusBadRequest)
			return
		}

		patient := models.Patient{
            PID:       input.PID,
			Name:      input.Name,
			Phone:     input.Phone,
			Email:     input.Email,
			Address:   input.Address,
			Age:       input.Age,
			Gender:    input.Gender,
			DOB:       parsedDOB,
            Occupation: input.Occupation,
            Language: input.Language,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&patient).Error; err != nil {
			log.Println("Database error while inserting patient:", err)
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
