package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
	"github.com/Rizabekus/effective-mobile-rest-api/pkg/loggers.go"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

func (handler *Handlers) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	loggers.DebugLog.Print("Received a request to UpdatePerson")
	vars := mux.Vars(r)
	personID := vars["id"]
	loggers.DebugLog.Print("Received an ID")

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
		loggers.DebugLog.Print("Checked ID for existence")

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

		loggers.DebugLog.Print("Decoded JSON")

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

			loggers.InfoLog.Print("Failed to update person")
			return
		}

		loggers.DebugLog.Print("Person updated successfully")

		response := models.ResponseStructure{
			Field: "Person updated successfully",
			Error: "",
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusAccepted)
	}
}
