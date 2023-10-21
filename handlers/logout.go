package handlers

import "net/http"

func (conf *ApiConfig) HandleLogout(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}
