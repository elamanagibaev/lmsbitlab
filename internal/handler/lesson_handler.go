package handler

import (
	"LMSBitLab/internal/api"
	"LMSBitLab/internal/dto"
	"LMSBitLab/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LessonHandler struct {
	service service.LessonService
}

func NewLessonHandler(service service.LessonService) *LessonHandler {
	return &LessonHandler{service: service}
}

func (h *LessonHandler) Create(c *gin.Context) {
	var input dto.CreateLessonDTO
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

func (h *LessonHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			Success: false,
			Error:   "invalid lesson id",
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

func (h *LessonHandler) GetByChapterID(c *gin.Context) {
	chapterID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			Success: false,
			Error:   "invalid chapter id",
		})
		return
	}

	result, err := h.service.GetByChapterID(uint(chapterID))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, api.Response{
		Success: true,
		Data:    result,
	})
}

func (h *LessonHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			Success: false,
			Error:   "invalid lesson id",
		})
		return
	}

	var input dto.UpdateLessonDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	result, err := h.service.Update(uint(id), input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, api.Response{
		Success: true,
		Data:    result,
	})
}

func (h *LessonHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			Success: false,
			Error:   "invalid lesson id",
		})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, api.Response{
		Success: true,
	})
}
