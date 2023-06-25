package repositories_test

import (
	"testing"

	"github.com/57th-street/oza-gueser/models"
	"github.com/57th-street/oza-gueser/repositories"
)

func TestCheckUserExist(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "normal case test",
			email:    "test@test.com",
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

func TestRegister(t *testing.T) {
	user := models.User{
		Email:    "test2@test.com",
		Password: "test1234",
	}
	expectedId := 2
	newUser, err := repositories.Register(testDB, user)
	if err != nil {
		t.Error(err)
	}
	if expectedId != newUser.ID {
		t.Errorf("got %d but want %d", newUser.ID, expectedId)
	}
	t.Cleanup(func() {
		const sqlStr = "delete from users where email = ?"
		testDB.Exec(sqlStr, user.Email)
	})
}

func TestLogin(t *testing.T) {
	user := models.User{
		Email:    "test@test.com",
		Password: "test1234",
	}
	expectedId := 1
	loginInUser, err := repositories.Login(testDB, user)
	if err != nil {
		t.Error(err)
	}
	if loginInUser.ID != expectedId {
		t.Errorf("got %d but want %d", loginInUser.ID, expectedId)
	}
}
