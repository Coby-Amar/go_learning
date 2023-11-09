package utils

import (
	"net/http"
	"time"

	"github.com/coby-amar/go_learning/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type jwtClaims struct {
	jwt.RegisteredClaims
	UserID pgtype.UUID
}

type ApiConfig struct {
	Connection     *pgx.Conn
	Queries        *database.Queries
	Store          *sessions.CookieStore
	JWT_SECRET_KEY string
}

type ConfigWithRequestAndResponse struct {
	Config  *ApiConfig
	W       http.ResponseWriter
	R       *http.Request
	Sparams *SessionParameters
}

type SessionParameters struct {
	UserID            pgtype.UUID
	Active            bool
	CountResetTime    time.Time
	RequestsPerMinute int
	RequestsCount     int
}

type LoginJson struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=70"`
}

type RegistrationJson struct {
	database.DailyLimit
	Username    string `json:"username" validate:"required,email"`
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phonenumber" validate:"required"`
	Password    string `json:"password" validate:"required,min=8,max=70"`
}

type RegistrationJsonResponse struct {
	database.DailyLimit
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phonenumber" validate:"required"`
}

type UserCreateReportWithEntries struct {
	Report  database.CreateReportParams          `json:"report" validate:"required"`
	Entries []database.CreateReportEntriesParams `json:"entriesToCreate" validate:"required,min=1,max=20"`
}

type UserUpdateReportWithEntries struct {
	Report          database.UpdateReportParams          `json:"report" validate:"required"`
	ExistingEntries []database.UpdateReportEntryParams   `json:"existingEntries" validate:"required,min=1,max=15"`
	EntriesToCreate []database.CreateReportEntriesParams `json:"entriesToCreate" validate:"required,min=1,max=15"`
}
