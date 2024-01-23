package handlers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
	"github.com/Rizabekus/effective-mobile-rest-api/pkg/loggers.go"
	"gopkg.in/go-playground/validator.v9"
)

func (handler *Handlers) CreatePerson(w http.ResponseWriter, r *http.Request) {
	loggers.DebugLog.Println("Received a request to CreatePerson")

	var newPerson models.Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Failed to decode JSON",
			Error: err.Error(),
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusBadRequest)

		loggers.InfoLog.Println("Failed to decode JSON")
		return
	}

	loggers.DebugLog.Println("Decoded JSON")

	validate := validator.New()
	err = validate.StructPartial(newPerson, "Name", "Surname", "Patronymic")
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			response := models.ResponseStructure{
				Field: "Internal Server Error",
				Error: err.Error(),
			}
			handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)

			loggers.InfoLog.Println("Internal Server Error")
			return
		}

		var errors []models.ResponseStructure
		for _, ve := range validationErrors {
			response := models.ResponseStructure{
				Field: fmt.Sprintf("Field: %s, Tag: %s\n", ve.Field(), ve.Tag()),
				Error: err.Error(),
			}
			errors = append(errors, response)
		}

		handler.Service.PersonService.SendResponse(errors[0], w, http.StatusBadRequest)

		loggers.InfoLog.Println("Validation Error")
		return
	}

	exist, err := handler.Service.PersonService.DoesExist(newPerson)
	if exist {
		loggers.DebugLog.Println("Person already exists")

		response := models.ResponseStructure{
			Field: "Person already exists",
			Error: "",
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusNoContent)
		return
	}
	if err != nil {
		response := models.ResponseStructure{
			Field: "Failed to get information",
			Error: err.Error(),
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.InfoLog.Println("Failed to get information")
		return
	}

	err = handler.Service.PersonService.AddPerson(newPerson)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Failed to add",
			Error: err.Error(),
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.InfoLog.Println("Failed to add person")
		return
	}

	loggers.DebugLog.Println("Person created successfully")

	response := models.ResponseStructure{
		Field: "Person created successfully",
		Error: "",
	}
	handler.Service.PersonService.SendResponse(response, w, http.StatusCreated)
	return
}
