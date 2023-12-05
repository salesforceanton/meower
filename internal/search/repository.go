package search_repo

import (
	"context"

	"github.com/salesforceanton/meower/internal/schema"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, message schema.Meow) error
	SearchMeows(ctx context.Context, query string, skip, take int64) ([]schema.Meow, error)
}
