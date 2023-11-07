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
	slog.Info("CreateJWTCookie")
	expiresAt := time.Now().Add(5 * time.Hour * 24)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		UserID: user_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	})
	token, err := jwtToken.SignedString([]byte(jwtUserSecretKey))
	if err != nil {
		slog.Error("SignedString", ERROR, err)
		return nil
	}
	secure, err := strconv.ParseBool(os.Getenv(PRODUCTION))
	if err != nil {
		slog.Error("ParseBool", ERROR, err)
		secure = true
	}
	return &http.Cookie{
		Name:     JWT_COOKIE,
		Domain:   "localhost, localhost:2318",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
		// SameSite: http.SameSiteStrictMode,
		MaxAge: int(expiresAt.Unix()),
		Secure: secure,
	}
}

func ValidateJWT(jwtCookie *http.Cookie, jwtUserSecretKey string) (pgtype.UUID, error) {
	slog.Info("ValidateJWT")
	claims := jwtClaims{}
	parsedToken, err := jwt.ParseWithClaims(jwtCookie.Value, &claims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			slog.Error("Failed to unsign jwt")
			return "", ErrorUnauthorized
		}
		if expTime, err := t.Claims.GetExpirationTime(); err != nil || expTime.Before(time.Now()) {
			slog.Error("Token expired or GetExpirationTime error", ERROR, err)
			return "", ErrorUnauthorized
		}
		return []byte(jwtUserSecretKey), nil
	})
	if err != nil {
		slog.Error("ParseWithClaims", ERROR, err)
		return pgtype.UUID{}, ErrorUnauthorized
	}
	if !parsedToken.Valid {
		slog.Error("Invalid token")
		return pgtype.UUID{}, ErrorUnauthorized
	}
	return claims.UserID, nil
}

func DeleteJWTCookie() *http.Cookie {
	return &http.Cookie{
		Name:     JWT_COOKIE,
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	}
}
