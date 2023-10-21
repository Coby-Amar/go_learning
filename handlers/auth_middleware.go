package handlers

import (
	"log/slog"
	"net/http"
	"strings"
)

func (conf *ApiConfig) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "auth") {
			next.ServeHTTP(w, r)
			return
		}
		session, err := conf.STORE.Get(r, SESSION)
		if err != nil {
			slog.Error("Error getting session", ERROR, err)
			respondWithUnauthorized(w)
			return
		}

		if _, ok := session.Values[SESSION_PARAMETERS]; ok {
			next.ServeHTTP(w, r)
			return
		}
		jwtCookie, err := r.Cookie(JWT_COOKIE)
		if err != nil {
			slog.Error("Error getting JWT cookie", ERROR, err)
			respondWithUnauthorized(w)
			return
		}
		user_id, validationErr := validateJWT(jwtCookie, conf.JWT_SECRET_KEY)
		if validationErr != nil {
			slog.Error("Error validating JWT cookie", ERROR, validationErr)
			respondWithUnauthorized(w)
			return
		}
		user, err := conf.DB.GetUserByID(r.Context(), user_id)
		if err != nil {
			slog.Error("Error getting userfrom DB", ERROR, err)
			respondWithUnauthorized(w)
			return
		}
		ok := conf.createUserSession(w, r, user.ID)
		if !ok {
			slog.Error("Error creating user", ERROR, err)
			respondWithUnauthorized(w)
			return
		}
	})

}
