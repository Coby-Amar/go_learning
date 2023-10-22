package handlers

import (
	"net/http"

	"github.com/coby-amar/go_learning/main/utils"
	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(cwrar *utils.ConfigWithRequestAndResponse, params *LoginJson) {
	user, err := cwrar.Config.DB.GetUserByEmail(cwrar.R.Context(), params.Username)
	if err != nil {
		utils.RespondWithMessage(cwrar.W, http.StatusBadRequest, utils.BAD_REQUEST_DATA)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		utils.RespondWithMessage(cwrar.W, http.StatusBadRequest, utils.BAD_REQUEST_DATA)
		return
	}
	cookie := utils.CreateJWTCookie(user.ID, cwrar.Config.JWT_SECRET_KEY)
	if cookie == nil {
		utils.RespondWithJSON(cwrar.W, http.StatusInternalServerError, utils.SOMETHING_WENT_WRONG)
		return
	}
	ok := utils.CreateUserSession(cwrar, user.ID)
	if !ok {
		utils.RespondWithJSON(cwrar.W, http.StatusInternalServerError, utils.SOMETHING_WENT_WRONG)
		return
	}
	http.SetCookie(cwrar.W, cookie)
	utils.RespondWithJSON(cwrar.W, http.StatusOK, user)
}
