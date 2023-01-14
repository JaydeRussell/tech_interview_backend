package interfaces

import "github.com/JaydeRussell/tech_interview_backend/data"

type DataServicer interface {
	GetQuestion(id string) data.Question
	GetQuestions(ids []string) []data.Question
	Search(filter data.Filter) (result []data.Question, totalFound int, err error)
	GetAnswersByQuestionId(id string) []data.Question
}
