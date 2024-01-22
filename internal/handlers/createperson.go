package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
	"gopkg.in/go-playground/validator.v9"
)

func (handler *Handlers) CreatePerson(w http.ResponseWriter, r *http.Request) {

	var newPerson models.Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.StructPartial(newPerson, "Name", "Surname", "Patronymic")
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		for _, ve := range validationErrors {
			response := models.ErrorResponse{
				Field: fmt.Sprintf("Field: %s, Tag: %s\n", ve.Field(), ve.Tag()),
				Error: err.Error(),
			}

			responseJSON, _ := json.Marshal(response)

			w.Header().Set("Content-Type", "application/json")
			http.Error(w, string(responseJSON), http.StatusBadRequest)
			return

		}

	}
	exist, err := handler.Service.PersonService.DoesExist(newPerson)
	if exist {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Person already exists"))
		return
	}
	if err != nil {
		http.Error(w, "Failed to get information", http.StatusInternalServerError)

		return
	}

	err = handler.Service.PersonService.AddPerson(newPerson)
	if err != nil {
		http.Error(w, "Failed to add", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Person created successfully"))
}
