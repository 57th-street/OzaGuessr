package controllers_test

import (
	"testing"

	"github.com/57th-street/oza-gueser/controllers"
	"github.com/57th-street/oza-gueser/controllers/testdata"
)

// 環境によってurlが変わるため今後改善の余地あり
var baseUrl = "http://localhost:8080/"
var aCon *controllers.AuthController
var uCon *controllers.UserController
var qCon *controllers.QuizController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewAuthController(ser)
	uCon = controllers.NewUserController(ser)
	qCon = controllers.NewQuizController(ser)
	m.Run()
}
