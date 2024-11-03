package main

import (
	"log"
	"net/http"

	"github.com/PragaL15/med_admin_backend/database"
	routers "github.com/PragaL15/med_admin_backend/src/routers/user"
	"github.com/gorilla/handlers"
)

func main() {
	database.InitializeDB()
	defer database.DB.Close()

	// Initialize router
	router := routers.SetupRoutes()

	// CORS settings
	corsOrigin := handlers.AllowedOrigins([]string{"http://localhost:5173"}) // Allow your frontend origin
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Accept"})

	// Start the server with CORS enabled
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", handlers.CORS(corsOrigin, corsMethods, corsHeaders)(router)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
