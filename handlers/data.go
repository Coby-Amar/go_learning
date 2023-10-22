package handlers

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

type UserCreateReportWithEntries struct {
	Report  database.CreateReportParams          `json:"report" validate:"required"`
	Entries []database.CreateReportEntriesParams `json:"entries" validate:"required,min=1,max=20"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type jsonHandler[T interface{}] func(http.ResponseWriter, *http.Request, T)
type jsonParams[T interface{}] struct {
	params T
}

type jwtClaims struct {
	jwt.RegisteredClaims
	UserID pgtype.UUID
}

type sessionParameters struct {
	UserID            pgtype.UUID
	Active            bool
	CountResetTime    time.Time
	RequestsPerMinute int
	RequestsCount     int
}

type registrationJson struct {
	Username    string `json:"username" validate:"required,email,max=50"`
	Name        string `json:"name" validate:"required,min=4,max=50"`
	PhoneNumber string `json:"phonenumber" validate:"required,min=8,max=10"`
	Password    string `json:"password" validate:"required,min=8,max=30,containsany=!@#?"`
}

type loginJson struct {
	Username string `json:"username" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}
