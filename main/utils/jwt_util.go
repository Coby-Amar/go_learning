package utils

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateJWTCookie(user_id pgtype.UUID, jwtUserSecretKey string) *http.Cookie {
	expiresAt := time.Now().Add(time.Minute * 30)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		UserID: user_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	})
	token, err := jwtToken.SignedString([]byte(jwtUserSecretKey))
	if err != nil {
		slog.Error("Failed to SignedString cookie", ERROR, err)
		return nil
	}
	secure, err := strconv.ParseBool(os.Getenv(PRODUCTION))
	if err != nil {
		secure = true
	}
	return &http.Cookie{
		Name:     JWT_COOKIE,
		Path:     "/",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   expiresAt.Second(),
		Secure:   secure,
	}
}

func ValidateJWT(jwtCookie *http.Cookie, jwtUserSecretKey string) (pgtype.UUID, error) {
	claims := jwtClaims{}
	parsedToken, err := jwt.ParseWithClaims(jwtCookie.Value, &claims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			slog.Error("Failed to unsign jwt")
			return "", UnauthorizedError
		}
		if expTime, err := t.Claims.GetExpirationTime(); err != nil || expTime.Before(time.Now()) {
			slog.Error("Token expired or GetExpirationTime error", ERROR, err)
			return "", UnauthorizedError
		}
		return []byte(jwtUserSecretKey), nil
	})
	if err != nil {
		slog.Error("Failed to Parse jwt", ERROR, err)
		return pgtype.UUID{}, UnauthorizedError
	}
	if !parsedToken.Valid {
		slog.Error("Invalid token")
		return pgtype.UUID{}, UnauthorizedError
	}
	return claims.UserID, nil
}
