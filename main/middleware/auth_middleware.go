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
			if strings.Contains(r.URL.Path, "auth") {
				next.ServeHTTP(w, r)
				return
			}
			configWithRequestAndResponse := &utils.ConfigWithRequestAndResponse{
				Config: config,
				W:      w,
				R:      r,
			}
			if sessionParams := utils.GetSessionParams(configWithRequestAndResponse); sessionParams != nil {
				next.ServeHTTP(w, r)
				return
			}
			jwtCookie, err := r.Cookie(utils.JWT_COOKIE)
			if err != nil {
				slog.Error("Error getting JWT cookie", utils.ERROR, err)
				utils.RespondWithUnauthorized(w)
				return
			}
			user_id, validationErr := utils.ValidateJWT(jwtCookie, config.JWT_SECRET_KEY)
			if validationErr != nil {
				slog.Error("Error validating JWT cookie", utils.ERROR, validationErr)
				utils.RespondWithUnauthorized(w)
				return
			}
			user, err := config.DB.GetUserByID(r.Context(), user_id)
			if err != nil {
				slog.Error("Error getting userfrom DB", utils.ERROR, err)
				utils.RespondWithUnauthorized(w)
				return
			}
			ok := utils.CreateUserSession(configWithRequestAndResponse, user.ID)
			if !ok {
				slog.Error("Error creating user", utils.ERROR, err)
				utils.RespondWithUnauthorized(w)
				return
			}
		})
	}
}
