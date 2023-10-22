package handlers

import (
	"context"
	"encoding/gob"
	"log/slog"
	"os"

	"github.com/coby-amar/go_learning/database"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
)

func CreateApiConfig() *ApiConfig {

	connection, err := pgx.Connect(context.Background(), os.Getenv(DATABASE_URL))
	if err != nil {
		slog.Error("Failed to connect to database", ERROR, err)
		return nil
	}

	gob.Register(&sessionParameters{})
	sAK := os.Getenv(SESSION_AUTHORIZATION_KEY)
	sEK := os.Getenv(SESSION_ENCRYPTION_KEY)
	jwtSK := os.Getenv(JWT_SECRET_KEY)

	return &ApiConfig{
		DB:             database.New(connection),
		STORE:          sessions.NewCookieStore([]byte(sAK), []byte(sEK)),
		JWT_SECRET_KEY: jwtSK,
	}

}
