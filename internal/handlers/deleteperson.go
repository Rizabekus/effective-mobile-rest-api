package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (handler *Handlers) DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personID := vars["id"]

	exist, err := handler.Service.PersonService.DoesExistByID(personID)
	if err != nil {
		http.Error(w, "Failed to get information", http.StatusInternalServerError)
		return
	}
	if exist {
		err = handler.Service.PersonService.DeleteByID(personID)
		if err != nil {
			http.Error(w, "Failed to get information", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Person deleted successfully"))
	} else {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Person doesn't exist"))
	}
}
