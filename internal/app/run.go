package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/handlers"
	"github.com/Rizabekus/effective-mobile-rest-api/internal/services"
	"github.com/Rizabekus/effective-mobile-rest-api/internal/storage"
	"github.com/Rizabekus/effective-mobile-rest-api/pkg/loggers.go"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	// migrate "github.com/rubenv/sql-migrate"
)

func Run() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	loggers.InfoLog.Println("Loaded the configuration data from .env")
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	loggers.InfoLog.Println("Successfully connected to database")
	storage := storage.StorageInstance(db)
	service := services.ServiceInstance(storage)
	handler := handlers.HandlersInstance(service)

	Routes(handler)

	defer db.Close()
}
