package testdata

import "github.com/57th-street/oza-gueser/models"

var UserTestData = models.User{
	ID:       1,
	Email:    "test@test.com",
	Password: "hashedPassword",
	Profile: models.Profile{
		Name:     "testUser",
		ImageUrl: "testImage.png",
		Intro:    "テスト用自己紹介文です。",
	},
}
