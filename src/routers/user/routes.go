package routers

import (
    addDetailsHandlers "github.com/PragaL15/med_admin_backend/src/handlers/user/AddDetails"
    dashboardHandlers "github.com/PragaL15/med_admin_backend/src/handlers/user/Dashboard"
    loginHandlers "github.com/PragaL15/med_admin_backend/src/handlers/user/login"
    recordHandlers "github.com/PragaL15/med_admin_backend/src/handlers/user/record"
    "github.com/PragaL15/med_admin_backend/src/middleware"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "net/http"
)

// SetupRoutes initializes and returns the configured API routes.
func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Define global CORS policy (apply to all routes)
    corsMiddleware := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Adjust origin for production
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Accept", "Authorization"}),
    )
    router.Use(corsMiddleware)
    
    // Public routes (no authentication required)
    router.HandleFunc("/login", loginHandlers.Login).Methods("POST")

    // API routes that require authentication
    apiRouter := router.PathPrefix("/api").Subrouter()
    apiRouter.Use(middleware.JWTAuthMiddleware)
		apiRouter.Use(corsMiddleware)

    // Setup specific API route groups
    setupRecordsRoutes(apiRouter.PathPrefix("/records").Subrouter())
    setupPatientsRoutes(apiRouter.PathPrefix("/patients").Subrouter())
    setupDashboardRoutes(apiRouter)
    setupDoctorsRoutes(apiRouter.PathPrefix("/doctors").Subrouter())

    return router
}

// setupRecordsRoutes configures routes related to medical records.
func setupRecordsRoutes(router *mux.Router) {
    router.Handle("", http.HandlerFunc(recordHandlers.GetRecords)).Methods("GET")
    router.Handle("/{id}", http.HandlerFunc(recordHandlers.GetRecordByID)).Methods("GET")
    router.Handle("", http.HandlerFunc(recordHandlers.CreateRecord)).Methods("POST")
    router.Handle("/{id}", http.HandlerFunc(recordHandlers.UpdateRecord)).Methods("PUT")
    router.Handle("/{id}", http.HandlerFunc(recordHandlers.DeleteRecord)).Methods("DELETE")
    router.Handle("/{p_id}/description", http.HandlerFunc(recordHandlers.UpdateDescriptionByPID)).Methods("PUT")
}

// setupPatientsRoutes configures routes related to patient management.
func setupPatientsRoutes(router *mux.Router) {
    router.Handle("", http.HandlerFunc(recordHandlers.GetAllPatients)).Methods("GET")
    router.Handle("", http.HandlerFunc(recordHandlers.CreatePatient)).Methods("POST")
    router.Handle("/{id}", http.HandlerFunc(recordHandlers.GetPatientByID)).Methods("GET")
    router.Handle("/{id}", http.HandlerFunc(recordHandlers.UpdatePatient)).Methods("PUT")
    router.Handle("/{id}", http.HandlerFunc(recordHandlers.DeletePatient)).Methods("DELETE")
}

// setupDashboardRoutes configures routes related to dashboard functionality.
func setupDashboardRoutes(router *mux.Router) {
    router.Handle("/patient-status", http.HandlerFunc(dashboardHandlers.GetPatientStatusForGraph)).Methods("GET")
    router.Handle("/patientDetails", http.HandlerFunc(addDetailsHandlers.AddPatient)).Methods("POST")
    router.Handle("/AppointmentTable", http.HandlerFunc(dashboardHandlers.GetAppointments)).Methods("GET")
    router.Handle("/AdmittedTable", http.HandlerFunc(dashboardHandlers.GetAdmitted)).Methods("GET")
    router.Handle("/RecentOperation", http.HandlerFunc(dashboardHandlers.RecentOperation)).Methods("GET")
}

// setupDoctorsRoutes configures routes related to doctors.
func setupDoctorsRoutes(router *mux.Router) {
    router.Handle("", http.HandlerFunc(recordHandlers.GetAllDoctors)).Methods("GET")
    router.Handle("", http.HandlerFunc(recordHandlers.CreateDoctor)).Methods("POST")
    router.Handle("/{id}", http.HandlerFunc(recordHandlers.GetDoctorByID)).Methods("GET")
    router.Handle("/{id}", http.HandlerFunc(recordHandlers.UpdateDoctor)).Methods("PUT")
    router.Handle("/{id}", http.HandlerFunc(recordHandlers.DeleteDoctor)).Methods("DELETE")
}
