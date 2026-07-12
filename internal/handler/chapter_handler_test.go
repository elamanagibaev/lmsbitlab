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

func setupChapterRouter(h *handler.ChapterHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(api.ErrorMiddleware())

	router.GET("/courses/:id/chapters", h.GetByCourseID)

	chapters := router.Group("/chapters")
	{
		chapters.POST("", h.Create)
		chapters.GET("/:id", h.GetByID)
		chapters.PUT("/:id", h.Update)
		chapters.DELETE("/:id", h.Delete)
	}

	return router
}

type chapterResponse struct {
	Success bool                   `json:"success"`
	Data    dto.ChapterResponseDTO `json:"data"`
	Error   string                 `json:"error"`
}

type chapterListResponse struct {
	Success bool                     `json:"success"`
	Data    []dto.ChapterResponseDTO `json:"data"`
	Error   string                   `json:"error"`
}

func TestChapterHandler_Create_Success(t *testing.T) {
	serviceMock := mocks.NewMockChapterService(t)

	input := dto.CreateChapterDTO{Name: "Control Structures", Order: 1, CourseID: 1}
	expected := dto.ChapterResponseDTO{ID: 1, Name: "Control Structures", Order: 1, CourseID: 1}

	serviceMock.EXPECT().
		Create(input).
		Return(expected, nil).
		Once()

	h := handler.NewChapterHandler(serviceMock)
	router := setupChapterRouter(h)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/chapters", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)

	var resp chapterResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, uint(1), resp.Data.ID)
}

func TestChapterHandler_Create_InvalidBody(t *testing.T) {
	serviceMock := mocks.NewMockChapterService(t)

	h := handler.NewChapterHandler(serviceMock)
	router := setupChapterRouter(h)

	// name отсутствует — required
	body := []byte(`{"order":1,"course_id":1}`)
	req := httptest.NewRequest(http.MethodPost, "/chapters", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestChapterHandler_GetByID_Success(t *testing.T) {
	serviceMock := mocks.NewMockChapterService(t)

	expected := dto.ChapterResponseDTO{ID: 1, Name: "Control Structures"}

	serviceMock.EXPECT().
		GetByID(uint(1)).
		Return(expected, nil).
		Once()

	h := handler.NewChapterHandler(serviceMock)
	router := setupChapterRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/chapters/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestChapterHandler_GetByID_NotFound(t *testing.T) {
	serviceMock := mocks.NewMockChapterService(t)

	serviceMock.EXPECT().
		GetByID(uint(999)).
		Return(dto.ChapterResponseDTO{}, apperrors.ErrChapterNotFound).
		Once()

	h := handler.NewChapterHandler(serviceMock)
	router := setupChapterRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/chapters/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)

	var resp chapterResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "chapter not found", resp.Error)
}

func TestChapterHandler_GetByCourseID_Success(t *testing.T) {
	serviceMock := mocks.NewMockChapterService(t)

	expected := []dto.ChapterResponseDTO{
		{ID: 1, Name: "Control Structures", CourseID: 1},
		{ID: 2, Name: "Data Types", CourseID: 1},
	}

	serviceMock.EXPECT().
		GetByCourseID(uint(1)).
		Return(expected, nil).
		Once()

	h := handler.NewChapterHandler(serviceMock)
	router := setupChapterRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/courses/1/chapters", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp chapterListResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp.Data, 2)
}

func TestChapterHandler_Update_Success(t *testing.T) {
	serviceMock := mocks.NewMockChapterService(t)

	input := dto.UpdateChapterDTO{Name: "New Name", Order: 2}
	expected := dto.ChapterResponseDTO{ID: 1, Name: "New Name", Order: 2}

	serviceMock.EXPECT().
		Update(uint(1), input).
		Return(expected, nil).
		Once()

	h := handler.NewChapterHandler(serviceMock)
	router := setupChapterRouter(h)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPut, "/chapters/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestChapterHandler_Update_NotFound(t *testing.T) {
	serviceMock := mocks.NewMockChapterService(t)

	input := dto.UpdateChapterDTO{Name: "New Name", Order: 2}

	serviceMock.EXPECT().
		Update(uint(999), input).
		Return(dto.ChapterResponseDTO{}, apperrors.ErrChapterNotFound).
		Once()

	h := handler.NewChapterHandler(serviceMock)
	router := setupChapterRouter(h)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPut, "/chapters/999", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestChapterHandler_Delete_Success(t *testing.T) {
	serviceMock := mocks.NewMockChapterService(t)

	serviceMock.EXPECT().
		Delete(uint(1)).
		Return(nil).
		Once()

	h := handler.NewChapterHandler(serviceMock)
	router := setupChapterRouter(h)

	req := httptest.NewRequest(http.MethodDelete, "/chapters/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestChapterHandler_Delete_NotFound(t *testing.T) {
	serviceMock := mocks.NewMockChapterService(t)

	serviceMock.EXPECT().
		Delete(uint(999)).
		Return(apperrors.ErrChapterNotFound).
		Once()

	h := handler.NewChapterHandler(serviceMock)
	router := setupChapterRouter(h)

	req := httptest.NewRequest(http.MethodDelete, "/chapters/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}
