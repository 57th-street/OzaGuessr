package testdata

import (
	"github.com/57th-street/oza-gueser/controllers"
	"github.com/57th-street/oza-gueser/models"
)

var userTestData = models.User{
	ID:       1,
	Email:    "test@test.com",
	Password: "hashedPassword",
	Profile: models.Profile{
		Name:     "testUser",
		ImageUrl: "testImage.png",
		Intro:    "テスト用自己紹介文です。",
	},
}

var UserInputTestData = controllers.UserInput{
	Email:    userTestData.Email,
	Password: userTestData.Password,
}

var ProfileTestData = models.Profile{
	Name:     userTestData.Profile.Name,
	ImageUrl: userTestData.Profile.ImageUrl,
	Intro:    userTestData.Profile.Intro,
}

var QuizTestData = []models.Quiz{
	models.Quiz{
		ID:       1,
		Quest:    "ozaki_tokyo_dome.png",
		Answer:   "1988",
		Descript: "東京ドーム公演の様子である。キャリア唯一であると同時に自身最多の5万6000人を動員した。",
		ThemeID:  1,
	},
	models.Quiz{
		ID:       2,
		Quest:    "ozaki_nihon_senenkan.png",
		Answer:   "1985",
		Descript: "日本青年館でのライブ模様である。尾崎にとって東京での初めてのホールコンサートとなった。",
		ThemeID:  1,
	},
	models.Quiz{
		ID:       3,
		Quest:    "honnoj_incident.png",
		Answer:   "1582",
		Descript: "本能寺の変の様子である。早朝、明智光秀が謀反を起こし、京都本能寺に滞在する主君・織田信長を襲撃した事件である。",
		ThemeID:  2,
	},
}
