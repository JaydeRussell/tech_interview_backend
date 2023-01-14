package data

import "github.com/JaydeRussell/tech_interview_backend/enum"

type Filter struct {
	SearchTerm string
	Page       int
	PageSize   int
	SortType   enum.SortType
}
