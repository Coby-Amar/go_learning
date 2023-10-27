package handlers

import (
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/main/utils"
)

func HandleHealthZ(cwrar *utils.ConfigWithRequestAndResponse) {
	slog.Info("HandleHealthZ")
	if session, err := cwrar.Config.Store.Get(cwrar.R, utils.SESSION); err == nil {
		if _, ok := session.Values[utils.SESSION_PARAMETERS]; ok {
			utils.RespondWithMessage(cwrar.W, http.StatusOK, "in session")
			return
		}
	}
	if jwtCookie, err := cwrar.R.Cookie(utils.JWT_COOKIE); err == nil {
		_, validationErr := utils.ValidateJWT(jwtCookie, cwrar.Config.JWT_SECRET_KEY)
		if validationErr == nil {
			utils.RespondWithMessage(cwrar.W, http.StatusOK, "in session")
			return
		}
	}
	utils.RespondWithUnauthorized(cwrar.W)
}
