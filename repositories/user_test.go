package repositories_test

import (
	"testing"

	"github.com/57th-street/oza-gueser/models"
	"github.com/57th-street/oza-gueser/repositories"
	"github.com/57th-street/oza-gueser/repositories/testdata"
)

func TestGetProfile(t *testing.T) {
	expected := testdata.UserTestData
	got, err := repositories.GetProfile(testDB, expected.ID)
	if err != nil {
		t.Error(err)
	}
	if got.Name != expected.Name {
		t.Errorf("got %s but want %s", got.Name, expected.Name)
	}
	if got.ImageUrl != expected.ImageUrl {
		t.Errorf("got %s but want %s", got.ImageUrl, expected.ImageUrl)
	}
	if got.Intro != expected.Intro {
		t.Errorf("got %s but want %s", got.Intro, expected.Intro)
	}
}

func TestUpdateProfile(t *testing.T) {
	expected := models.Profile{
		Name:     "testUser2",
		ImageUrl: "testImage2.png",
		Intro:    "テスト用自己紹介文2",
	}
	if err := repositories.UpdateProfile(testDB, testdata.UserTestData.ID, expected); err != nil {
		t.Error(err)
	}
	const verifySqlStr = "select name, image_url, intro from users where id = ?"
	var got models.Profile
	err := testDB.QueryRow(verifySqlStr, testdata.UserTestData.ID).Scan(&got.Name, &got.ImageUrl, &got.Intro)
	if err != nil {
		t.Error(err)
	}
	if got.Name != expected.Name {
		t.Errorf("got %s but want %s", got.Name, expected.Name)
	}
	if got.ImageUrl != expected.ImageUrl {
		t.Errorf("got %s but want %s", got.ImageUrl, expected.ImageUrl)
	}
	if got.Intro != expected.Intro {
		t.Errorf("got %s but want %s", got.Intro, expected.Intro)
	}
}
