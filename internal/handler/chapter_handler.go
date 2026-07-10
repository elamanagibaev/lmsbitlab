package handler

import (
	"LMSBitLab/internal/api"
	"LMSBitLab/internal/dto"
	"LMSBitLab/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChapterHandler struct {
	service service.ChapterService
}

func NewChapterHandler(service service.ChapterService) *ChapterHandler {
	return &ChapterHandler{service: service}
}

func (h *ChapterHandler) Create(c *gin.Context) {
	var input dto.CreateChapterDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	result, err := h.service.Create(input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, api.Response{
		Success: true,
		Data:    result,
	})
}

func (h *ChapterHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			Success: false,
			Error:   "invalid course id",
		})
		return
	}

	result, err := h.service.GetByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, api.Response{
		Success: true,
		Data:    result,
	})
}
