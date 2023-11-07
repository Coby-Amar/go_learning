package handlers

import (
	"net/http"

	"github.com/coby-amar/go_learning/main/utils"
)

func HandleLogout(cwrar *utils.ConfigWithRequestAndResponse) {
	utils.DeleteSessionParams(cwrar.Config.Store, cwrar.W, cwrar.R)
	jwtCookie := utils.DeleteJWTCookie()
	http.SetCookie(cwrar.W, jwtCookie)
	utils.RespondWithMessage(cwrar.W, http.StatusOK, "loggout")
}
