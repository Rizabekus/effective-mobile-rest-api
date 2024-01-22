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
	if err != nil {
		http.Error(w, "Failed to get people", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		http.Error(w, "Invalid 'page' parameter", http.StatusBadRequest)
		return
	}
	fmt.Println(people)
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if err != nil || pageSize < 1 {
		http.Error(w, "Invalid 'pageSize' parameter", http.StatusBadRequest)
		return
	}
	if page >= 1 && pageSize >= 1 {
		people = handler.Service.PersonService.Pagination(page, pageSize, people)
	}
	// if people == nil {
	// 	http.Error(w, "Invalid 'page' parameter", http.StatusBadRequest)
	// 	return
	// }
	jsonData, err := json.Marshal(people)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}
