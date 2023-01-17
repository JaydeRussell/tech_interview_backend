package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/JaydeRussell/tech_interview_backend/constants"
	"github.com/JaydeRussell/tech_interview_backend/data"
	"github.com/JaydeRussell/tech_interview_backend/enum"
	"github.com/JaydeRussell/tech_interview_backend/interfaces"
	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
	dataService interfaces.DataServicer
}

func NewQuestionHandler(dataService interfaces.DataServicer) *QuestionHandler {
	return &QuestionHandler{
		dataService: dataService,
	}
}

func (q *QuestionHandler) HandleSearch(c *gin.Context) {
	filter, err := generateSearchFilterFromQuery(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, data.ErrorResponse{Message: "unable to parse parameters", Err: err.Error()})
		return
	}

	results, totalFound, err := q.dataService.Search(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, data.ErrorResponse{Message: "Unable to perform search", Err: err.Error()})
		return
	}

	result := data.SearchResponse{
		Results:    results,
		TotalFound: totalFound,
		SearchTerm: filter.SearchTerm,
		Page:       filter.Page,
		PageSize:   filter.PageSize,
		SortType:   filter.SortType,
	}

	c.JSON(http.StatusOK, result)
	return
}

func generateSearchFilterFromQuery(c *gin.Context) (filter data.Filter, err error) {
	filter.SearchTerm = c.Query("search-term")
	filter.Page = constants.DefaultPage
	filter.PageSize = constants.DefaultPageSize
	filter.SortType = constants.DefaultSortType

	if c.Query("page") != "" {
		filter.Page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			err = errors.New("invalid page parameter provided")
			return
		}
	}
	if c.Query("page-size") != "" {
		filter.PageSize, err = strconv.Atoi(c.Query("page-size"))
		if err != nil {
			err = errors.New("invalid page-size parameter provided")
			return
		}
	}

	if c.Query("sort-type") != "" {
		// this will return an empty string of the provided sort-type isn't one we support
		sortType := enum.StringToSortType(c.Query("sort-type"))
		if sortType != "" {
			filter.SortType = sortType
		} else {
			err = errors.New("invalid sort-type parameter provided")
			return
		}
	}
	return
}
