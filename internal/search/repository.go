package search

import (
	"context"

	"github.com/salesforceanton/meower/internal/schema"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, message schema.Meow) error
	searchMeows(ctx context.Context, query string, skip, take int64) ([]schema.Meow, error)
}
