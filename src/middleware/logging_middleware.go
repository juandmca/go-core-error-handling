package middleware

import (
	"encoding/json"

	over "github.com/Trendyol/overlog"
	"github.com/gin-gonic/gin"
)

func RubikLooger() func(c *gin.Context) {
	return func(c *gin.Context) {
		over.NewDefault()
		over.Log().Info("Inicio de ejecucion operacion:" + c.FullPath())

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
