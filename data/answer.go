package data

type Answer struct {
	ID         string `json:"id" db:"id"`
	Body       string `json:"body" db:"body"`
	CreatedAt  string `json:"createdAt" db:"created_at"`
	QuestionID string `json:"questionId" db:"question_id"`
}
