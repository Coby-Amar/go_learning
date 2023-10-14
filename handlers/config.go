package handlers

import (
	"database/sql"
	"log"
	"os"

	"github.com/coby-amar/go_learning/database"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func CreateApiConfig() *ApiConfig {
	connection, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database -", err)
		return nil
	}
	apiConfig := ApiConfig{
		DB: database.New(connection),
	}
	return &apiConfig

}
