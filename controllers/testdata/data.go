package testdata

import (
	"github.com/57th-street/oza-gueser/controllers"
	"github.com/57th-street/oza-gueser/models"
)

var UserInputTestData = controllers.UserInput{
	Email:    testdata.UserTestData.Email,
	Password: testdata.UserTestData.Password,
}

var ProfileTestData = models.Profile{
	Name:     "testUser",
	ImageUrl: "testImage.png",
	Intro:    "テスト用自己紹介文です。",
}
