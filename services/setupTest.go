package services

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
)

func SetupTest() (*Service, *sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	s := NewService(db)
	return s, db, mock
}
