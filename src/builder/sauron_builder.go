package builder

import (
	"encoding/json"
	"net/http"

	"github.com/juandmca/go-core-error-handling/v2/src/error/model"
	"github.com/juandmca/go-core-error-handling/v2/src/logger"
)

// Funcion que construye un nuevo SauronError en base a otro error
func BuildSauronError(r *http.Request, sauronError *model.SauronError,
	detail []model.SauronErrorDetail) *model.SauronError {
	sauronError.Path = r.URL.Path
	if detail != nil {
		sauronError.ErrorDetail = detail
	}
	logger.LogMessage(r, sauronError)
	return sauronError
}

// Funcion que arma una respuesta por defecto para una api tomando como base un objeto preconstruido
func BuildDefaultResponse(rw http.ResponseWriter, data interface{}, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	output, _ := json.Marshal(&data)
	rw.Write(output)
}
