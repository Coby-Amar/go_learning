package config

import (
	"context"
	"encoding/gob"
	"log/slog"
	"os"

	"github.com/coby-amar/go_learning/database"
	"github.com/coby-amar/go_learning/main/utils"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
)

func Setup() *utils.ApiConfig {
	connection, err := pgx.Connect(context.Background(), os.Getenv(DATABASE_URL))
	if err != nil {
		slog.Error("Failed to connect to database", utils.ERROR, err)
		return nil
	}

	gob.Register(&utils.SessionParameters{})
	sAK := os.Getenv(SESSION_AUTHORIZATION_KEY)
	sEK := os.Getenv(SESSION_ENCRYPTION_KEY)
	jwtSK := os.Getenv(utils.JWT_SECRET_KEY)

	return &utils.ApiConfig{
		Connection:     connection,
		Queries:        database.New(connection),
		Store:          sessions.NewCookieStore([]byte(sAK), []byte(sEK)),
		JWT_SECRET_KEY: jwtSK,
	}

}
