package services

import (
	"database/sql"

	"github.com/JaydeRussell/tech_interview_backend/data"
)

type MySQLService struct {
	db *sql.DB
}

func NewMySQLService(db *sql.DB) *MySQLService {
	return &MySQLService{
		db: db,
	}
}

func (s *MySQLService) GetQuestion(id string) data.Question {
	panic("NOT YET IMPLEMENTED")
}
func (s *MySQLService) GetQuestions(ids []string) []data.Question {
	panic("NOT YET IMPLEMENTED")
}
func (s *MySQLService) Search(filter data.Filter) (result []data.Question, totalFound int, err error) {
	return
}
func (s *MySQLService) GetAnswersByQuestionId(id string) []data.Question {
	panic("NOT YET IMPLEMENTED")
}
