package config

import (
	"fmt"
)

// Variable for connect to postgres DB

const (
	dbUser     = "postgres"
	dbPassword = "1234"
	dbName     = "postgres"
	dbHost     = "localhost"
	dbPort     = "5432"
)

// Return string for connect to postgres DB

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable ",
		dbHost,
		dbPort,
		dbUser,
		dbName,
		dbPassword,
	)
	return dataBase
}
