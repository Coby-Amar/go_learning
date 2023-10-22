package handlers

import (
	"net/http"

	"github.com/coby-amar/go_learning/main/utils"
)

func HandleLogout(cwrar *utils.ConfigWithRequestAndResponse) {
	utils.RespondWithJSON(cwrar.W, http.StatusOK, struct{}{})
}
