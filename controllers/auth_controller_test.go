package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/57th-street/oza-gueser/controllers"
	"github.com/57th-street/oza-gueser/repositories/testdata"
)

// 環境によってurlが変わるため今後改善の余地あり
var baseUrl = "http://localhost:8080/"
var userInput = controllers.UserInput{
	Email:    testdata.UserTestData.Email,
	Password: testdata.UserTestData.Password,
}

func TestRegisterController(t *testing.T) {
	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(userInput); err != nil {
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
	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(userInput); err != nil {
		t.Error(err)
	}
	req := httptest.NewRequest(http.MethodPost, baseUrl+"login", buffer)
	res := httptest.NewRecorder()
	aCon.LoginController(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("unexpected StatusCode: got %d but want %d", res.Code, http.StatusOK)
	}
}
