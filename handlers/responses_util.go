package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		slog.Error("Failed to marshal JSON response", "payload", payload)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Contact your IT consultant"))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithNotFound(w http.ResponseWriter) {
	respondWithJSON(w, http.StatusNotFound, struct{}{})

}

func respondWithMessage(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, errors.New(message))
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}

func respondWithInternalServerError(w http.ResponseWriter) {
	respondWithJSON(w, http.StatusInternalServerError, SOMETHING_WENT_WRONG)
}

func respondWithUnauthorized(w http.ResponseWriter) {
	respondWithError(w, http.StatusUnauthorized, unauthorizedError)
}

func respondWithForbidden(w http.ResponseWriter) {
	respondWithJSON(w, http.StatusForbidden, forbiddenError)
}

func respondWithBadRequest(w http.ResponseWriter) {
	respondWithJSON(w, http.StatusBadRequest, BAD_REQUEST)
}
