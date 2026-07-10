package handler

import (
	"LMSBitLab/internal/apperrors"
	"errors"

	"github.com/gin-gonic/gin"
)

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test404(c *gin.Context) {
	c.Error(apperrors.ErrCourseNotFound)
}

func (h *TestHandler) Test500(c *gin.Context) {
	c.Error(errors.New("something went wrong"))
}
