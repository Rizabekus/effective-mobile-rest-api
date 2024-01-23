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
		response := models.ResponseStructure{
			Field: "Doesn't exist",
			Error: err.Error(),
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusNoContent)
		return
	}

	if exist {
		var updatedValues models.UpdatedPerson
		err := json.NewDecoder(r.Body).Decode(&updatedValues)
		if err != nil {
			response := models.ResponseStructure{
				Field: "Failed to decode JSON",
				Error: err.Error(),
			}
			handler.Service.PersonService.SendResponse(response, w, http.StatusBadRequest)
			return
		}

		validate := validator.New()
		err = validate.Struct(updatedValues)
		if err != nil {
			validationErrors, ok := err.(validator.ValidationErrors)
			if !ok {
				response := models.ResponseStructure{
					Field: "Internal Server Error",
					Error: err.Error(),
				}
				handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)
				return
			}

			for _, ve := range validationErrors {
				response := models.ResponseStructure{
					Field: fmt.Sprintf("Field: %s, Tag: %s\n", ve.Field(), ve.Tag()),
					Error: err.Error(),
				}

				handler.Service.PersonService.SendResponse(response, w, http.StatusBadRequest)
				return
			}
		}

		err = handler.Service.PersonService.UpdatePerson(updatedValues, personID)
		if err != nil {
			response := models.ResponseStructure{
				Field: "Internal Server Error",
				Error: err.Error(),
			}
			handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		response := models.ResponseStructure{
			Field: "Person updated successfully",
			Error: "",
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusNoContent)
	}
}
