package logger

import (
	"encoding/json"
	"net/http"
	"os"

	over "github.com/Trendyol/overlog"
	"github.com/rs/zerolog"
)

// Funcion que inicaliza el logger de transacciones
func init() {
	zlogger := zerolog.New(os.Stderr).With().Logger()
	over.New(zlogger)
}

// Funcion que se encarga de loggear un error dentro del contexto de una sauron app
// Tambien aplica controles PII para obfuscar informacion sensible de los usuarios
func LogError(r *http.Request, data interface{}) {
	over.Log().Info("An error was found during the execution of the operation: " + r.URL.Path)
	over.MDC().Set("x-request-id", r.Header.Get("x-request-id"))
	over.AddGlobalFields("x-request-id")
	output, _ := json.Marshal(&data)
	over.Log().Error(output)
}
