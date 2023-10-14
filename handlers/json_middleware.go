package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
)

type jsonHandler[T interface{}] func(http.ResponseWriter, *http.Request, T)
type ParamsType[T interface{}] struct {
	params T
}

func ParseJSONAndValidateMiddleware[T interface{}](handler jsonHandler[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paramsContainer := ParamsType[T]{}
		decodeErr := json.NewDecoder(r.Body).Decode(&paramsContainer.params)
		if decodeErr != nil {
			log.Println("Decode body error:", decodeErr)
			if strings.Contains(decodeErr.Error(), "cannot unmarshal") {
				respondWithError(w, http.StatusBadRequest, "Malformed request")
			} else {
				respondWithError(w, http.StatusInternalServerError, "")
			}
			return
		}
		if validationErr := validator.New().Struct(paramsContainer.params); validationErr != nil {
			log.Println("Validate body error: ", validationErr)
			respondWithError(w, http.StatusBadRequest, "Malformed request")
			return
		}
		handler(w, r, paramsContainer.params)
	}

}
