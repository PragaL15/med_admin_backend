package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "github.com/PragaL15/med_admin_backend/database"
    models "github.com/PragaL15/med_admin_backend/src/model"
)

// GetAppointments handles fetching all appointment records
func GetAppointments(w http.ResponseWriter, r *http.Request) {
    // Enable CORS headers
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

    // SQL query to retrieve all appointments with patient details
    query := `
        SELECT a.id, a.p_id, p.p_name, p.p_number, a.created_at, a.updated_at, a.app_date, 
               a.p_health, a.d_id, a.time, a.problem_hint ,a.appo_status
        FROM appointments a
        JOIN patient_id p ON a.p_id = p.p_id;
    `

    // Execute the query
    rows, err := database.DB.Query(context.Background(), query)
    if err != nil {
        http.Error(w, "Error fetching appointments", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var appointments []models.Appointment

    // Iterate over rows and scan data into appointment struct
    for rows.Next() {
        var appointment models.Appointment
        var dbTime time.Time

        if err := rows.Scan(
            &appointment.ID, &appointment.PID, &appointment.PName, &appointment.PNumber,
            &appointment.CreatedAt, &appointment.UpdatedAt, &appointment.AppDate, &appointment.PHealth,
            &appointment.DID, &dbTime, &appointment.ProblemHint,&appointment.AppoStatus ,
        ); err != nil {
            http.Error(w, "Error scanning appointments", http.StatusInternalServerError)
            return
        }

        // Format dbTime to 12-hour format and store in appointment.Time as a string
        appointment.Time = dbTime.Format("03:04 PM")

        // Append the appointment to the list
        appointments = append(appointments, appointment)
    }

    // Respond with the fetched appointment data as JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(appointments)
}
