package utils

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateUserSession(store *sessions.CookieStore, w http.ResponseWriter, r *http.Request, userId pgtype.UUID) bool {
	slog.Info("CreateUserSession")
	session, err := store.Get(r, SESSION)
	if err != nil {
		slog.Error("Store.Get", ERROR, err)
		return false
	}
	secure, err := strconv.ParseBool(os.Getenv(PRODUCTION))
	if err != nil {
		slog.Error("ParseBool", ERROR, err)
		secure = true
	}
	session.Options = &sessions.Options{
		HttpOnly: true,
		Domain:   "localhost, localhost:2318",
		Path:     "/",
		// SameSite: http.SameSiteStrictMode,
		Secure: secure,
		MaxAge: int(time.Now().Add(time.Minute * 20).Unix()),
	}
	sessionParams := SessionParameters{
		UserID:            userId,
		Active:            true,
		CountResetTime:    time.Now().Add(time.Minute),
		RequestsPerMinute: 100,
		RequestsCount:     0,
	}
	session.Values[SESSION_PARAMETERS] = &sessionParams
	if err := session.Save(r, w); err != nil {
		slog.Error("session.Save", ERROR, err)
		return false
	}
	return true
}

func GetSessionParams(store *sessions.CookieStore, r *http.Request) *SessionParameters {
	slog.Info("GetSessionParams")
	session, err := store.Get(r, SESSION)
	if err != nil {
		slog.Error("Store.Get", ERROR, err)
		return nil
	}
	val := session.Values[SESSION_PARAMETERS]
	params, ok := val.(*SessionParameters)
	if !ok {
		slog.Error("type-assert", ERROR, params)
		return nil
	}
	return params
}

func DeleteSessionParams(store *sessions.CookieStore, w http.ResponseWriter, r *http.Request) bool {
	slog.Info("GetSessionParams")
	session, err := store.Get(r, SESSION)
	if err != nil {
		slog.Error("Store.Get", ERROR, err)
		return false
	}
	session.Options = &sessions.Options{
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	}
	session.Values[SESSION_PARAMETERS] = nil
	if err := session.Save(r, w); err != nil {
		slog.Error("session.Save", ERROR, err)
		return false
	}
	return true
}
