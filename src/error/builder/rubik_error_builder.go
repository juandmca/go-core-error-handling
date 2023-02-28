package builder

import (
	"github.com/gin-gonic/gin"
	"github.com/juandmca/go-core-error-handling/v2/src/error/model"
)

// Funcion que construye un nuevo RubikError en base a otro error
func BuildRubikError(c *gin.Context, statusCode int, friendlyMessage string, technicalMessage string, detail []model.RubikErrorDetail, errorCategory string) *model.RubikError {

	rubikError := &model.RubikError{
		StatusCode:       statusCode,
		FriendlyMessage:  friendlyMessage,
		TechnicalMessage: technicalMessage,
		ErrorCategory:    errorCategory,
		ErrorDetail:      detail,
		Path:             c.FullPath(),
	}
	return rubikError
}
