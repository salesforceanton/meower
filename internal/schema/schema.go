package schema

import "time"

type Meow struct {
	Id        string    `json:"id"         db:"id"`
	Body      string    `json:"body"       db:"body"`
	CreatedAt time.Time `json:"createdAt"  db:"created_at"`
}
