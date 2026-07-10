package api

import (
	"LMSBitLab/internal/apperrors"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		status := apperrors.StatusCode(err)

		c.JSON(status, Response{
			Success: false,
			Error:   err.Error(),
		})
	}
}
