package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/57th-street/oza-gueser/controllers/testdata"
	"github.com/57th-street/oza-gueser/models"
	"github.com/57th-street/oza-gueser/utils"
)

func TestGetProfileController(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, baseUrl+"profile", nil)
	req = req.WithContext(utils.SetUserID(req.Context(), 1))
	res := httptest.NewRecorder()
	uCon.GetProfileController(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("unexpected StatusCode: got %d but want %d", res.Code, http.StatusOK)
	}
	var gotProfile models.Profile
	if err := json.NewDecoder(res.Body).Decode(&gotProfile); err != nil {
		t.Error(err)
	}
	expectedProfile := testdata.ProfileTestData
	if gotProfile.Name != expectedProfile.Name {
		t.Errorf("unexpected profile name: got %s but want %s", gotProfile.Name, expectedProfile.Name)
	}
	if gotProfile.ImageUrl != expectedProfile.ImageUrl {
		t.Errorf("unexpected profile image_url: got %s but want %s", gotProfile.ImageUrl, expectedProfile.ImageUrl)
	}
	if gotProfile.Intro != expectedProfile.Intro {
		t.Errorf("unexpected profile intro: got %s but want %s", gotProfile.Intro, expectedProfile.Intro)
	}
}

func TestUpdateProfileController(t *testing.T) {
	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(testdata.ProfileTestData); err != nil {
		t.Error(err)
	}
	req := httptest.NewRequest(http.MethodPost, baseUrl+"update-profile", buffer)
	req = req.WithContext(utils.SetUserID(req.Context(), 1))
	res := httptest.NewRecorder()
	uCon.UpdateProfileController(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("unexpected StatusCode: got %d but want %d", res.Code, http.StatusOK)
	}
}
