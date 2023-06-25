package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/57th-street/oza-gueser/models"
)

var baseUrl = "http://localhost:8080/"

func TestRegisterController(t *testing.T) {
	newUser := models.User{
		Email:    "test@test.com",
		Password: "test1234",
	}
	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(newUser); err != nil {
		t.Error(err)
	}
	req := httptest.NewRequest(http.MethodPost, baseUrl+"register", buffer)
	res := httptest.NewRecorder()
	aCon.RegisterController(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("unexpected StatusCode: got %d but want %d", res.Code, http.StatusOK)
	}
}

func TestLoginController(t *testing.T) {
	loginUser := models.User{
		Email:    "test@test.com",
		Password: "test1234",
	}
	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(loginUser); err != nil {
		t.Error(err)
	}
	req := httptest.NewRequest(http.MethodPost, baseUrl+"login", buffer)
	res := httptest.NewRecorder()
	aCon.LoginController(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("unexpected StatusCode: got %d but want %d", res.Code, http.StatusOK)
	}
}
