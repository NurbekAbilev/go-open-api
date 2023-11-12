package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDatabase() (*sql.DB, error) {
	connstr := os.Getenv("DB_URL")
	if connstr == "" {
		return nil, errors.New("empty DB_URL variable, set inside .env")
	}

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitPgxConnect(ctx context.Context, config *pgxpool.Config) (c *pgxpool.Pool, err error) {
	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("pgx: could not connect to postgres database : %w", err)
	}

	return conn, nil
}
