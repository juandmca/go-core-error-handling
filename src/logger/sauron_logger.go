package logger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"

	over "github.com/Trendyol/overlog"
	"github.com/juandmca/go-core-error-handling/v2/src/error/model"
	"github.com/rs/zerolog"
)

// Funcion que inicaliza el logger de transacciones
func init() {
	zlogger := zerolog.New(os.Stderr).With().Logger()
	over.New(zlogger)
}

// Funcion que se encarga de loggear un mensaje dentro del contexto de una sauron app
// Tambien aplica controles PII para obfuscar informacion sensible de los usuarios
func LogMessage(r *http.Request, data interface{}) {
	over.Log().Info("An error was found during the execution of the operation: " + r.URL.Path)
	over.MDC().Set("x-request-id", r.Header.Get("x-request-id"))
	over.AddGlobalFields("x-request-id")
	output, _ := json.Marshal(&data)

	fmt.Println(reflect.ValueOf(&data).Type().String())
	fmt.Println(reflect.TypeOf(model.SauronError{}).String())
	if reflect.ValueOf(&data).Type() == reflect.TypeOf(model.SauronError{}) {
		over.Log().Error(string(output))
	} else {
		over.Log().Info(string(output))
	}
}
