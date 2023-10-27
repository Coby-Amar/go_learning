package utils

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	slog.Info("RespondWithJSON")
	data, err := json.Marshal(payload)
	if err != nil {
		slog.Error("Marshal", "payload", payload)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Contact your IT consultant"))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func RespondWithNotFound(w http.ResponseWriter) {
	RespondWithJSON(w, http.StatusNotFound, struct{}{})

}

func RespondWithMessage(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, errors.New(message))
}

func RespondWithError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}

func RespondWithInternalServerError(w http.ResponseWriter) {
	RespondWithJSON(w, http.StatusInternalServerError, SOMETHING_WENT_WRONG)
}

func RespondWithUnauthorized(w http.ResponseWriter) {
	RespondWithError(w, http.StatusUnauthorized, ErrorUnauthorized)
}

func RespondWithForbidden(w http.ResponseWriter) {
	RespondWithError(w, http.StatusForbidden, ErrorForbidden)
}

func RespondWithBadRequest(w http.ResponseWriter) {
	RespondWithError(w, http.StatusBadRequest, ErrorBadRequest)
}
