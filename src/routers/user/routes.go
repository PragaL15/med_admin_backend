package routers

import (
	handlers "github.com/PragaL15/med_admin_backend/src/handlers/user"
	"github.com/gorilla/mux"
)

// SetupRoutes initializes the API routes.
func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Define routes and associate them with handlers
    router.HandleFunc("/records", handlers.GetRecords).Methods("GET")
    router.HandleFunc("/records/{id}", handlers.GetRecordByID).Methods("GET")
    router.HandleFunc("/records", handlers.CreateRecord).Methods("POST")
    router.HandleFunc("/records/{id}", handlers.UpdateRecord).Methods("PUT")
    router.HandleFunc("/records/{id}", handlers.DeleteRecord).Methods("DELETE")

    return router
		
}
