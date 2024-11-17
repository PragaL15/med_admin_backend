package routers

import (
    addDetailsHandlers "github.com/PragaL15/med_admin_backend/src/handlers/user/AddDetails"
    dashboardHandlers "github.com/PragaL15/med_admin_backend/src/handlers/user/Dashboard"
    loginHandlers "github.com/PragaL15/med_admin_backend/src/handlers/user/login"
    recordHandlers "github.com/PragaL15/med_admin_backend/src/handlers/user/record"
		appointmentHandlers "github.com/PragaL15/med_admin_backend/src/handlers/user/BookAppointment" 
	
    "github.com/PragaL15/med_admin_backend/src/middleware"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *mux.Router {
    router := mux.NewRouter()

    corsMiddleware := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:5173"}), 
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Accept", "Authorization"}),
    )
    router.Use(corsMiddleware)

    router.HandleFunc("/login", loginHandlers.Login).Methods("POST")

    apiRouter := router.PathPrefix("/api").Subrouter()
    apiRouter.Use(middleware.RoleBasedAccessMiddleware(db)) 
    apiRouter.Use(corsMiddleware) 

    setupRecordsRoutes(apiRouter.PathPrefix("/records").Subrouter(), db)
    setupPatientsRoutes(apiRouter.PathPrefix("/patients").Subrouter(), db)
    setupDashboardRoutes(apiRouter.PathPrefix("/dashboard").Subrouter(), db)
    setupDoctorsRoutes(apiRouter.PathPrefix("/doctors").Subrouter(), db)
	setupAppointmentsRoutes(apiRouter.PathPrefix("/appointments").Subrouter(), db)
       
    return router
}

func setupRecordsRoutes(router *mux.Router, db *gorm.DB) {
    router.HandleFunc("", recordHandlers.GetRecords(db)).Methods("GET")
    router.HandleFunc("/{id}", recordHandlers.GetRecordByID(db)).Methods("GET")
    router.HandleFunc("", recordHandlers.CreateRecord(db)).Methods("POST")
    router.HandleFunc("/{id}", recordHandlers.UpdateRecord(db)).Methods("PUT")
    router.HandleFunc("/{id}", recordHandlers.DeleteRecord(db)).Methods("DELETE")
    router.HandleFunc("/{p_id}/description", recordHandlers.UpdateDescriptionByPID(db)).Methods("PUT")
}

func setupPatientsRoutes(router *mux.Router, db *gorm.DB) {
    router.HandleFunc("", recordHandlers.GetAllPatients(db)).Methods("GET")
    router.HandleFunc("", recordHandlers.CreatePatient(db)).Methods("POST")
    router.HandleFunc("/{id}", recordHandlers.GetPatientByID(db)).Methods("GET")
    router.HandleFunc("/{id}", recordHandlers.UpdatePatient(db)).Methods("PUT")
    router.HandleFunc("/{id}", recordHandlers.DeletePatient(db)).Methods("DELETE")
}


func setupDashboardRoutes(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/patient-status", dashboardHandlers.GetPatientStatusForGraph(db)).Methods("GET")
	router.HandleFunc("/patientDetails", addDetailsHandlers.AddPatient(db)).Methods("POST")
	router.HandleFunc("/AppointmentTable", dashboardHandlers.GetAppointments(db)).Methods("GET")
	router.HandleFunc("/AdmittedTable", dashboardHandlers.GetAdmittedPatients(db)).Methods("GET")
	router.HandleFunc("/RecentOperation", dashboardHandlers.RecentOperation(db)).Methods("GET")
}
func setupAppointmentsRoutes(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/create", appointmentHandlers.CreateAppointment(db)).Methods("POST", "OPTIONS") 
    router.HandleFunc("/doctors-patients", appointmentHandlers.GetDoctorsAndPatients(db)).Methods("GET")
}
func setupDoctorsRoutes(router *mux.Router, db *gorm.DB) {
    router.HandleFunc("", recordHandlers.GetAllDoctors(db)).Methods("GET")
    router.HandleFunc("", recordHandlers.CreateDoctor(db)).Methods("POST")
    router.HandleFunc("/{id}", recordHandlers.GetDoctorByID(db)).Methods("GET")
    router.HandleFunc("/{id}", recordHandlers.UpdateDoctor(db)).Methods("PUT")
    router.HandleFunc("/{id}", recordHandlers.DeleteDoctor(db)).Methods("DELETE")
}