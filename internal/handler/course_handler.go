package handler

import (
	"net/http"
	"strconv"

	"LMSBitLab/internal/api"
	"LMSBitLab/internal/dto"
	"LMSBitLab/internal/service"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	service service.CourseService
}

func NewCourseHandler(service service.CourseService) *CourseHandler {
	return &CourseHandler{service: service}
}

// Create godoc
// @Summary      Создать курс
// @Description  Создаёт новый курс
// @Tags         courses
// @Accept       json
// @Produce      json
// @Param        input  body      dto.CreateCourseDTO  true  "Данные курса"
// @Success      201    {object}  api.Response{data=dto.CourseResponseDTO}
// @Failure      400    {object}  api.Response
// @Failure      500    {object}  api.Response
// @Router       /courses [post]
func (h *CourseHandler) Create(c *gin.Context) {
	var input dto.CreateCourseDTO
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
// @Summary      Получить курс по ID
// @Description  Возвращает курс по его идентификатору
// @Tags         courses
// @Produce      json
// @Param        id   path      int  true  "ID курса"
// @Success      200  {object}  api.Response{data=dto.CourseResponseDTO}
// @Failure      400  {object}  api.Response
// @Failure      404  {object}  api.Response
// @Router       /courses/{id} [get]
func (h *CourseHandler) GetByID(c *gin.Context) {
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

// GetAll godoc
// @Summary      Получить все курсы
// @Description  Возвращает список всех курсов
// @Tags         courses
// @Produce      json
// @Success      200  {object}  api.Response{data=[]dto.CourseResponseDTO}
// @Failure      500  {object}  api.Response
// @Router       /courses [get]
func (h *CourseHandler) GetAll(c *gin.Context) {
	result, err := h.service.GetAll()
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
// @Summary      Обновить курс
// @Description  Обновляет существующий курс по ID
// @Tags         courses
// @Accept       json
// @Produce      json
// @Param        id     path      int                   true  "ID курса"
// @Param        input  body      dto.UpdateCourseDTO   true  "Обновлённые данные курса"
// @Success      200    {object}  api.Response{data=dto.CourseResponseDTO}
// @Failure      400    {object}  api.Response
// @Failure      404    {object}  api.Response
// @Failure      500    {object}  api.Response
// @Router       /courses/{id} [put]
func (h *CourseHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			Success: false,
			Error:   "invalid course id",
		})
		return
	}

	var input dto.UpdateCourseDTO
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
// @Summary      Удалить курс
// @Description  Удаляет курс по ID
// @Tags         courses
// @Produce      json
// @Param        id   path      int  true  "ID курса"
// @Success      200  {object}  api.Response
// @Failure      400  {object}  api.Response
// @Failure      404  {object}  api.Response
// @Failure      500  {object}  api.Response
// @Router       /courses/{id} [delete]
func (h *CourseHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.Response{
			Success: false,
			Error:   "invalid course id",
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
