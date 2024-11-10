package routers

import (
	handlers "github.com/PragaL15/med_admin_backend/src/handlers/user"
		"github.com/gorilla/mux"
)

// SetupRoutes initializes the API routes.
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

//login route
router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Define routes and associate them with handlers for records
	router.HandleFunc("/records", handlers.GetRecords).Methods("GET")
	router.HandleFunc("/records/{id}", handlers.GetRecordByID).Methods("GET")
	router.HandleFunc("/records", handlers.CreateRecord).Methods("POST")
	router.HandleFunc("/records/{id}", handlers.UpdateRecord).Methods("PUT")
	router.HandleFunc("/records/{id}", handlers.DeleteRecord).Methods("DELETE")
  router.HandleFunc("/records/{p_id}/description", handlers.UpdateDescriptionByPID).Methods("PUT")

	// Define routes and associate them with handlers for patients
	router.HandleFunc("/patients", handlers.GetAllPatients).Methods("GET")
	router.HandleFunc("/patients", handlers.CreatePatient).Methods("POST")
	router.HandleFunc("/patients/{id}", handlers.GetPatientByID).Methods("GET")
	router.HandleFunc("/patients/{id}", handlers.UpdatePatient).Methods("PUT")
	router.HandleFunc("/patients/{id}", handlers.DeletePatient).Methods("DELETE")
  router.HandleFunc("/patient-status", handlers.GetPatientStatusForGraph).Methods("GET")
  router.HandleFunc("/patientDetails", handlers.AddPatient).Methods("POST")
	router.HandleFunc("/AppointmentTable", handlers.GetAppointments).Methods("GET")
	router.HandleFunc("/AdmittedTable", handlers.GetAdmitted).Methods("GET")
	router.HandleFunc("/RecentOperation",handlers.RecentOperation).Methods("GET")

	//Defined the routes associated with handlers for Doctors.
router.HandleFunc("/doctors", handlers.CreateDoctor).Methods("POST")
router.HandleFunc("/doctors", handlers.GetAllDoctors).Methods("GET")
router.HandleFunc("/doctors/{id}", handlers.GetDoctorByID).Methods("GET")
router.HandleFunc("/doctors/{id}", handlers.UpdateDoctor).Methods("PUT")
router.HandleFunc("/doctors/{id}", handlers.DeleteDoctor).Methods("DELETE")

	return router
}
