package storage

import "database/sql"

type UserDB struct {
	DB *sql.DB
}

func CreatePersonStorage(db *sql.DB) *UserDB {
	return &UserDB{DB: db}
}
