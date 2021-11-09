package repository

import (
	"fmt"
	"github.com/jackc/pgx/v4"
)

// POSTGRES REPOSITORY...
type PostgresRepository struct {
	DB *pgx.Conn
}

// NEW POSTGRES REPOSITORY...
func NewPostgresRepository(conn *pgx.Conn) *PostgresRepository {
	fmt.Println("Connect to DATABASE postgres")
	return &PostgresRepository{
		DB: conn,
	}
}