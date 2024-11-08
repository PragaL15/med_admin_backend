package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "github.com/PragaL15/med_admin_backend/database"
)

// PatientStatusRecord holds the structure of each record returned for graphing
type PatientStatusRecord struct {
    PatientID int    `json:"p_id"`
    Month     string `json:"month"`
    Status    string `json:"p_status"`
}

// GetPatientStatusForGraph retrieves patient status data for graphing by month
func GetPatientStatusForGraph(w http.ResponseWriter, r *http.Request) {
    query := `
        SELECT 
            record.p_id,
            record.date,
            patient_id.p_status
        FROM 
            record
        JOIN 
            patient_id ON record.p_id = patient_id.p_id;
    `

    rows, err := database.DB.Query(context.Background(), query)
    if err != nil {
        http.Error(w, "Error fetching patient status data", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var records []PatientStatusRecord
    for rows.Next() {
        var record PatientStatusRecord
        var date time.Time
        err := rows.Scan(&record.PatientID, &date, &record.Status)
        if err != nil {
            http.Error(w, "Error scanning patient status data", http.StatusInternalServerError)
            return
        }
        // Format the date to display only the month
        record.Month = date.Format("January")
        records = append(records, record)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(records)
}
