package config

import (
	"fmt"
)

var (
	DBUser = "postgres"
	DBPassword = "1234"
	DBName     = "postgres"
	DBHost     = "localhost"
	DBPort     = "5432"
	DBType     = "postgres"
)

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable ",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword,
	)
	return dataBase
}
