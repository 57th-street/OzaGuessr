package services_test

import (
	"database/sql"
	"testing"

	"github.com/57th-street/oza-gueser/repositories/testdata"
	"github.com/57th-street/oza-gueser/services"
	"github.com/57th-street/oza-gueser/utils"
	"github.com/DATA-DOG/go-sqlmock"
)

var (
	email    = testdata.UserTestData.Email
	password = testdata.UserTestData.Password
)

func TestRegister(t *testing.T) {
	s, db, mock := services.SetupTest()
	defer db.Close()

	tests := []struct {
		name       string
		userExists bool
		errIs      bool
	}{
		{"successRegister", false, false},
		{"userAlreadyExists", true, true},
	}
	sqlStrCheckUserExist := "select exists"
	sqlStrCreateUser := "insert into users"

	for _, tt := range tests {
		mock.ExpectQuery(sqlStrCheckUserExist).WithArgs(email).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(tt.userExists))
		if !tt.userExists {
			mock.ExpectExec(sqlStrCreateUser).WithArgs(email, password).WillReturnResult(sqlmock.NewResult(1, 1))
		}
		err := s.Register(email, password)
		if !tt.errIs && err != nil {
			t.Errorf("unexpected error: %v", err)
		} else if tt.errIs && err == nil {
			t.Errorf("error was expected but got nil")
		}
	}
}

func TestLogin(t *testing.T) {
	s, db, mock := services.SetupTest()
	defer db.Close()

	tests := []struct {
		name   string
		noData bool
		errIs  bool
	}{
		{"successLogin", false, false},
		{"noData", true, true},
	}
	sqlStrGetUserPassword := "select password from users"
	hashedPassword, _ := utils.HashPassword(password)

	for _, tt := range tests {
		if !tt.noData {
			mock.ExpectQuery(sqlStrGetUserPassword).WithArgs(email).WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(hashedPassword))
		} else {
			mock.ExpectQuery(sqlStrGetUserPassword).WithArgs(email).WillReturnError(sql.ErrNoRows)
		}
		err := s.Login(email, password)
		if !tt.errIs && err != nil {
			t.Errorf("unexpected error: %v", err)
		} else if tt.errIs && err == nil {
			t.Errorf("error was expected but got nil")
		}
	}
}
