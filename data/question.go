package data

import "time"

type Question struct {
	ID        string    `json:"id" db:"id"`
	Body      string    `json:"body" db:"body"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	Answers   []Answer  `json:"answers"`
}
