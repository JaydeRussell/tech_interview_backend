package services

import (
	"database/sql"

	"github.com/JaydeRussell/tech_interview_backend/data"
)

type PostgresService struct {
	db *sql.DB
}

func NewPostgresService(db *sql.DB) *PostgresService {
	return &PostgresService{
		db: db,
	}
}

func (s *PostgresService) GetQuestion(id string) data.Question {
	panic("NOT YET IMPLEMENTED")
}
func (s *PostgresService) GetQuestions(ids []string) []data.Question {
	panic("NOT YET IMPLEMENTED")
}
func (s *PostgresService) Search(filter data.Filter) (result []data.Question, totalFound int, err error) {
	return
}
func (s *PostgresService) GetAnswersByQuestionId(id string) []data.Question {
	panic("NOT YET IMPLEMENTED")
}
