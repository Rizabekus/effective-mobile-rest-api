package services

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
	"github.com/Rizabekus/effective-mobile-rest-api/pkg/external_api"
)

type PersonService struct {
	storage models.PersonStorage
}

func CreatePersonService(storage models.PersonStorage) *PersonService {
	return &PersonService{storage: storage}
}

func (PersonService *PersonService) GetPeople() ([]models.Person, error) {
	return PersonService.storage.GetPeople()
}

func (PersonService *PersonService) AddPerson(data models.Person) error {
	data.Age = external_api.Agify(data.Name)
	data.Gender = external_api.Genderize(data.Name)
	data.Nationality = external_api.Nationalize(data.Name)
	return PersonService.storage.AddPerson(data)
}
func (PersonService *PersonService) DoesExist(data models.Person) (bool, error) {
	return PersonService.storage.DoesExist(data)
}
func (PersonService *PersonService) DoesExistByID(id string) (bool, error) {
	return PersonService.storage.DoesExistByID(id)
}
func (PersonService *PersonService) DeleteByID(id string) error {
	return PersonService.storage.DeleteByID(id)
}

func (PersonService *PersonService) GetPersonByID(id string) (models.Person, error) {
	return PersonService.storage.GetPersonByID(id)
}

func (PersonService *PersonService) UpdatePerson(update models.UpdatedPerson, personID string) error {
	return PersonService.storage.UpdatePerson(update, personID)
}
func (PersonService *PersonService) FilteredSearch(queries url.Values) ([]models.Person, error) {
	return PersonService.storage.FilteredSearch(queries)
}
func (PersonService *PersonService) Pagination(page, pageSize int, people []models.Person) []models.Person {
	startIdx := (page - 1) * pageSize
	endIdx := startIdx + pageSize

	if startIdx >= len(people) {
		return nil
	}
	if endIdx > len(people) {
		endIdx = len(people)
	}

	return people[startIdx:endIdx]
}

func (PersonService *PersonService) SendResponse(response models.ResponseStructure, w http.ResponseWriter, statusCode int) {
	responseJSON, err := json.Marshal(response)
	if err != nil {

		internalError := models.ResponseStructure{
			Field: "Internal Server Error",
			Error: "Failed to marshal JSON response",
		}
		internalErrorJSON, _ := json.Marshal(internalError)

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, string(internalErrorJSON), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
