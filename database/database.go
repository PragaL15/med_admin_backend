package database

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/jackc/pgx/v4/pgxpool"
    "github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func InitializeDB() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
    }

    // Print environment variables to verify they are loaded
    log.Println("DB_USER:", os.Getenv("DB_USER"))
    log.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
    log.Println("DB_HOST:", os.Getenv("DB_HOST"))
    log.Println("DB_PORT:", os.Getenv("DB_PORT"))
    log.Println("DB_NAME:", os.Getenv("DB_NAME"))

    databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"))

    DB, err = pgxpool.Connect(context.Background(), databaseURL)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
        os.Exit(1)
    }
    log.Println("Connected to the database")
}
