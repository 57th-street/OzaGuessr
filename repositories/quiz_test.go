package repositories_test

import (
	"testing"

	"github.com/57th-street/oza-gueser/models"
	"github.com/57th-street/oza-gueser/repositories"
)

func TestCreateQuiz(t *testing.T) {
	testQuiz := models.Quiz{
		ID:       4,
		Quest:    "ozaki_ruid.png",
		Answer:   "1984",
		Descript: "尾崎豊のデビューライブの様子である。通っていた高校の卒業式の日に敢行した。",
		ThemeID:  1,
	}
	result, err := repositories.CreateQuiz(testDB, testQuiz)
	if err != nil {
		t.Fatalf("Failed to create quiz: %v", err)
	}
	gotID, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("Failed to LastInsertId(): %v", err)
	}
	wantID := int64(testQuiz.ID)
	if gotID != wantID {
		t.Fatalf("got %d but want %d", gotID, wantID)
	}
	t.Cleanup(func() {
		const sqlStr = "delete from quizzes where id = ?"
		testDB.Exec(sqlStr, testQuiz.ID)
	})
}

func TestGetQuizzes(t *testing.T) {
	wantNum := 2
	got, err := repositories.GetQuizzes(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}
	if num := len(got); num != wantNum {
		t.Errorf("want %d but got %d articles\n", wantNum, num)
	}
}
