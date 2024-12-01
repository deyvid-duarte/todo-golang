package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateRequest[S any](input S, w http.ResponseWriter) bool {
	if err := validate.Struct(input); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return true
	}
	return false
}