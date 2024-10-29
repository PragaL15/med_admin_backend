// database/database.go
package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitializeDB establishes a connection to the PostgreSQL database and returns the DB instance
func InitializeDB() *gorm.DB {
	// Define the DSN (Data Source Name) with sslmode=disable for local development
	dsn := "host=localhost user=postgres password=pragalya123 dbname=admin_med port=5432"
	
	// Connect to the database using GORM and PostgreSQL driver
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Print a detailed error message if connection fails
		panic(fmt.Sprintf("failed to connect to the database: %v", err))
	}

	fmt.Println("Database connection established successfully")
	return db
}
