package builder

import (
	"encoding/json"
	"net/http"

	"github.com/juandmca/go-core-error-handling/v2/src/error/model"
	"github.com/juandmca/go-core-error-handling/v2/src/logger"
)

// Funcion que construye un nuevo RubikError en base a otro error
func BuildRubikError(r *http.Request, statusCode int, friendlyMessage string, technicalMessage string,
	detail []model.SauronErrorDetail, errorCategory string) *model.SauronError {

	rubikError := &model.SauronError{
		StatusCode:       statusCode,
		FriendlyMessage:  friendlyMessage,
		TechnicalMessage: technicalMessage,
		ErrorCategory:    errorCategory,
		ErrorDetail:      detail,
		Path:             r.URL.Path,
	}
	logger.LogMessage(r, rubikError)
	return rubikError
}

func BuildDefaultResponse(rw http.ResponseWriter, data interface{}, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	output, _ := json.Marshal(&data)
	rw.Write(output)
}
