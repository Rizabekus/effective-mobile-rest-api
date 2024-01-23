package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/handlers"
	"github.com/Rizabekus/effective-mobile-rest-api/pkg/loggers.go"
	"github.com/gorilla/mux"
)

func Routes(h *handlers.Handlers) {
	r := mux.NewRouter()

	r.HandleFunc("/people", h.GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", h.GetPerson).Methods("GET")
	r.HandleFunc("/people", h.CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", h.UpdatePerson).Methods("PUT")
	r.HandleFunc("/people/{id}", h.DeletePerson).Methods("DELETE")
	fmt.Println("http://localhost:8000")
	loggers.InfoLog.Println("Started the server")
	log.Fatal(http.ListenAndServe(":8000", r))

}
