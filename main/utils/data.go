package utils

import (
	"net/http"
	"time"

	"github.com/coby-amar/go_learning/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgtype"
)

type ApiConfig struct {
	DB             *database.Queries
	STORE          *sessions.CookieStore
	JWT_SECRET_KEY string
}

type ConfigWithRequestAndResponse struct {
	Config *ApiConfig
	W      http.ResponseWriter
	R      *http.Request
}

type SessionParameters struct {
	UserID            pgtype.UUID
	Active            bool
	CountResetTime    time.Time
	RequestsPerMinute int
	RequestsCount     int
}

type jwtClaims struct {
	jwt.RegisteredClaims
	UserID pgtype.UUID
}
