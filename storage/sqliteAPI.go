package storage

import (
	"database/sql"
	"fmt"
	"gin-twitter/mapping"
	"gin-twitter/types"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteStorage struct {
	sqlDb *sql.DB
}

func NewSqliteStorage(dbPath string) (*SqliteStorage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return &SqliteStorage{sqlDb: db}, nil
}

// TODO: fix opening and closing of database
func (s *SqliteStorage) CreateTable(tableName, schema string) error {
	log.Println("Entering CreateTable")
	_, err := s.sqlDb.Exec(schema)
	if err != nil {
		log.Printf("Error in creating table: %v\n", err)
	}
	log.Printf("Table %s was created!", tableName)
	return nil
}

func (s *SqliteStorage) Get(tableName, whereCondition string, args ...interface{}) (interface{}, error) {
	row := s.sqlDb.QueryRow("SELECT * FROM "+tableName+" WHERE "+whereCondition, args...)

	if err := row.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	person, mapErr := mapping.MapPerson(row)

	if mapErr != nil {
		return nil, mapErr
	}

	log.Println(person)
	return person, nil
}

func (s *SqliteStorage) Create(tableName string, data interface{}) error {

	log.Println("Creating person in database!")
	if tableName != "persons" {
		fmt.Print("Wrong table name")
		return nil
	}
	var person *types.Person
	var ok bool
	person, ok = data.(*types.Person)

	if !ok {
		return fmt.Errorf("Invalid data")
	}

	insertSQL := `
		INSERT INTO persons (firstname, lastname, username, verified, joined)
		VALUES (?, ?, ?, ?, ?)
	`
	t := time.Now()

	_, err := s.sqlDb.Exec(
		insertSQL,
		person.FirstName,
		person.LastName,
		person.UserName,
		person.Verified,
		t,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Person with username: %s at time %s, ID : %d\n", person.UserName, person.Joined, person.ID)
	return nil
}

func (s *SqliteStorage) Close() {
	if s.sqlDb != nil {
		if err := s.sqlDb.Close(); err != nil {
			log.Printf("Error closing database: %v\n", err)
		}
	}
}

func (s *SqliteStorage) Update(tableName, setClause, whereCondition string, args ...interface{}) error {
	updateSQL := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, setClause, whereCondition)

	_, err := s.sqlDb.Exec(updateSQL, args...)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Record in table %s updated successfully", tableName)
	return nil
}

func (s *SqliteStorage) Delete(tableName, whereCondition string, args ...interface{}) error {
	deleteSQL := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, whereCondition)

	_, err := s.sqlDb.Exec(deleteSQL, args...)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Delete person succes!")
	return nil
}
