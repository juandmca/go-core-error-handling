package builder

import (
	"net/http"

	"github.com/juandmca/go-core-error-handling/v2/src/error/model"
)

// Funcion que construye un nuevo RubikError en base a otro error
func BuildRubikError(r *http.Request, statusCode int, friendlyMessage string, technicalMessage string, detail []model.RubikErrorDetail, errorCategory string) *model.RubikError {

	return &model.RubikError{
		StatusCode:       statusCode,
		FriendlyMessage:  friendlyMessage,
		TechnicalMessage: technicalMessage,
		ErrorCategory:    errorCategory,
		ErrorDetail:      detail,
		Path:             r.URL.Path,
	}
}
