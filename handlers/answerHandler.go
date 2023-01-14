package handlers

import "github.com/JaydeRussell/tech_interview_backend/interfaces"

type AnswerHandler struct {
	dataService interfaces.DataServicer
}

func NewAnswerHandler(dataService interfaces.DataServicer) *AnswerHandler {
	return &AnswerHandler{
		dataService: dataService,
	}
}

func (a *AnswerHandler) HandleGet() {
	panic("NOT YET IMPLEMENTED")
}

func (a *AnswerHandler) HandleCreate() {
	panic("NOT YET IMPLEMENTED")
}

func (a *AnswerHandler) HandleUpdate() {
	panic("NOT YET IMPLEMENTED")
}

func (a *AnswerHandler) HandleDelete() {
	panic("NOT YET IMPLEMENTED")
}
