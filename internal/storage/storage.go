package storage

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnect(connStr string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, errors.Wrap(err, "GetConnect pgx.Connect")
	}
	return pool, err
}
