package middleware

import (
	"net/http"

	"github.com/juandmca/go-core-error-handling/v2/src/builder"
	"github.com/juandmca/go-core-error-handling/v2/src/error/constants"
	"github.com/juandmca/go-core-error-handling/v2/src/error/model"
	"github.com/mercadolibre/fury_go-core/pkg/web"
)

// Funcion que valida la estructura de las cabeceras asegurandose que vengan los valores
// requeridos para una correcta trazabilidad de la peticion
func HeaderValidator() web.Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			headers := []string{"x-request-id"}
			detail := []model.SauronErrorDetail{}
			for _, header := range headers {
				if exist := r.Header.Get(header); exist == "" {
					detail = append(detail, model.SauronErrorDetail{
						ErrorDescription: "header field missing:" + header,
						ErrorComponent:   "Error checcking header structure",
					})
				}
			}
			if len(detail) > 0 {
				sauron_error := &model.SauronError{
					StatusCode:       http.StatusBadRequest,
					FriendlyMessage:  "An unexpected error happened when checking your request",
					TechnicalMessage: "Missing or incorrect headers in the request",
					ErrorCategory:    constants.BUSINESS_ERROR,
				}
				builder.BuildDefaultResponse(w, builder.BuildSauronError(r, sauron_error, detail), http.StatusBadRequest)
			} else {
				h.ServeHTTP(w, r)
			}
		}
	}
}
