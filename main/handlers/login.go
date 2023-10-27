package handlers

import (
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/main/utils"
	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(cwrar *utils.ConfigWithRequestAndResponse, params utils.LoginJson) {
	slog.Error("HandleLogin")
	user, err := cwrar.Config.Queries.GetUserByEmail(cwrar.R.Context(), params.Username)
	if err != nil {
		slog.Error("GetUserByEmail", utils.ERROR, err)
		utils.RespondWithBadRequest(cwrar.W)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		slog.Error("CompareHashAndPassword", utils.ERROR, err)
		utils.RespondWithMessage(cwrar.W, http.StatusBadRequest, "Username/password don't match")
		return
	}
	cookie := utils.CreateJWTCookie(user.ID, cwrar.Config.JWT_SECRET_KEY)
	if cookie == nil {
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	ok := utils.CreateUserSession(cwrar.Config.Store, cwrar.W, cwrar.R, user.ID)
	if !ok {
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	http.SetCookie(cwrar.W, cookie)
	utils.RespondWithJSON(cwrar.W, http.StatusOK, user)
}
