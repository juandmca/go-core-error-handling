package builder

import (
	"encoding/json"
	"fmt"
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

func BuildDefaultResponse(rw http.ResponseWriter, data interface{}, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	output, _ := json.Marshal(&data)
	fmt.Fprintln(rw, string(output))
}
