package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"LMSBitLab/internal/api"
	"LMSBitLab/internal/apperrors"
	"LMSBitLab/internal/dto"
	"LMSBitLab/internal/handler"
	"LMSBitLab/internal/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupCourseRouter(h *handler.CourseHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(api.ErrorMiddleware())

	courses := router.Group("/courses")
	{
		courses.POST("", h.Create)
		courses.GET("", h.GetAll)
		courses.GET("/:id", h.GetByID)
		courses.PUT("/:id", h.Update)
		courses.DELETE("/:id", h.Delete)
	}

	return router
}

type courseResponse struct {
	Success bool                  `json:"success"`
	Data    dto.CourseResponseDTO `json:"data"`
	Error   string                `json:"error"`
}

type courseListResponse struct {
	Success bool                    `json:"success"`
	Data    []dto.CourseResponseDTO `json:"data"`
	Error   string                  `json:"error"`
}

func TestCourseHandler_Create_Success(t *testing.T) {
	serviceMock := mocks.NewMockCourseService(t)

	expected := dto.CourseResponseDTO{ID: 1, Name: "Go Developer", Description: "desc"}

	serviceMock.EXPECT().
		Create(dto.CreateCourseDTO{Name: "Go Developer", Description: "desc"}).
		Return(expected, nil).
		Once()

	h := handler.NewCourseHandler(serviceMock)
	router := setupCourseRouter(h)

	body, _ := json.Marshal(dto.CreateCourseDTO{Name: "Go Developer", Description: "desc"})
	req := httptest.NewRequest(http.MethodPost, "/courses", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)

	var resp courseResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.True(t, resp.Success)
	assert.Equal(t, uint(1), resp.Data.ID)
	assert.Equal(t, "Go Developer", resp.Data.Name)
}

func TestCourseHandler_Create_InvalidBody(t *testing.T) {
	serviceMock := mocks.NewMockCourseService(t)

	h := handler.NewCourseHandler(serviceMock)
	router := setupCourseRouter(h)

	// name отсутствует, а он "required" в DTO — ожидаем 400 ещё до вызова сервиса
	body := []byte(`{"description":"desc"}`)
	req := httptest.NewRequest(http.MethodPost, "/courses", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCourseHandler_GetByID_Success(t *testing.T) {
	serviceMock := mocks.NewMockCourseService(t)

	expected := dto.CourseResponseDTO{ID: 1, Name: "Go Developer"}

	serviceMock.EXPECT().
		GetByID(uint(1)).
		Return(expected, nil).
		Once()

	h := handler.NewCourseHandler(serviceMock)
	router := setupCourseRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/courses/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp courseResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "Go Developer", resp.Data.Name)
}

func TestCourseHandler_GetByID_NotFound(t *testing.T) {
	serviceMock := mocks.NewMockCourseService(t)

	serviceMock.EXPECT().
		GetByID(uint(999)).
		Return(dto.CourseResponseDTO{}, apperrors.ErrCourseNotFound).
		Once()

	h := handler.NewCourseHandler(serviceMock)
	router := setupCourseRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/courses/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)

	var resp courseResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.False(t, resp.Success)
	assert.Equal(t, "course not found", resp.Error)
}

func TestCourseHandler_GetByID_InvalidID(t *testing.T) {
	serviceMock := mocks.NewMockCourseService(t)

	h := handler.NewCourseHandler(serviceMock)
	router := setupCourseRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/courses/abc", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCourseHandler_GetAll_Success(t *testing.T) {
	serviceMock := mocks.NewMockCourseService(t)

	expected := []dto.CourseResponseDTO{
		{ID: 1, Name: "Go Developer"},
		{ID: 2, Name: "Python Developer"},
	}

	serviceMock.EXPECT().
		GetAll().
		Return(expected, nil).
		Once()

	h := handler.NewCourseHandler(serviceMock)
	router := setupCourseRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/courses", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp courseListResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp.Data, 2)
}

func TestCourseHandler_Update_Success(t *testing.T) {
	serviceMock := mocks.NewMockCourseService(t)

	input := dto.UpdateCourseDTO{Name: "New Name", Description: "new desc"}
	expected := dto.CourseResponseDTO{ID: 1, Name: "New Name", Description: "new desc"}

	serviceMock.EXPECT().
		Update(uint(1), input).
		Return(expected, nil).
		Once()

	h := handler.NewCourseHandler(serviceMock)
	router := setupCourseRouter(h)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPut, "/courses/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp courseResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "New Name", resp.Data.Name)
}

func TestCourseHandler_Update_NotFound(t *testing.T) {
	serviceMock := mocks.NewMockCourseService(t)

	input := dto.UpdateCourseDTO{Name: "New Name", Description: "new desc"}

	serviceMock.EXPECT().
		Update(uint(999), input).
		Return(dto.CourseResponseDTO{}, apperrors.ErrCourseNotFound).
		Once()

	h := handler.NewCourseHandler(serviceMock)
	router := setupCourseRouter(h)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPut, "/courses/999", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestCourseHandler_Delete_Success(t *testing.T) {
	serviceMock := mocks.NewMockCourseService(t)

	serviceMock.EXPECT().
		Delete(uint(1)).
		Return(nil).
		Once()

	h := handler.NewCourseHandler(serviceMock)
	router := setupCourseRouter(h)

	req := httptest.NewRequest(http.MethodDelete, "/courses/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestCourseHandler_Delete_NotFound(t *testing.T) {
	serviceMock := mocks.NewMockCourseService(t)

	serviceMock.EXPECT().
		Delete(uint(999)).
		Return(apperrors.ErrCourseNotFound).
		Once()

	h := handler.NewCourseHandler(serviceMock)
	router := setupCourseRouter(h)

	req := httptest.NewRequest(http.MethodDelete, "/courses/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}
