package schema

import "time"

type Meow struct {
	Id         string    `json:"id"         db:"id"`
	Body       string    `json:"body"       db:"body"`
	Created_at time.Time `json:"created_at" db:"created_at"`
}
