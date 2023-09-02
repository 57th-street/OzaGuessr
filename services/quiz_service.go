package services

import (
	"github.com/57th-street/oza-gueser/apperrors"
	"github.com/57th-street/oza-gueser/models"
	"github.com/57th-street/oza-gueser/repositories"
)

func (s *Service) CreateQuizService(quiz models.Quiz) error {
	_, err := repositories.CreateQuiz(s.db, quiz)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return err
	}
	return nil
}

func (s *Service) GetQuizzesService(categoryID int) ([]models.Quiz, error) {
	quizzes, err := repositories.GetQuizzes(s.db, categoryID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return []models.Quiz{}, err
	}
	return quizzes, nil
}
