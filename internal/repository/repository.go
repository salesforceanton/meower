package repository

import (
	"context"

	"github.com/salesforceanton/meower/internal/schema"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, message schema.Meow) error
	GetMeowsList(ctx context.Context, skip, take int64) ([]schema.Meow, error)
}
