package handlers

import (
	"net/http"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
	"github.com/gorilla/mux"
)

func (handler *Handlers) DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personID := vars["id"]

	exist, err := handler.Service.PersonService.DoesExistByID(personID)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Failed to get information",
			Error: err.Error(),
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)
		return
	}

	var response models.ResponseStructure
	if exist {
		err = handler.Service.PersonService.DeleteByID(personID)
		if err != nil {
			response = models.ResponseStructure{
				Field: "Failed to get information",
				Error: err.Error(),
			}
			handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)
			return
		}
		response = models.ResponseStructure{
			Field: "Person deleted successfully",
			Error: "",
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusCreated)
	} else {
		response = models.ResponseStructure{
			Field: "Person doesn't exist",
			Error: "",
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusNoContent)
	}
}
