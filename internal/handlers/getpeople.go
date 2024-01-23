package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (handler *Handlers) GetPeople(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	people, err := handler.Service.PersonService.FilteredSearch(queryParams)
	fmt.Println(people)
	if err != nil {
		http.Error(w, "Failed to get people", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	pageQuery := r.URL.Query().Get("page")
	var page int
	if pageQuery == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			http.Error(w, "Invalid 'page' parameter", http.StatusBadRequest)
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
			http.Error(w, "Invalid 'pageSzie' parameter", http.StatusBadRequest)
			return
		}
	}

	if page >= 1 && pageSize >= 1 {
		people = handler.Service.PersonService.Pagination(page, pageSize, people)
	}

	jsonData, err := json.Marshal(people)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}
