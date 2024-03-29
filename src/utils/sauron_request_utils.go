package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/juandmca/go-core-error-handling/v2/src/builder"
	"github.com/juandmca/go-core-error-handling/v2/src/error/constants"
	"github.com/juandmca/go-core-error-handling/v2/src/error/model"
)

// Funcion que valida la estructura del request de una peticion y mapea
// los errores de validacion en una estructura SauronError
func RequestValidator(r *http.Request, data interface{}) *model.SauronError {
	var sauronError *model.SauronError
	detail := []model.SauronErrorDetail{}
	if decodeError := json.NewDecoder(r.Body).Decode(&data); decodeError != nil {
		detail = append(detail, model.SauronErrorDetail{
			ErrorDescription: "error reading the body structure",
			ErrorComponent:   decodeError.Error(),
		})
	}
	if data != nil {
		validate := validator.New()
		if requestErrors := validate.Struct(data); requestErrors != nil {
			for _, requestError := range strings.Split(requestErrors.Error(), "\n") {
				detail = append(detail, model.SauronErrorDetail{
					ErrorDescription: "Missing or incorrect structure of body property",
					ErrorComponent:   requestError,
				})
			}
		}
	}
	if len(detail) > 0 {
		sauronError = &model.SauronError{
			StatusCode:       http.StatusBadRequest,
			FriendlyMessage:  "An unexpected error happened when checking your request",
			TechnicalMessage: "Incorrect body structure",
			ErrorCategory:    constants.TECHNICAL_ERROR,
		}
		return builder.BuildSauronError(r, sauronError, detail)
	} else {
		return sauronError
	}
}
