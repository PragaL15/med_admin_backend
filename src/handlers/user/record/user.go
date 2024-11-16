package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	models "github.com/PragaL15/med_admin_backend/src/model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// GetRecords retrieves all records.
func GetRecords(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var records []models.Record

		// Use GORM to find all records
		if err := db.Find(&records).Error; err != nil {
			http.Error(w, "Failed to fetch records", http.StatusInternalServerError)
			log.Println("Database query error:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(records)
	}
}

// GetRecordByID retrieves a record by ID.
func GetRecordByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var record models.Record
		if err := db.First(&record, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "Record not found", http.StatusNotFound)
			} else {
				http.Error(w, "Failed to fetch record", http.StatusInternalServerError)
			}
			log.Println("Record fetch error:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(record)
	}
}

func CreateRecord(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record models.Record
		if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		record.CreatedAt = time.Now()
		record.UpdatedAt = time.Now()

		if err := db.Create(&record).Error; err != nil {
			http.Error(w, "Failed to create record", http.StatusInternalServerError)
			log.Println("Record creation error:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(record)
	}
}

func UpdateRecord(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		record.UpdatedAt = time.Now()
		if err := db.Model(&models.Record{}).Where("id = ?", id).Updates(map[string]interface{}{
			"PID":         record.PID,
			"DID":         record.DID,
			"Date":        record.Date,
			"Description": record.Description,
			"Prescription": record.Prescription,
			"UpdatedAt":   record.UpdatedAt,
		}).Error; err != nil {
			http.Error(w, "Failed to update record", http.StatusInternalServerError)
			log.Println("Record update error:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}
func UpdateDescriptionByPID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		if data.Description == "" {
			http.Error(w, "Description cannot be empty", http.StatusBadRequest)
			return
		}

		if err := db.Model(&models.Record{}).Where("p_id = ?", pID).Update("description", data.Description).Error; err != nil {
			http.Error(w, "Failed to update description", http.StatusInternalServerError)
			log.Println("Description update error:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdatePrescription(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type UpdateData struct {
			IDs          []int  `json:"ids"`
			Prescription string `json:"prescription"`
		}
		var data UpdateData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if data.Prescription == "" {
			http.Error(w, "Prescription cannot be empty", http.StatusBadRequest)
			return
		}

		if err := db.Model(&models.Record{}).Where("id IN ?", data.IDs).Update("prescription", data.Prescription).Error; err != nil {
			http.Error(w, "Failed to update prescription", http.StatusInternalServerError)
			log.Println("Prescription update error:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteRecord(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := db.Delete(&models.Record{}, id).Error; err != nil {
			http.Error(w, "Failed to delete record", http.StatusInternalServerError)
			log.Println("Record deletion error:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}
