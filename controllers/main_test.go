package controllers_test

import (
	"testing"

	"github.com/57th-street/oza-gueser/controllers"
	"github.com/57th-street/oza-gueser/controllers/testdata"
)

var aCon *controllers.AuthController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewAuthController(ser)
	m.Run()
}
