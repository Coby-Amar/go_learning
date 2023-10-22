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

func CreateUserSession(cwrar *ConfigWithRequestAndResponse, userId pgtype.UUID) bool {
	session, err := cwrar.Config.STORE.Get(cwrar.R, SESSION)
	if err != nil {
		slog.Error("Error getting session", ERROR, err)
		return false
	}
	secure, err := strconv.ParseBool(os.Getenv(PRODUCTION))
	if err != nil {
		secure = true
	}
	session.Options = &sessions.Options{
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		Secure:   secure,
		MaxAge:   time.Now().Add(time.Second * 10).Second(),
	}
	sessionParams := SessionParameters{
		UserID:            userId,
		Active:            true,
		CountResetTime:    time.Now().Add(time.Minute),
		RequestsPerMinute: 100,
		RequestsCount:     0,
	}
	session.Values[SESSION_PARAMETERS] = &sessionParams
	if err := session.Save(cwrar.R, cwrar.W); err != nil {
		slog.Error("Saved session fail", ERROR, err)
		return false
	}
	slog.Info("Created and saved session", SESSION_PARAMETERS, sessionParams)
	return true
}

func GetSessionParams(cwrar *ConfigWithRequestAndResponse) *SessionParameters {
	session, err := cwrar.Config.STORE.Get(cwrar.R, SESSION)
	if err != nil {
		slog.Error("getSessionParams failed", ERROR, err)
		return nil
	}
	val := session.Values[SESSION_PARAMETERS]
	params, ok := val.(*SessionParameters)
	if !ok {
		slog.Error("getSessionParams failed to type-assert params", ERROR, params)
		return nil
	}
	return params
}
