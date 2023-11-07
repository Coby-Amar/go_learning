package handlers

import (
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/main/utils"
)

func HandleGetUser(cwrar *utils.ConfigWithRequestAndResponse) {
	slog.Error("HandleGetUser")
	user, err := cwrar.Config.Queries.GetUserByID(cwrar.R.Context(), cwrar.Sparams.UserID)
	if err != nil {
		slog.Error("GetUserByID", utils.ERROR, err)
		utils.RespondWithUnauthorized(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusOK, user)
}
