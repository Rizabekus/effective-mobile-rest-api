package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
	"github.com/Rizabekus/effective-mobile-rest-api/pkg/loggers.go"
	"github.com/gorilla/mux"
)

func (handler *Handlers) GetPerson(w http.ResponseWriter, r *http.Request) {
	loggers.DebugLog.Println("Received a request to GetPerson")
	vars := mux.Vars(r)
	personID := vars["id"]
	loggers.DebugLog.Println("Received an ID")
	exist, err := handler.Service.PersonService.DoesExistByID(personID)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Failed to check person existence",
			Error: err.Error(),
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)
		return
	}

	if exist {
		loggers.DebugLog.Println("Checked ID for existence")
		person, err := handler.Service.PersonService.GetPersonByID(personID)
		if err != nil {
			response := models.ResponseStructure{
				Field: "Failed to get person",
				Error: err.Error(),
			}
			handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)
			return
		}

		personJSON, err := json.Marshal(person)
		if err != nil {
			response := models.ResponseStructure{
				Field: "Failed to encode person JSON",
				Error: err.Error(),
			}
			handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(personJSON)
		loggers.DebugLog.Println("GetPerson completed successfully")
	} else {
		response := models.ResponseStructure{
			Field: "Person doesn't exist",
			Error: "",
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusNoContent)
	}
}
