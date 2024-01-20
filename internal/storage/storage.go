package storage

import (
	"database/sql"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
)

type PersonStorage struct {
	Storage models.PersonStorage
}

func StorageInstance(db *sql.DB) *PersonStorage {
	return &PersonStorage{Storage: CreatePersonStorage(db)}
}
