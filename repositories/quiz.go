package repositories

import (
	"database/sql"

	"github.com/57th-street/oza-gueser/models"
)

func CreateQuiz(db *sql.DB, quiz models.Quiz) (sql.Result, error) {
	const testSqlStr = `SELECT id FROM themes;`

	const sqlStr = `insert into quizzes (quest, answer, descript, theme_id) values (?, ?, ?, ?); `
	return db.Exec(sqlStr, quiz.Quest, quiz.Answer, quiz.Descript, quiz.ThemeID)
}

func GetQuizzes(db *sql.DB, themeID int) ([]models.Quiz, error) {
	const sqlStr = `select id, quest, answer, descript from quizzes where theme_id = ?;`
	rows, err := db.Query(sqlStr, themeID)
	if err != nil {
		return []models.Quiz{}, err
	}
	defer rows.Close()
	quizzes := []models.Quiz{}
	for rows.Next() {
		var quiz models.Quiz
		if err := rows.Scan(&quiz.ID, &quiz.Quest, &quiz.Answer, &quiz.Descript); err != nil {
			return []models.Quiz{}, err
		}
		quizzes = append(quizzes, quiz)
	}
	if err := rows.Err(); err != nil {
		return []models.Quiz{}, err
	}
	return quizzes, err
}
