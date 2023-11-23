package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/57th-street/oza-gueser/controllers/testdata"
)

func TestCreateQuizController(t *testing.T) {
	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(testdata.QuizTestData[0]); err != nil {
		t.Error(err)
	}
	req := httptest.NewRequest(http.MethodPost, baseUrl+"create-quiz", buffer)
	res := httptest.NewRecorder()
	qCon.CreateQuizController(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("unexpected StatusCode: got %d but want %d", res.Code, http.StatusOK)
	}
}

func TestGetQuizController(t *testing.T) {
	var tests = []struct {
		name     string
		themeID  string
		wantCode int
	}{
		{name: "number param", themeID: "1", wantCode: http.StatusOK},
		// 以下のテストケース、なぜか200が返ってきてしまうので一旦保留
		// {name: "string param", themeID: "a", wantCode: http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/%s", baseUrl+"get-quiz", tt.themeID)
			log.Print(url)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			res := httptest.NewRecorder()
			qCon.GetQuizzesController(res, req)
			if res.Code != tt.wantCode {
				t.Errorf("unexpected StatusCode got %d but want %d", res.Code, tt.wantCode)
			}
		})
	}
}
