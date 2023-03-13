package middleware

import (
	"net/http"
	"os"

	over "github.com/Trendyol/overlog"
	"github.com/juandmca/go-core-error-handling/v2/src/builder"
	"github.com/juandmca/go-core-error-handling/v2/src/error/constants"
	"github.com/juandmca/go-core-error-handling/v2/src/error/model"
	"github.com/mercadolibre/fury_go-core/pkg/web"
	"github.com/rs/zerolog"
)

func init() {
	zlogger := zerolog.New(os.Stderr).With().Str("foo", "bar").Logger()
	over.New(zlogger)
}

func RubikLogger() web.Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			over.NewDefault()
			over.Log().Info("Operation Start: " + r.URL.Path)
		}
	}
}

// Funcion que valida la estructura de las cabeceras asegurandose que vengan los valores
// requeridos para una correcta trazabilidad de la peticion
func HeaderValidator() web.Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			headers := []string{"x-request-id"}
			detail := []model.RubikErrorDetail{}
			for _, header := range headers {
				if exist := r.Header.Get(header); exist == "" {
					detail = append(detail, model.RubikErrorDetail{
						ErrorDescription: "header field missing:" + header,
						ErrorComponent:   "Error checcking header structure",
					})
				}
			}
			if len(detail) > 0 {
				rubik_error := builder.BuildRubikError(r, http.StatusBadRequest, "An unexpected error happened when checking your request",
					"Missing or incorrect headers in the request", detail, constants.TECHNICAL_ERROR)
				builder.BuildDefaultResponse(w, rubik_error, http.StatusBadRequest)
			}
		}
	}
}

// func RubikLooger() {

// 	return func(c *gin.Context) {
// 		over.NewDefault()
// 		over.Log().Info("Operation Start: " + c.FullPath())

// 		//Header check
// 		if header_error := headerChecker(c); header_error != nil {
// 			over.Log().Error("header-check-error", header_error)
// 			//c.JSON(http.StatusBadRequest, &header_error)
// 			c.AbortWithStatusJSON(http.StatusBadRequest, &header_error)
// 			return
// 		}
// 		//Header Logger
// 		over.MDC().Set("uuid", c.GetHeader("UUID"))

// 		if request_body, error := c.GetRawData(); error != nil {
// 			over.Log().Error("Error Obteniendo Body de la Peticion", error)
// 		} else {
// 			request_body, _ := json.Marshal(&request_body)
// 			over.MDC().Set("request-body", request_body)
// 		}
// 		over.MDC().Set("query-parameters", c.Request.URL.Query())
// 		c.Next()
// 	}
// }

// func headerChecker(c *gin.Context) *model.RubikError {
// 	over.NewDefault()
// 	over.Log().Info("Start Header Checker:")
// 	kh := c.GetHeader("UUID")
// 	if kh == "" {
// 		detail := []model.RubikErrorDetail{
// 			{
// 				CustomCode:       "AVSBD12",
// 				ErrorDescription: "Error in header validation: UUID header is missing",
// 			},
// 		}
// 		over.Log().Error("Finish Header Checking With Error")
// 		return builder.BuildRubikError(c, http.StatusBadRequest, "error amistoso", "technical error", detail, constants.BUSINESS_ERROR)
// 	}
// 	over.Log().Error("Finish Header Checking With Success")
// 	return nil
// }
