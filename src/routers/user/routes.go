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

// ChainMiddlewares chains multiple middlewares to a handler function
func ChainMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
    for _, middleware := range middlewares {
        handler = middleware(handler)
    }
    return handler
}

// SetupRoutes initializes and returns the configured API routes.
func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Define CORS policy (globally applied)
    corsMiddleware := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Change to env-config if needed
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Accept", "Authorization"}),
    )

    // Public routes (no authentication required)
    router.HandleFunc("/login", loginHandlers.Login).Methods("POST")

    // API routes requiring authentication and/or specific roles
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
    router.Handle("", ChainMiddlewares(http.HandlerFunc(recordHandlers.GetRecords), middleware.RoleAuthMiddleware("2"))).Methods("GET")
    router.Handle("/{id}", ChainMiddlewares(http.HandlerFunc(recordHandlers.GetRecordByID), middleware.RoleAuthMiddleware("2"))).Methods("GET")
    router.Handle("", ChainMiddlewares(http.HandlerFunc(recordHandlers.CreateRecord), middleware.RoleAuthMiddleware("2"))).Methods("POST")
    router.Handle("/{id}", ChainMiddlewares(http.HandlerFunc(recordHandlers.UpdateRecord), middleware.RoleAuthMiddleware("2"))).Methods("PUT")
    router.Handle("/{id}", ChainMiddlewares(http.HandlerFunc(recordHandlers.DeleteRecord), middleware.RoleAuthMiddleware("2"))).Methods("DELETE")
    router.Handle("/{p_id}/description", ChainMiddlewares(http.HandlerFunc(recordHandlers.UpdateDescriptionByPID), middleware.RoleAuthMiddleware("3"))).Methods("PUT")
}

// setupPatientsRoutes configures routes related to patient management.
func setupPatientsRoutes(router *mux.Router) {
    router.Handle("", ChainMiddlewares(http.HandlerFunc(recordHandlers.GetAllPatients), middleware.RoleAuthMiddleware("2"))).Methods("GET")
    router.Handle("", ChainMiddlewares(http.HandlerFunc(recordHandlers.CreatePatient), middleware.RoleAuthMiddleware("2"))).Methods("POST")
    router.Handle("/{id}", ChainMiddlewares(http.HandlerFunc(recordHandlers.GetPatientByID), middleware.RoleAuthMiddleware("2"))).Methods("GET")
    router.Handle("/{id}", ChainMiddlewares(http.HandlerFunc(recordHandlers.UpdatePatient), middleware.RoleAuthMiddleware("2"))).Methods("PUT")
    router.Handle("/{id}", ChainMiddlewares(http.HandlerFunc(recordHandlers.DeletePatient), middleware.RoleAuthMiddleware("2"))).Methods("DELETE")
}

// setupDashboardRoutes configures routes related to dashboard functionality.
func setupDashboardRoutes(router *mux.Router) {
    router.Handle("/patient-status", ChainMiddlewares(http.HandlerFunc(dashboardHandlers.GetPatientStatusForGraph), middleware.RoleAuthMiddleware("2"))).Methods("GET")
    router.Handle("/patientDetails", ChainMiddlewares(http.HandlerFunc(addDetailsHandlers.AddPatient), middleware.RoleAuthMiddleware("1"))).Methods("POST")
    router.Handle("/AppointmentTable", ChainMiddlewares(http.HandlerFunc(dashboardHandlers.GetAppointments), middleware.RoleAuthMiddleware("2"))).Methods("GET")
    router.Handle("/AdmittedTable", ChainMiddlewares(http.HandlerFunc(dashboardHandlers.GetAdmitted), middleware.RoleAuthMiddleware("2"))).Methods("GET")
    router.Handle("/RecentOperation", ChainMiddlewares(http.HandlerFunc(dashboardHandlers.RecentOperation), middleware.RoleAuthMiddleware("2"))).Methods("GET")
}

// setupDoctorsRoutes configures routes related to doctors.
func setupDoctorsRoutes(router *mux.Router) {
    router.Handle("", ChainMiddlewares(http.HandlerFunc(recordHandlers.GetAllDoctors), middleware.RoleAuthMiddleware("3"))).Methods("GET")
    router.Handle("", ChainMiddlewares(http.HandlerFunc(recordHandlers.CreateDoctor), middleware.RoleAuthMiddleware("2"))).Methods("POST")
    router.Handle("/{id}", ChainMiddlewares(http.HandlerFunc(recordHandlers.GetDoctorByID), middleware.RoleAuthMiddleware("3"))).Methods("GET")
    router.Handle("/{id}", ChainMiddlewares(http.HandlerFunc(recordHandlers.UpdateDoctor), middleware.RoleAuthMiddleware("2"))).Methods("PUT")
    router.Handle("/{id}", ChainMiddlewares(http.HandlerFunc(recordHandlers.DeleteDoctor), middleware.RoleAuthMiddleware("2"))).Methods("DELETE")
}
