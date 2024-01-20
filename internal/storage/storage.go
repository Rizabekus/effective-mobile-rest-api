package storage

import (
	"database/sql"

	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
)

type Storage struct {
	PersonStorage models.PersonStorage
}

func StorageInstance(db *sql.DB) *Storage {
	return &Storage{PersonStorage: CreatePersonStorage(db)}
}
