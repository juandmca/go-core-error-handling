package middleware

import (
	"encoding/json"
	"net/http"

	over "github.com/Trendyol/overlog"
	"github.com/gin-gonic/gin"
	"github.com/juandmca/go-core-error-handling/v2/src/error/builder"
	"github.com/juandmca/go-core-error-handling/v2/src/error/constants"
	"github.com/juandmca/go-core-error-handling/v2/src/error/model"
)

func RubikLooger() func(c *gin.Context) {
	return func(c *gin.Context) {
		over.NewDefault()
		over.Log().Info("Operation Start: " + c.FullPath())

		//Header check
		headerChecker()

		if request_body, error := c.GetRawData(); error != nil {
			over.Log().Error("Error Obteniendo Body de la Peticion", error)
		} else {
			request_body, _ := json.Marshal(&request_body)
			over.MDC().Set("request-body", request_body)
		}
		over.MDC().Set("query-parameters", c.Request.URL.Query())
		c.Next()
	}
}

func headerChecker() func(c *gin.Context) {
	return func(c *gin.Context) {
		over.NewDefault()
		over.Log().Info("Start Header Checker:")
		kh := c.GetHeader("UUID")
		if kh == "" {
			detail := []model.RubikErrorDetail{
				{
					CustomCode:       "AVSBD12",
					ErrorDescription: "Error in header validation: UUID header is missing",
				},
			}
			error := builder.BuildRubikError(c, http.StatusBadRequest, "error amistoso", "technical error", detail, constants.BUSINESS_ERROR)
			c.JSON(error.StatusCode, &error)

			c.JSON(http.StatusBadRequest, &error)
			over.Log().Error("Finish Header Checking With Error: ", error)
			c.Abort()
			return
		}
		over.Log().Error("Finish Header Checking With Success")
		c.Next()
	}
}
