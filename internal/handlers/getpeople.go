package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
	"github.com/Rizabekus/effective-mobile-rest-api/pkg/loggers.go"
)

func (handler *Handlers) GetPeople(w http.ResponseWriter, r *http.Request) {
	loggers.DebugLog.Println("Received a request to GetPeople")
	queryParams := r.URL.Query()
	loggers.DebugLog.Println("Received query parameters")
	people, err := handler.Service.PersonService.FilteredSearch(queryParams)

	loggers.DebugLog.Println("Filtered search results")

	if err != nil {
		errorMsg := fmt.Sprintf("Failed to get people: %s", err.Error())
		response := models.ResponseStructure{
			Field: "Failed to get people",
			Error: errorMsg,
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.InfoLog.Println(errorMsg)
		return
	}

	pageQuery := r.URL.Query().Get("page")
	var page int
	if pageQuery == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			errorMsg := fmt.Sprintf("Invalid 'page' parameter: %s", err.Error())
			response := models.ResponseStructure{
				Field: "Invalid 'page' parameter",
				Error: errorMsg,
			}
			handler.Service.PersonService.SendResponse(response, w, http.StatusBadRequest)

			loggers.InfoLog.Println(errorMsg)
			return
		}
	}

	pageSizeQuery := r.URL.Query().Get("pageSize")
	var pageSize int
	if pageSizeQuery == "" {
		pageSize = 15
	} else {
		pageSize, err = strconv.Atoi(pageSizeQuery)
		if err != nil {
			errorMsg := fmt.Sprintf("Invalid 'pageSize' parameter: %s", err.Error())
			response := models.ResponseStructure{
				Field: "Invalid 'pageSize' parameter",
				Error: errorMsg,
			}
			handler.Service.PersonService.SendResponse(response, w, http.StatusBadRequest)

			loggers.InfoLog.Println(errorMsg)
			return
		}
	}

	if page >= 1 && pageSize >= 1 {
		people = handler.Service.PersonService.Pagination(page, pageSize, people)
	}
	loggers.DebugLog.Println("Successfully made pagination")
	jsonData, err := json.Marshal(people)
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to encode JSON: %s", err.Error())
		response := models.ResponseStructure{
			Field: "Failed to encode JSON",
			Error: errorMsg,
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.InfoLog.Println(errorMsg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	loggers.DebugLog.Println("GetPeople completed successfully")
}
