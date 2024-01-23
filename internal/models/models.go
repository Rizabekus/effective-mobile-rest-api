package models

import (
	"net/http"
	"net/url"
)

type PersonService interface {
	GetPeople() ([]Person, error)
	AddPerson(person Person) error
	DoesExist(person Person) (bool, error)
	DoesExistByID(id string) (bool, error)
	DeleteByID(id string) error
	GetPersonByID(id string) (Person, error)
	UpdatePerson(UpdatedPerson, string) error
	FilteredSearch(queries url.Values) ([]Person, error)
	Pagination(int, int, []Person) []Person
	SendResponse(ResponseStructure, http.ResponseWriter, int)
}
type PersonStorage interface {
	GetPeople() ([]Person, error)
	AddPerson(person Person) error
	DoesExist(person Person) (bool, error)
	DoesExistByID(id string) (bool, error)
	DeleteByID(id string) error
	GetPersonByID(id string) (Person, error)
	UpdatePerson(UpdatedPerson, string) error
	FilteredSearch(queries url.Values) ([]Person, error)
}
type Person struct {
	Name        string `json:"name" validate:"required,min=1,max=50,alpha"`
	Surname     string `json:"surname" validate:"required,min=1,max=50,alpha"`
	Patronymic  string `json:"patronymic" validate:"required,min=1,max=50,alpha"`
	Gender      string `json:"gender" validate:"omitempty,eq=Male|eq=Female|eq=male|eq=female"`
	Nationality string `json:"nationality" validate:"omitempty,alpha,len=2"`
	Age         int    `json:"age" validate:"omitempty,min=0,max=150"`
}
type UpdatedPerson struct {
	Name        string `json:"name" validate:"omitempty,min=1,max=50,alpha"`
	Surname     string `json:"surname" validate:"omitempty,min=1,max=50,alpha"`
	Patronymic  string `json:"patronymic" validate:"omitempty,min=1,max=50,alpha"`
	Gender      string `json:"gender" validate:"omitempty,eq=Male|eq=Female|eq=male|eq=female"`
	Nationality string `json:"nationality" validate:"omitempty,alpha,len=2"`
	Age         int    `json:"age" validate:"omitempty,min=0,max=150"`
}
type ResponseStructure struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
