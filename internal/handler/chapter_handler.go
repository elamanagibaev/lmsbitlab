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

// Create godoc
// @Summary      Создать главу
// @Description  Создаёт новую главу внутри курса
// @Tags         chapters
// @Accept       json
// @Produce      json
// @Param        input  body      dto.CreateChapterDTO  true  "Данные главы"
// @Success      201    {object}  api.Response{data=dto.ChapterResponseDTO}
// @Failure      400    {object}  api.Response
// @Failure      500    {object}  api.Response
// @Router       /chapters [post]
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

// GetByID godoc
// @Summary      Получить главу по ID
// @Description  Возвращает главу по её идентификатору
// @Tags         chapters
// @Produce      json
// @Param        id   path      int  true  "ID главы"
// @Success      200  {object}  api.Response{data=dto.ChapterResponseDTO}
// @Failure      400  {object}  api.Response
// @Failure      404  {object}  api.Response
// @Router       /chapters/{id} [get]
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

// GetByCourseID godoc
// @Summary      Получить главы курса
// @Description  Возвращает список глав конкретного курса
// @Tags         chapters
// @Produce      json
// @Param        id   path      int  true  "ID курса"
// @Success      200  {object}  api.Response{data=[]dto.ChapterResponseDTO}
// @Failure      400  {object}  api.Response
// @Failure      500  {object}  api.Response
// @Router       /courses/{id}/chapters [get]
func (h *ChapterHandler) GetByCourseID(c *gin.Context) {
	courseID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			Success: false,
			Error:   "invalid course id",
		})
		return
	}

	result, err := h.service.GetByCourseID(uint(courseID))
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
// @Summary      Обновить главу
// @Description  Обновляет существующую главу по ID
// @Tags         chapters
// @Accept       json
// @Produce      json
// @Param        id     path      int                   true  "ID главы"
// @Param        input  body      dto.UpdateChapterDTO  true  "Обновлённые данные главы"
// @Success      200    {object}  api.Response{data=dto.ChapterResponseDTO}
// @Failure      400    {object}  api.Response
// @Failure      404    {object}  api.Response
// @Failure      500    {object}  api.Response
// @Router       /chapters/{id} [put]
func (h *ChapterHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			Success: false,
			Error:   "invalid chapter id",
		})
		return
	}

	var input dto.UpdateChapterDTO
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
// @Summary      Удалить главу
// @Description  Удаляет главу по ID
// @Tags         chapters
// @Produce      json
// @Param        id   path      int  true  "ID главы"
// @Success      200  {object}  api.Response
// @Failure      400  {object}  api.Response
// @Failure      404  {object}  api.Response
// @Failure      500  {object}  api.Response
// @Router       /chapters/{id} [delete]
func (h *ChapterHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			Success: false,
			Error:   "invalid chapter id",
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
