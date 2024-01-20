package services

import "github.com/Rizabekus/effective-mobile-rest-api/internal/models"

type PersonService struct {
	storage models.PersonStorage
}

func CreatePersonService(storage models.PersonStorage) *PersonService {
	return &PersonService{storage: storage}
}

func (PersonService *PersonService) GetPeople() ([]models.Person, error) {
	return PersonService.storage.GetPeople()
}
