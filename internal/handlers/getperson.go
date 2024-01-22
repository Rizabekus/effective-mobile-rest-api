package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (handler *Handlers) GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personID := vars["id"]
	exist, err := handler.Service.PersonService.DoesExistByID(personID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if exist {
		person, err := handler.Service.PersonService.GetPersonByID(personID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		personJSON, err := json.Marshal(person)
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(personJSON)
	} else {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Person doesn't exist"))
	}
}
