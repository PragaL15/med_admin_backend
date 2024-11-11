package routers

import (
	addDetailsHandlers "github.com/PragaL15/med_admin_backend/src/handlers/user/AddDetails"
	dashboardHandlers "github.com/PragaL15/med_admin_backend/src/handlers/user/Dashboard"
	loginHandlers "github.com/PragaL15/med_admin_backend/src/handlers/user/login"
	recordHandlers "github.com/PragaL15/med_admin_backend/src/handlers/user/record"
	"github.com/gorilla/mux"
)

// SetupRoutes initializes the API routes.
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Login route
	router.HandleFunc("/login", loginHandlers.Login).Methods("POST")

	// Define routes and associate them with handlers for records
	router.HandleFunc("/records", recordHandlers.GetRecords).Methods("GET")
	router.HandleFunc("/records/{id}", recordHandlers.GetRecordByID).Methods("GET")
	router.HandleFunc("/records", recordHandlers.CreateRecord).Methods("POST")
	router.HandleFunc("/records/{id}", recordHandlers.UpdateRecord).Methods("PUT")
	router.HandleFunc("/records/{id}", recordHandlers.DeleteRecord).Methods("DELETE")
	router.HandleFunc("/records/{p_id}/description", recordHandlers.UpdateDescriptionByPID).Methods("PUT")

	// Define routes and associate them with handlers for patients
	router.HandleFunc("/patients", recordHandlers.GetAllPatients).Methods("GET")
	router.HandleFunc("/patients", recordHandlers.CreatePatient).Methods("POST")
	router.HandleFunc("/patients/{id}",recordHandlers.GetPatientByID).Methods("GET")
	router.HandleFunc("/patients/{id}", recordHandlers.UpdatePatient).Methods("PUT")
	router.HandleFunc("/patients/{id}",recordHandlers.DeletePatient).Methods("DELETE")
	router.HandleFunc("/patient-status", dashboardHandlers.GetPatientStatusForGraph).Methods("GET")
	router.HandleFunc("/patientDetails", addDetailsHandlers.AddPatient).Methods("POST")
	router.HandleFunc("/AppointmentTable", dashboardHandlers.GetAppointments).Methods("GET")
	router.HandleFunc("/AdmittedTable", dashboardHandlers.GetAdmitted).Methods("GET")
	router.HandleFunc("/RecentOperation", dashboardHandlers.RecentOperation).Methods("GET")

	// Define routes and associate them with handlers for doctors
	router.HandleFunc("/doctors", recordHandlers.CreateDoctor).Methods("POST")
	router.HandleFunc("/doctors",recordHandlers.GetAllDoctors).Methods("GET")
	router.HandleFunc("/doctors/{id}", recordHandlers.GetDoctorByID).Methods("GET")
	router.HandleFunc("/doctors/{id}", recordHandlers.UpdateDoctor).Methods("PUT")
	router.HandleFunc("/doctors/{id}",recordHandlers.DeleteDoctor).Methods("DELETE")

	return router
}
