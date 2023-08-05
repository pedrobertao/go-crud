package database

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/pedrobertao/go-crud/models"
)

/*

In this section, you can define the database setup and creation. You also have the option to create subfolders for organizing query handlers. For demonstration purposes, a simple slice implementation will be used.
By organizing query handlers into separate files, you achieve a clean and modular codebase. This approach ensures a clear separation of concerns, making maintenance and scalability easier as your project evolves.
*/

type DBData struct {
	ID  string `json:"id"`
	DBB []byte `json:"dbb"`
}

var myDB []DBData

func Start() {
	myDB = make([]DBData, 0)
}

func FindAll() ([]models.Person, error) {
	people := make([]models.Person, 0)

	if len(myDB) == 0 {
		return people, nil
	}

	for _, obj := range myDB {
		var person models.Person
		if err := json.Unmarshal(obj.DBB, &person); err != nil {
			return people, err
		}
		person.ID = obj.ID
		people = append(people, person)
	}

	return people, nil
}

func FindById(id string) (models.Person, error) {
	for _, obj := range myDB {
		var person models.Person
		if err := json.Unmarshal(obj.DBB, &person); err != nil {
			return person, err
		}
		if person.ID == id {
			person.ID = obj.ID
			return person, nil
		}
	}

	return models.Person{}, nil
}

func Post(p models.Person) (string, error) {
	buff, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	id := uuid.New()
	uid := id.String()
	myDB = append(myDB, DBData{
		ID:  uid,
		DBB: buff,
	})
	return uid, nil
}

func DeleteByID(id string) int {
	newDB := make([]DBData, 0)
	count := 0
	for _, obj := range myDB {
		if obj.ID != id {
			newDB = append(newDB, obj)
		} else {
			count++
		}
	}
	myDB = newDB
	return count
}
