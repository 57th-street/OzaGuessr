package repositories_test

import (
	"testing"

	"github.com/57th-street/oza-gueser/repositories"
	"github.com/57th-street/oza-gueser/repositories/testdata"
)

func TestCheckUserExist(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "normal case test",
			email:    testdata.UserTestData.Email,
			expected: true,
		},
		{
			name:     "edge case test",
			email:    "test@gmail.com",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := repositories.CheckUserExist(testDB, test.email)
			if err != nil {
				t.Fatal(err)
			}
			if result != test.expected {
				t.Errorf("got %v but want %v", result, test.expected)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	email := "test2@test.com"
	password := "hashedPassword2"
	err := repositories.CreateUser(testDB, email, password)
	if err != nil {
		t.Error(err)
	}
	var actualCount int
	expectedCount := 1
	sqlStr := "select count(*) from users where email = ? and password = ?"
	err = testDB.QueryRow(sqlStr, email, password).Scan(&actualCount)
	if actualCount != expectedCount {
		t.Errorf("got %d but want %d", actualCount, expectedCount)
	}
	t.Cleanup(func() {
		const sqlStr = "delete from users where email = ?"
		testDB.Exec(sqlStr, email)
	})
}

func TestGetUserPassword(t *testing.T) {
	email := testdata.UserTestData.Email
	password := testdata.UserTestData.Password
	gotPassword, err := repositories.GetUserPassword(testDB, email)
	if err != nil {
		t.Error(err)
	}
	if gotPassword != password {
		t.Errorf("got %s but want %s", gotPassword, password)
	}
}
