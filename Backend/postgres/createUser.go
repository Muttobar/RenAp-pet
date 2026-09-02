package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateTableUser(connect *pgx.Conn, ctx context.Context) (pgconn.CommandTag, error) {
	sqlQwery := ` CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(60) NOT NULL UNIQUE,
	name VARCHAR(20), 
	familyName VARCHAR(30),
	patronymic VARCHAR(30),
	created_at DATE ,
	setting JSON
	);`
	return connect.Exec(ctx, sqlQwery)
}
