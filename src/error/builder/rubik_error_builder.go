package builder

import (
	"go-core-error-handling/src/error/model"

	"github.com/gin-gonic/gin"
)

// Funcion que construye un nuevo RubikError en base a otro error
func BuildRubikError(c *gin.Context, statusCode int, errorDescription string, detail *model.RubikErrorDetail) *model.RubikError {

	return &model.RubikError{
		Error:            nil,
		StatusCode:       statusCode,
		ErrorDescription: errorDescription,
		ErrorCategory:    "",
		ErrorDetail:      []model.RubikErrorDetail{*detail},
	}
}
