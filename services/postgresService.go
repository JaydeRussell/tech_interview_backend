package services

import (
	"database/sql"
	"log"

	"github.com/JaydeRussell/tech_interview_backend/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresService struct {
	db *gorm.DB
}

func NewPostgresService(db *sql.DB) *PostgresService {
	gDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to create gorm instance")
	}
	return &PostgresService{
		db: gDB,
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
