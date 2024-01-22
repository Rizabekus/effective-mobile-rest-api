package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/handlers"
	"github.com/gorilla/mux"
)

func Routes(h *handlers.Handlers) {
	r := mux.NewRouter()
	r.Use(LoggerMiddleware)
	r.HandleFunc("/people", h.GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", h.GetPerson).Methods("GET")
	r.HandleFunc("/people", h.CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", h.UpdatePerson).Methods("PUT")
	r.HandleFunc("/people/{id}", h.DeletePerson).Methods("DELETE")
	fmt.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		logLevel := os.Getenv("LOG_LEVEL")

		switch logLevel {
		case "DEBUG":
			log.Printf(
				"[DEBUG] %s %s %s %s",
				r.Method,
				r.RequestURI,
				r.RemoteAddr,
				time.Since(start),
			)
		case "INFO":
			log.Printf(
				"[INFO] %s %s %s %s",
				r.Method,
				r.RequestURI,
				r.RemoteAddr,
				time.Since(start),
			)
		default:
			// Если уровень логирования не задан, выводим только общую информацию
			log.Printf(
				"%s %s %s %s",
				r.Method,
				r.RequestURI,
				r.RemoteAddr,
				time.Since(start),
			)
		}
	})
}
