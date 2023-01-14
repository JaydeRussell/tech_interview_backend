package handlers

import "github.com/JaydeRussell/tech_interview_backend/interfaces"

type QuestionHandler struct {
	dataService interfaces.DataServicer
}

func NewQuestionHandler(dataService interfaces.DataServicer) *QuestionHandler {
	return &QuestionHandler{
		dataService: dataService,
	}
}

func (q *QuestionHandler) HandleCreate() {
	panic("NOT YET IMPLEMENTED")
}

func (q *QuestionHandler) HandleGet() {
	panic("NOT YET IMPLEMENTED")
}

func (q *QuestionHandler) HandleUpdate() {
	panic("NOT YET IMPLEMENTED")
}

func (q *QuestionHandler) HandleDelete() {
	panic("NOT YET IMPLEMENTED")
}

func (q *QuestionHandler) HandleSearch() {
	panic("NOT YET IMPLEMENTED")
}
