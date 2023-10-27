package middleware

import (
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/main/utils"
)

func ConfigInjectorMiddleware(config *utils.ApiConfig, next configHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("ConfigInjectorMiddleware")
		next(&utils.ConfigWithRequestAndResponse{
			Config:  config,
			W:       w,
			R:       r,
			Sparams: utils.GetSessionParams(config.Store, r),
		})
	}
}
