package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/salesforceanton/meower/internal/schema"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) Close() {
	r.db.Close()
}

func (r *PostgresRepo) InsertMeow(ctx context.Context, message schema.Meow) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (id, body, created_at) VALUES ($1, $2, $3) RETURNING id
	`, MEOWS_TABLE_NAME)

	_, err := r.db.ExecContext(ctx, query, message.Id, message.Body, message.CreatedAt.Format(time.RFC3339))
	return err
}

func (r *PostgresRepo) GetMeowsList(ctx context.Context, skip, take int64) ([]schema.Meow, error) {
	var result []schema.Meow

	query := fmt.Sprintf(`
		SELECT id, body, created_at
		FROM %s 
		ORDER BY id DESC
		OFFSET $1 LIMIT $2 
	`, MEOWS_TABLE_NAME)

	if err := r.db.SelectContext(ctx, &result, query, skip, take); err != nil {
		return nil, err
	}
	return result, nil
}
