package data

type Answer struct {
	ID         string `db:"id"`
	QuestionID string `db:"question_id"`
	Body       string `db:"body"`
}
