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
	result, err := repositories.CreateUser(testDB, email, password)
	if err != nil {
		t.Error(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Error(err)
	}
	expectedID := 2
	if int(id) != expectedID {
		t.Errorf("got %d but want %d", id, expectedID)
	}
	t.Cleanup(func() {
		const sqlStr = "delete from users where email = ?"
		testDB.Exec(sqlStr, email)
	})
}

func TestGetUser(t *testing.T) {
	id := testdata.UserTestData.ID
	email := testdata.UserTestData.Email
	password := testdata.UserTestData.Password
	user, err := repositories.GetUser(testDB, email)
	if err != nil {
		t.Error(err)
	}
	if int(user.ID) != id {
		t.Errorf("got %d but want %d", user.ID, id)
	}
	if user.Email != email {
		t.Errorf("got %s but want %s", user.Email, email)
	}
	if user.HashedPassword != password {
		t.Errorf("got %s but want %s", user.HashedPassword, password)
	}
}
