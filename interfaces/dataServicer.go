package interfaces

import "github.com/JaydeRussell/tech_interview_backend/data"

type DataServicer interface {
	Search(filter data.Filter) (result []data.Question, totalFound int, err error)
}
