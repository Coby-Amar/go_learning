package middleware

import (
	"net/http"

	"github.com/coby-amar/go_learning/main/utils"
)

func ConfigInjectorMiddleware(config *utils.ApiConfig, next configHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(&utils.ConfigWithRequestAndResponse{
			Config: config,
			W:      w,
			R:      r,
		})
	}
}
