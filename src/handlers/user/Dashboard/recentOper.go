package handlers

import (
    "context"
    "encoding/json"
    "log"
    "net/http"

    "github.com/PragaL15/med_admin_backend/database"
    models "github.com/PragaL15/med_admin_backend/src/model"
)

// RecentOperations handles fetching all records from the admitted table
func RecentOperation(w http.ResponseWriter, r *http.Request) {
    // Enable CORS headers for all domains (can restrict this to specific domains if needed)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // Handle OPTIONS request for CORS preflight
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    // Restrict to GET method only
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // SQL query to retrieve all records from the admitted table
    query := `
    SELECT 
        a.id,                   -- ID from admitted table
        a.p_id,                 -- Patient ID from admitted table
        p.p_name,               -- Patient name from patient_id table
        a.p_health,             -- Patient health from admitted table
        a.p_operation,          -- Operation type from admitted table
        a.p_operation_date,     -- Operation date from admitted table
        a.p_operated_doctor,    -- Operating doctor from admitted table
        a.duration_admit,       -- Duration of admission from admitted table
        a.ward_no               -- Ward number from admitted table
    FROM 
        admitted a
    JOIN 
        patient_id p 
        ON a.p_id = p.p_id;  -- Join admitted table with patient_id table using p_id
    `

    // Query execution
    rows, err := database.DB.Query(context.Background(), query)
    if err != nil {
        log.Printf("Error executing query: %v", err) // Log error details
        http.Error(w, "Error fetching admitted records", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var admittedRecords []models.Admitted

    // Scan data into Admitted model
    for rows.Next() {
        var admitted models.Admitted
        if err := rows.Scan(
            &admitted.ID, &admitted.PID, &admitted.PName, &admitted.PHealth, 
            &admitted.POperation, &admitted.POperationDate, &admitted.POperatedDoctor, 
            &admitted.DurationAdmit, &admitted.WardNo,
        ); err != nil {
            log.Printf("Error scanning row: %v", err) // Log scanning error
            http.Error(w, "Error scanning admitted records", http.StatusInternalServerError)
            return
        }
        admittedRecords = append(admittedRecords, admitted)
    }

    // Check if any records were found
    if len(admittedRecords) == 0 {
        http.Error(w, "No admitted records found", http.StatusNotFound)
        return
    }

    // Send the records as a JSON response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    err = json.NewEncoder(w).Encode(admittedRecords)
    if err != nil {
        log.Printf("Error encoding JSON response: %v", err)
        http.Error(w, "Error sending response", http.StatusInternalServerError)
    }
}
