package main

import (
	"log"
	"net/http"

	"github.com/PragaL15/med_admin_backend/database"
	"github.com/PragaL15/med_admin_backend/src/routers/user"
	"github.com/gorilla/handlers"

)

func main() {
	db, err := database.InitializeDB() 
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Failed to get raw database connection: %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}()

	router := routers.SetupRoutes(db)

	corsOrigin := handlers.AllowedOrigins([]string{"http://localhost:5173"}) 
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}) 
	corsHeaders := handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Accept", "Authorization"}) 

	corsMiddleware := handlers.CORS(corsOrigin, corsMethods, corsHeaders)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", corsMiddleware(router)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
