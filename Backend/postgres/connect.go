package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConnection(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, "postgres://postgres:8787898@localhost:5432/postgres")
}
