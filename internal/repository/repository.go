package repository

import (
	"context"

	"github.com/salesforceanton/meower/internal/schema"
)

type Repository interface {
	Close()
	InsertMeow(context.Context, schema.Meow) error
	GetMeowsList(context.Context) ([]schema.Meow, error)
}
