package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/coby-amar/go_learning/main/utils"
)

func AuthenticationMiddleware(config *utils.ApiConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			slog.Info("AuthenticationMiddleware")
			if strings.Contains(r.URL.Path, "auth") {
				next.ServeHTTP(w, r)
				return
			}
			if sessionParams := utils.GetSessionParams(config.Store, r); sessionParams != nil {
				next.ServeHTTP(w, r)
				return
			}
			jwtCookie, err := r.Cookie(utils.JWT_COOKIE)
			if err != nil {
				slog.Error("Failed getting JWT cookie", utils.ERROR, err)
				utils.RespondWithUnauthorized(w)
				return
			}
			user_id, validationErr := utils.ValidateJWT(jwtCookie, config.JWT_SECRET_KEY)
			if validationErr != nil {
				utils.RespondWithUnauthorized(w)
				return
			}
			user, err := config.Queries.GetUserByID(r.Context(), user_id)
			if err != nil {
				slog.Error("GetUserByID", utils.ERROR, err)
				utils.RespondWithUnauthorized(w)
				return
			}
			ok := utils.CreateUserSession(config.Store, w, r, user.ID)
			if !ok {
				utils.RespondWithUnauthorized(w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
