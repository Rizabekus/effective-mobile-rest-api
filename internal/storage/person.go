package storage

import (
	"database/sql"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
)

type PersonDB struct {
	DB *sql.DB
}

func CreatePersonStorage(db *sql.DB) *PersonDB {
	return &PersonDB{DB: db}
}
func (PersonDB *PersonDB) GetPeople() ([]models.Person, error) {
	return []models.Person{}, nil
}
