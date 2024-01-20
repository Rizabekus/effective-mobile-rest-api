package handlers

import (
	"encoding/json"
	"net/http"
)

func (handler *Handlers) GetPeople(w http.ResponseWriter, r *http.Request) {
	people, err := handler.Service.PersonService.GetPeople()
	if err != nil {
		http.Error(w, "Failed to get people", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(people)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}
