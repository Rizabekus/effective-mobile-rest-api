package models

type PersonService interface {
	GetPeople() ([]Person, error)
}
type PersonStorage interface {
	GetPeople() ([]Person, error)
}
type Person struct {
	Name        string
	Surname     string
	Patryonomic string
	Gender      string
	Nationality string
	Age         string
}
