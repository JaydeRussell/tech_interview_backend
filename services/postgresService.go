package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JaydeRussell/tech_interview_backend/data"
	"github.com/JaydeRussell/tech_interview_backend/enum"
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

func (s *PostgresService) Search(filter data.Filter) (results []data.Question, totalFound int, err error) {
	// global always do things. Mostly just ignore deleted questions
	tx := s.db.Table("question").
		Where("deleted = false")

	// if we were provided a search term, apply it here
	// NOTE: this is case insensitive
	if filter.SearchTerm != "" {
		tx.Where("LOWER(body) like LOWER(?)", fmt.Sprintf("%%%s%%", filter.SearchTerm))
	}

	// apply the sortType to the query
	switch filter.SortType {
	case enum.CreatedAtAsc:
		tx.Order("created_at ASC")
	case enum.CreatedAtDesc:
		tx.Order("created_at DESC")
	}

	// offset calculates the results that would exist given a page number
	tx.Offset((filter.Page - 1) * filter.PageSize)
	// Limit will always give us a given page size
	tx.Limit(filter.PageSize)
	// this will return our results
	tx.Find(&results)
	err = tx.Error
	if err != nil {
		log.Printf("Error attempting Query with given filter: %+v, err: %s", filter, err.Error())
		return
	}
	// we're going to repeat the query without any pagination to calculate the total
	totalFound, err = s.calculateQueryTotal(filter)
	if err != nil {
		log.Printf("Error attempting to calculate query total: %s", err.Error())
		return
	}

	// this could probably be done with gorm preloading
	// but because we set up the tables with migrations instead of gorm it's a lot more complicated
	results, err = s.fetchAndAttachAnswers(results)
	if err != nil {
		log.Printf("Error attempting to attach answers to the questions: %s", err.Error())
		return
	}
	return
}

func (s *PostgresService) calculateQueryTotal(filter data.Filter) (total int, err error) {
	tx := s.db.Table("question").
		Where("deleted = false")

	// if we were provided a search term, apply it here
	// NOTE: this is case insensitive
	if filter.SearchTerm != "" {
		tx.Where("LOWER(body) like LOWER(?)", fmt.Sprintf("%%%s%%", filter.SearchTerm))
	}

	var count int64
	err = tx.Count(&count).Error
	if err != nil {
		return
	}
	total = int(count)
	return
}

func (s *PostgresService) fetchAndAttachAnswers(questions []data.Question) ([]data.Question, error) {
	ids := s.getIDsFromQuestions(questions)
	answers, err := s.getAnswersByQuestionIds(ids)
	if err != nil {
		return questions, err
	}

	mappedAnswers := make(map[string][]data.Answer)
	for _, answer := range answers {
		mappedAnswers[answer.QuestionID] = append(mappedAnswers[answer.QuestionID], answer)
	}

	for index, question := range questions {
		questions[index].Answers = append(questions[index].Answers, mappedAnswers[question.ID]...)
	}

	return questions, err
}

func (s *PostgresService) getIDsFromQuestions(questions []data.Question) (ids []string) {
	for _, question := range questions {
		ids = append(ids, question.ID)
	}
	return
}

func (s *PostgresService) getAnswersByQuestionIds(id []string) (answers []data.Answer, err error) {
	err = s.db.Table("answer").
		Where("deleted = false").
		Where("question_id in ?", id).
		Find(&answers).Error
	return
}
