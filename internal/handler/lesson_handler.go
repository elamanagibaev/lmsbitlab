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

// Create godoc
// @Summary      Создать урок
// @Description  Создаёт новый урок внутри главы
// @Tags         lessons
// @Accept       json
// @Produce      json
// @Param        input  body      dto.CreateLessonDTO  true  "Данные урока"
// @Success      201    {object}  api.Response{data=dto.LessonResponseDTO}
// @Failure      400    {object}  api.Response
// @Failure      500    {object}  api.Response
// @Router       /lessons [post]
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

// GetByID godoc
// @Summary      Получить урок по ID
// @Description  Возвращает урок по его идентификатору
// @Tags         lessons
// @Produce      json
// @Param        id   path      int  true  "ID урока"
// @Success      200  {object}  api.Response{data=dto.LessonResponseDTO}
// @Failure      400  {object}  api.Response
// @Failure      404  {object}  api.Response
// @Router       /lessons/{id} [get]
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

// GetByChapterID godoc
// @Summary      Получить уроки главы
// @Description  Возвращает список уроков конкретной главы
// @Tags         lessons
// @Produce      json
// @Param        id   path      int  true  "ID главы"
// @Success      200  {object}  api.Response{data=[]dto.LessonResponseDTO}
// @Failure      400  {object}  api.Response
// @Failure      500  {object}  api.Response
// @Router       /chapters/{id}/lessons [get]
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

// Update godoc
// @Summary      Обновить урок
// @Description  Обновляет существующий урок по ID
// @Tags         lessons
// @Accept       json
// @Produce      json
// @Param        id     path      int                  true  "ID урока"
// @Param        input  body      dto.UpdateLessonDTO  true  "Обновлённые данные урока"
// @Success      200    {object}  api.Response{data=dto.LessonResponseDTO}
// @Failure      400    {object}  api.Response
// @Failure      404    {object}  api.Response
// @Failure      500    {object}  api.Response
// @Router       /lessons/{id} [put]
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

// Delete godoc
// @Summary      Удалить урок
// @Description  Удаляет урок по ID
// @Tags         lessons
// @Produce      json
// @Param        id   path      int  true  "ID урока"
// @Success      200  {object}  api.Response
// @Failure      400  {object}  api.Response
// @Failure      404  {object}  api.Response
// @Failure      500  {object}  api.Response
// @Router       /lessons/{id} [delete]
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
