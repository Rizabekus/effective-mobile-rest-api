package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

func (handler *Handlers) UpdatePerson(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	personID := vars["id"]
	exist, err := handler.Service.PersonService.DoesExistByID(personID)
	if err != nil {
		http.Error(w, "Doesn't exist", http.StatusNoContent)
		return
	}
	if exist {
		var updatedValues models.UpdatedPerson
		err := json.NewDecoder(r.Body).Decode(&updatedValues)
		if err != nil {
			http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
			return
		}
		validate := validator.New()
		err = validate.Struct(updatedValues)

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

		err = handler.Service.PersonService.UpdatePerson(updatedValues, personID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Person updated successfully"))

	}
}
