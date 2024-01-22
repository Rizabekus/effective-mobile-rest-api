package storage

import (
	"database/sql"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

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
func (PersonDB *PersonDB) AddPerson(p models.Person) error {

	sqlStatement := `
		INSERT INTO people (name, surname, patronymic, gender, nationality, age)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := PersonDB.DB.Exec(sqlStatement, p.Name, p.Surname, p.Patronymic, p.Gender, p.Nationality, p.Age)
	if err != nil {

		return fmt.Errorf("error executing SQL statement in AddPerson: %v", err)
	}

	return nil
}
func (PersonDB *PersonDB) DoesExist(person models.Person) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1 FROM people
			WHERE name = $1 AND surname = $2 AND patronymic = $3
		)
	`

	var exists bool
	err := PersonDB.DB.QueryRow(query, person.Name, person.Surname, person.Patronymic).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking existence: %v", err)
	}

	return exists, nil
}
func (PersonDB *PersonDB) DoesExistByID(id string) (bool, error) {
	if id == "" {
		return false, fmt.Errorf("empty ID provided")
	}

	query := "SELECT EXISTS(SELECT 1 FROM people WHERE id = $1)"

	var exists bool
	err := PersonDB.DB.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking existence: %v", err)
	}

	return exists, nil
}
func (PersonDB *PersonDB) DeleteByID(id string) error {
	query := "DELETE FROM people WHERE id = $1"

	_, err := PersonDB.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting person: %v", err)
	}

	return nil
}
func (PersonDB *PersonDB) GetPersonByID(id string) (models.Person, error) {
	var person models.Person

	row := PersonDB.DB.QueryRow("SELECT name, surname, patronymic, gender, nationality, age FROM people WHERE id = $1", id)

	err := row.Scan(&person.Name, &person.Surname, &person.Patronymic, &person.Gender, &person.Nationality, &person.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return person, fmt.Errorf("person not found with ID %s", id)
		}
		return person, fmt.Errorf("failed to scan person data: %w", err)
	}

	return person, nil
}

func (PersonDB *PersonDB) UpdatePerson(update models.UpdatedPerson, personID string) error {
	var setValues []string
	var values []interface{}

	typ := reflect.TypeOf(update)

	val := reflect.ValueOf(update)

	for i := 0; i < typ.NumField(); i++ {

		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface()

		if reflect.Zero(typ.Field(i).Type).Interface() == fieldValue {
			continue
		}

		setValues = append(setValues, fmt.Sprintf("%s = $%d", strings.ToLower(fieldName), len(values)+1))
		values = append(values, fieldValue)
	}

	if len(setValues) == 0 {

		return nil
	}

	query := fmt.Sprintf("UPDATE people SET %s WHERE id = $%d", strings.Join(setValues, ", "), len(values)+1)

	stmt, err := PersonDB.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	values = append(values, personID)

	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}

	return nil
}
func (PersonDB *PersonDB) FilteredSearch(queries url.Values) ([]models.Person, error) {
	var filteredPeople []models.Person

	query := "SELECT * FROM people WHERE true"
	params := make([]interface{}, 0)

	if name, ok := queries["name"]; ok && len(name) > 0 {
		query += " AND name = $1"
		params = append(params, name[0])
	}

	if surname, ok := queries["surname"]; ok && len(surname) > 0 {
		query += " AND surname = $" + strconv.Itoa(len(params)+1)
		params = append(params, surname[0])
	}

	if patronymic, ok := queries["patronymic"]; ok && len(patronymic) > 0 {
		query += " AND patronymic = $" + strconv.Itoa(len(params)+1)
		params = append(params, patronymic[0])
	}

	if gender, ok := queries["gender"]; ok && len(gender) > 0 {
		query += " AND gender = $" + strconv.Itoa(len(params)+1)
		params = append(params, gender[0])
	}

	if nationality, ok := queries["nationality"]; ok && len(nationality) > 0 {
		query += " AND nationality = $" + strconv.Itoa(len(params)+1)
		params = append(params, nationality[0])
	}

	var minAgeInt int
	if minAge, ok := queries["minAge"]; ok && len(minAge) > 0 {
		tmp, err := strconv.Atoi(minAge[0])
		minAgeInt = tmp
		if err == nil {
			query += " AND age >= $" + strconv.Itoa(len(params)+1)
			params = append(params, minAgeInt)
		}
	}

	var maxAgeInt int
	if maxAge, ok := queries["maxAge"]; ok && len(maxAge) > 0 {
		tmp, err := strconv.Atoi(maxAge[0])
		maxAgeInt = tmp
		if err == nil && maxAgeInt >= minAgeInt {
			query += " AND age <= $" + strconv.Itoa(len(params)+1)
			params = append(params, maxAgeInt)
		}
	}

	rows, err := PersonDB.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var person models.Person
		var id sql.NullInt64
		err := rows.Scan(
			&id,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Gender,
			&person.Nationality,
			&person.Age,
		)
		if err != nil {
			return nil, err
		}
		filteredPeople = append(filteredPeople, person)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return filteredPeople, nil
}
