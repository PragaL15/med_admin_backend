package database

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() (*gorm.DB, error) {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
        return nil, err
    }

    // Print environment variables to verify they are loaded
    log.Println("DB_USER:", os.Getenv("DB_USER"))
    log.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
    log.Println("DB_HOST:", os.Getenv("DB_HOST"))
    log.Println("DB_PORT:", os.Getenv("DB_PORT"))
    log.Println("DB_NAME:", os.Getenv("DB_NAME"))

    // Construct the database connection URL
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"))

    // Open a connection to the database using GORM
   
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
        return nil, err
    }
    log.Println("Connected to the database")

    return DB, nil
}
