package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/salesforceanton/meower/internal/config"
)

const (
	MEOWS_TABLE_NAME  = "meows"
	POSTGRESS_DB_TYPE = "postgres"
)

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresName))
	db, err := sqlx.Open(POSTGRESS_DB_TYPE, pgUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
