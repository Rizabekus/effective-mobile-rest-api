package handlers

import (
	"net/http"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
	"github.com/Rizabekus/effective-mobile-rest-api/pkg/loggers.go"
	"github.com/gorilla/mux"
)

func (handler *Handlers) DeletePerson(w http.ResponseWriter, r *http.Request) {
	loggers.DebugLog.Println("Received a request to DeletePerson")

	vars := mux.Vars(r)
	personID := vars["id"]

	exist, err := handler.Service.PersonService.DoesExistByID(personID)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Failed to get information",
			Error: err.Error(),
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.InfoLog.Println("Failed to get information")
		return
	}

	var response models.ResponseStructure
	if exist {
		loggers.DebugLog.Println("Person exists - attempting to delete")

		err = handler.Service.PersonService.DeleteByID(personID)
		if err != nil {
			response = models.ResponseStructure{
				Field: "Failed to delete person",
				Error: err.Error(),
			}
			handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)

			loggers.InfoLog.Println("Failed to delete person")
			return
		}

		loggers.DebugLog.Println("Person deleted successfully")

		response = models.ResponseStructure{
			Field: "Person deleted successfully",
			Error: "",
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusCreated)
	} else {
		loggers.DebugLog.Println("Person doesn't exist")

		response = models.ResponseStructure{
			Field: "Person doesn't exist",
			Error: "",
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusNoContent)
	}
}
