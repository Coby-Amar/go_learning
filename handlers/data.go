package handlers

import (
	"net/http"
	"time"

	"github.com/coby-amar/go_learning/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

type ApiConfig struct {
	DB             *database.Queries
	STORE          *sessions.CookieStore
	JWT_SECRET_KEY string
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
}

type sessionParameters struct {
	UserID            uuid.UUID
	Active            bool
	CountResetTime    time.Time
	RequestsPerMinute int
	RequestsCount     int
}

type registrationJson struct {
	Username    string `json:"username" validate:"required,email"`
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phonenumber" validate:"required"`
	Password    string `json:"password" validate:"required,min=8,max=70"`
}

type loginJson struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=70"`
}

type userReport struct {
	Date time.Time `json:"date"`
}

type userGetReport struct {
	userReport
	Entries []database.GetReportEntriesRow `json:"entries"`
}

type UserRequestReport struct {
	userReport
	Entries []database.CreateReportEntryParams `json:"entries"`
}

type userCreatedReport struct {
	userReport
	Entries []database.ReportEntry `json:"entries"`
}
