package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/57th-street/oza-gueser/apperrors"
	"github.com/57th-street/oza-gueser/controllers/services"
	"github.com/57th-street/oza-gueser/models"
	"github.com/57th-street/oza-gueser/utils"
)

type QuizController struct {
	service services.QuizServicer
}

func NewQuizController(s services.QuizServicer) *QuizController {
	return &QuizController{service: s}
}

func (q *QuizController) CreateQuizController(w http.ResponseWriter, req *http.Request) {
	var input models.Quiz
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	if err := q.service.CreateQuizService(input); err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (q *QuizController) GetQuizzesController(w http.ResponseWriter, req *http.Request) {
	param, err := utils.GetPathParam(req.URL.Path)
	if err != nil {
		err = apperrors.InvalidPathParam.Wrap(err, "invalid path parameter")
		return
	}
	quizzes, err := q.service.GetQuizzesService(param)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(quizzes)
}
