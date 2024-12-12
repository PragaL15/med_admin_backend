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
##### Decoding the request body

`var patient models.Patient
if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
    http.Error(w, "Invalid request body", http.StatusBadRequest)
    return
}
`
