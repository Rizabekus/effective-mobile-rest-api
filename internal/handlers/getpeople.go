package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
)

func (handler *Handlers) GetPeople(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	people, err := handler.Service.PersonService.FilteredSearch(queryParams)
	fmt.Println(people)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Failed to get people",
			Error: err.Error(),
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)

		return
	}

	pageQuery := r.URL.Query().Get("page")
	var page int
	if pageQuery == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			response := models.ResponseStructure{
				Field: "Invalid 'page' parameter",
				Error: err.Error(),
			}
			handler.Service.PersonService.SendResponse(response, w, http.StatusBadRequest)
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
			response := models.ResponseStructure{
				Field: "Invalid 'pageSize' parameter",
				Error: err.Error(),
			}
			handler.Service.PersonService.SendResponse(response, w, http.StatusBadRequest)
			return
		}
	}

	if page >= 1 && pageSize >= 1 {
		people = handler.Service.PersonService.Pagination(page, pageSize, people)
	}

	jsonData, err := json.Marshal(people)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Failed to encode JSON",
			Error: err.Error(),
		}
		handler.Service.PersonService.SendResponse(response, w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
