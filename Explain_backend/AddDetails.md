##### imports 
`import (
    "encoding/json"
    "net/http"
    "time"
    models "github.com/PragaL15/med_admin_backend/src/model"
    "gorm.io/gorm"
)`
1. **encoding/json** - Built in go package to work with JSON , it includes the `marshaling` and `unmarshaling`.
2. `gorm.io/gorm` - External Library used for database interaction.
---
##### Function Definition 

`func AddPatient(db *gorm.DB) http.HandlerFunc {
`
1. `db *gorm.DB` - recives the GORM database connection as a dependency.

2. `http.HandlerFunc` - returns a functn that confirms to HTTP handler interface.

---
##### Handler Logic
`    return func(w http.ResponseWriter, r *http.Request) {
`
1. `w http.ResponseWriter`: Used to write the HTTP response back to the client.
2. `r *http.Request`: Represents the incoming HTTP request, which includes headers, method, body, etc.
---
##### HTTP method check and content-type validation 

`        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }
`

1. Checks if the HTTP method is not POST. If not, it sends a 405 Method Not Allowed error and ends the handler.

`        if r.Header.Get("Content-Type") != "application/json" {
            http.Error(w, "Invalid Content-Type, expected application/json", http.StatusUnsupportedMediaType)
            return
        }
`
1. Ensures the Content-Type header is application/json. If not, it sends a 415 Unsupported Media Type error and ends the handler.
---
1. **Data in JSON Format**:
   - The client sends a request (e.g., `POST`) to the server with data in JSON format.
   - Example JSON payload from the client:
     ```json
     {
       "name": "John Doe",
       "phone": "9876543210",
       "email": "john.doe@example.com",
       "status": "active",
       "address": "123 Elm Street",
       "mode": "offline",
       "age": 30,
       "gender": "male"
     }
     ```

2. **Decoding JSON**:
   - The server receives this JSON payload in the HTTP request body.
   - To work with this data in Go, you use `json.NewDecoder(r.Body).Decode(&patient)`:
     - This decodes the JSON into a Go struct (`Patient`).
     - The fields in the JSON are mapped to corresponding fields in the `Patient` struct based on their `json` tags.

3. **Validation**:
   - Once decoded, you can validate the data in the `Patient` struct:
     - For example, ensure `Name` is not empty, `Age` is greater than 0, and `Gender` is valid.

4. **Saving to the Database**:
   - After validating the data, you save the `Patient` object into the database using an ORM like GORM.
   - This is done with:
     ```go
     db.Create(&patient)
     ```
     - `Create` takes the populated `Patient` struct and saves it as a new record in the database.

5. **Response to the Client**:
   - Once the patient is successfully saved, the server sends back a response confirming the operation.
   - For example:
     ```json
     {
       "message": "Patient created successfully",
       "patient": {
         "id": 1,
         "name": "John Doe",
         "phone": "9876543210",
         "email": "john.doe@example.com",
         "status": "active",
         "address": "123 Elm Street",
         "mode": "offline",
         "age": 30,
         "gender": "male",
         "created_at": "2024-12-11T12:00:00Z",
         "updated_at": "2024-12-11T12:00:00Z"
       }
     }
     ```

### **Why We Need `json.Decode`**
- JSON is a universal data format used in APIs.
- In Go, to work with this JSON data, you need to:

  1. **Parse it**: Convert the raw JSON into a usable Go struct (`Patient`).
  2. **Validate it**: Ensure the data meets your application’s requirements.
  3. **Process it**: Save it to the database or use it for further logic.


