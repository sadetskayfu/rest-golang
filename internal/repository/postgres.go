package repository

import (
	"fmt"
	"github.com/jackc/pgx/v4"
)

// PostgresRepository ...

type PostgresRepository struct {
	DB *pgx.Conn
}

// NewPostgresRepository ...

func NewPostgresRepository(conn *pgx.Conn) *PostgresRepository {
	fmt.Println("Connect to DATABASE postgres")
	return &PostgresRepository{
		DB: conn,
	}
}