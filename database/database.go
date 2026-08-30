package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) *pgx.Conn {
	conn_string := os.Getenv("CONN_STRING_DOCKER")
	conn, err := pgx.Connect(ctx, conn_string)
	if err != nil {
		panic(err)
	}
	return conn
}
