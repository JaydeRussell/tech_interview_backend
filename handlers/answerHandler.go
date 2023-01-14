package handlers

import (
	"github.com/JaydeRussell/tech_interview_backend/interfaces"
	"github.com/gin-gonic/gin"
)

type AnswerHandler struct {
	dataService interfaces.DataServicer
}

func NewAnswerHandler(dataService interfaces.DataServicer) *AnswerHandler {
	return &AnswerHandler{
		dataService: dataService,
	}
}

func (a *AnswerHandler) HandleGet(c *gin.Context) {
	panic("NOT YET IMPLEMENTED")
}

func (a *AnswerHandler) HandleCreate(c *gin.Context) {
	panic("NOT YET IMPLEMENTED")
}

func (a *AnswerHandler) HandleUpdate(c *gin.Context) {
	panic("NOT YET IMPLEMENTED")
}

func (a *AnswerHandler) HandleDelete(c *gin.Context) {
	panic("NOT YET IMPLEMENTED")
}
