package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator"
)

func parseJSONAndValidateFromRequest[T interface{}](body io.ReadCloser) (T, error) {
	paramsContainer := jsonParams[T]{}
	decodeErr := json.NewDecoder(body).Decode(&paramsContainer.params)
	if decodeErr != nil {
		slog.Error("Failed to decode", ERROR, decodeErr)
		return paramsContainer.params, errors.New("couldnt decode json")
	}
	if validationErr := validator.New().Struct(paramsContainer.params); validationErr != nil {
		slog.Error("Failed to validate", ERROR, validationErr)
		return paramsContainer.params, errors.New("couldnt decode json")
	}
	return paramsContainer.params, nil

}

func ParseJSONAndValidateMiddleware[T interface{}](handler jsonHandler[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseJSONAndValidateFromRequest[T](r.Body)
		if err != nil {
			respondWithBadRequest(w)
			return
		}
		handler(w, r, params)
	}

}
