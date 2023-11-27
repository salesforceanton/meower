package schema

import "time"

type Meow struct {
	Id        string    `json:"Dd"         db:"id"`
	Body      string    `json:"Body"       db:"body"`
	CreatedAt time.Time `json:"CreatedAt" db:"created_at"`
}
