
### **Package and Imports**
```go
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
	models "github.com/PragaL15/med_admin_backend/src/model"
)
```
1. **Package Declaration**:
   - Declares the `handlers` package, which contains HTTP handler functions for different routes.
2. **Imports**:
   - `encoding/json`: To encode and decode JSON data.
   - `log`: For logging errors or debugging information.
   - `net/http`: To handle HTTP requests and responses.
   - `time`: To handle date and time formatting/parsing.
   - `gorm.io/gorm`: GORM is used as the ORM for database operations.
   - `models`: Imports the `AppointmentPost` struct (defined in `src/model`) to represent appointment data.

---
### **Function Definition**
```go
func CreateAppointment(db *gorm.DB) http.HandlerFunc {
```
- **Function Name**: `CreateAppointment`
- **Parameter**: `db` (a GORM database instance for interacting with the database).
- **Returns**: An `http.HandlerFunc`, which can handle HTTP requests.

---

### **Handler Function**
```go
return func(w http.ResponseWriter, r *http.Request) {
```
- Returns an anonymous function that takes `w` (HTTP response writer) and `r` (HTTP request) as arguments.

---

### **CORS Preflight Handling**
```go
if r.Method == http.MethodOptions {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	return
}
```
- Handles **CORS Preflight Requests**:
  - Allows cross-origin requests by specifying headers and methods.
  - Sends an `HTTP 200 OK` response for `OPTIONS` requests.

---

### **Method Check**
```go
if r.Method != http.MethodPost {
	http.Error(w, `{"status": false, "message": "Method not allowed"}`, http.StatusMethodNotAllowed)
	return
}
```
- Ensures that only `POST` requests are allowed.
- If any other method is used, it returns a `405 Method Not Allowed` error with a JSON response.

---

### **JSON Decoding**
```go
var appointment models.AppointmentPost
if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
	log.Println("Error decoding request body:", err)
	http.Error(w, `{"status": false, "message": "Invalid request body"}`, http.StatusBadRequest)
	return
}
```
1. Creates a variable `appointment` of type `models.AppointmentPost`.
2. Uses `json.NewDecoder(r.Body).Decode` to parse the JSON body into the `appointment` struct.
3. If decoding fails (invalid JSON or mismatch with struct), it logs the error and returns a `400 Bad Request` error.

---

### **Required Field Validation**
```go
if appointment.PID == 0 || appointment.DID == 0 || appointment.AppDate == "" || appointment.Time == "" {
	http.Error(w, `{"status": false, "message": "Missing required fields"}`, http.StatusBadRequest)
	return
}
```
- Checks if required fields (`PID`, `DID`, `AppDate`, `Time`) are missing or invalid.
- If any field is invalid, it returns a `400 Bad Request` error with a JSON response.

---

### **Date and Time Parsing**
```go
parsedDate, err := time.Parse("02-01-2006", appointment.AppDate)
if err != nil {
	log.Println("Error parsing app_date:", err)
	http.Error(w, `{"status": false, "message": "Invalid date format. Expected DD-MM-YYYY"}`, http.StatusBadRequest)
	return
}

parsedTime, err := time.Parse("15:04:05", appointment.Time)
if err != nil {
	log.Println("Error parsing time:", err)
	http.Error(w, `{"status": false, "message": "Invalid time format. Expected HH:mm:ss"}`, http.StatusBadRequest)
	return
}
```
1. **Date Parsing**:
   - Uses `time.Parse` to validate and parse `AppDate` (e.g., "11-12-2024") into a `time.Time` object.
   - Expected format: `DD-MM-YYYY`.
   - If parsing fails, logs the error and returns a `400 Bad Request` error.
2. **Time Parsing**:
   - Validates and parses the `Time` field (e.g., "14:30:00").
   - Expected format: `HH:mm:ss`.

---

### **Combine Date and Time**
```go
combinedDateTime := time.Date(
	parsedDate.Year(), parsedDate.Month(), parsedDate.Day(),
	parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), 0, time.UTC,
)
```
- Combines the parsed `AppDate` and `Time` into a single `time.Time` object, `combinedDateTime`.

---

### **Format Date and Time**
```go
appointment.AppDate = combinedDateTime.Format("2006-01-02") 
appointment.Time = combinedDateTime.Format("15:04:05")
```
- Converts `combinedDateTime` back into formatted strings:
  - `AppDate`: `YYYY-MM-DD` (e.g., "2024-12-11").
  - `Time`: `HH:mm:ss` (e.g., "14:30:00").

---

### **Database Insertion**
```go
if err := db.Create(&appointment).Error; err != nil {
	log.Printf("Error creating appointment: %v", err)
	http.Error(w, `{"status": false, "message": "Failed to create appointment"}`, http.StatusInternalServerError)
	return
}
```
- Saves the `appointment` object into the database using GORM.
- If the insertion fails, it logs the error and returns a `500 Internal Server Error` response.

---

### **Response to Client**
```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
response := map[string]interface{}{
	"status":  true,
	"message": "Appointment created successfully",
	"data":    appointment,
}
if err := json.NewEncoder(w).Encode(response); err != nil {
	log.Printf("Error encoding JSON response: %v", err)
	http.Error(w, `{"status": false, "message": "Error sending response"}`, http.StatusInternalServerError)
}
```
1. Sets the response `Content-Type` to `application/json`.
2. Responds with `201 Created` status code.
3. Sends a JSON response containing:
   - Success status.
   - Confirmation message.
   - The created appointment record.

If JSON encoding fails, it logs the error and sends a `500 Internal Server Error`.

---

### **Summary**
This function:
1. Validates the request method and JSON format.
2. Parses and validates the input data.
3. Combines and formats the date and time.
4. Saves the appointment in the database.
5. Sends an appropriate JSON response to the client.

