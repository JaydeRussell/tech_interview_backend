package data

import "github.com/JaydeRussell/tech_interview_backend/enum"

type SearchResponse struct {
	Results    []Question    `json:"results"`
	TotalFound int           `json:"totalFound"`
	SearchTerm string        `json:"search-term"`
	Page       int           `json:"page"`
	PageSize   int           `json:"pageSize"`
	SortType   enum.SortType `json:"sortType"`
}
