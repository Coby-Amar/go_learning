package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

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

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, errorResponse{
		Error: message,
	})

}
