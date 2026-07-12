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

func setupLessonRouter(h *handler.LessonHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(api.ErrorMiddleware())

	router.GET("/chapters/:id/lessons", h.GetByChapterID)

	lessons := router.Group("/lessons")
	{
		lessons.POST("", h.Create)
		lessons.GET("/:id", h.GetByID)
		lessons.PUT("/:id", h.Update)
		lessons.DELETE("/:id", h.Delete)
	}

	return router
}

type lessonResponse struct {
	Success bool                  `json:"success"`
	Data    dto.LessonResponseDTO `json:"data"`
	Error   string                `json:"error"`
}

type lessonListResponse struct {
	Success bool                    `json:"success"`
	Data    []dto.LessonResponseDTO `json:"data"`
	Error   string                  `json:"error"`
}

func TestLessonHandler_Create_Success(t *testing.T) {
	serviceMock := mocks.NewMockLessonService(t)

	input := dto.CreateLessonDTO{Name: "If-else Statement", Content: "текст", Order: 1, ChapterID: 1}
	expected := dto.LessonResponseDTO{ID: 1, Name: "If-else Statement", ChapterID: 1}

	serviceMock.EXPECT().
		Create(input).
		Return(expected, nil).
		Once()

	h := handler.NewLessonHandler(serviceMock)
	router := setupLessonRouter(h)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/lessons", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)

	var resp lessonResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, uint(1), resp.Data.ID)
}

func TestLessonHandler_Create_InvalidBody(t *testing.T) {
	serviceMock := mocks.NewMockLessonService(t)

	h := handler.NewLessonHandler(serviceMock)
	router := setupLessonRouter(h)

	// content отсутствует — required
	body := []byte(`{"name":"If-else Statement","order":1,"chapter_id":1}`)
	req := httptest.NewRequest(http.MethodPost, "/lessons", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLessonHandler_GetByID_Success(t *testing.T) {
	serviceMock := mocks.NewMockLessonService(t)

	expected := dto.LessonResponseDTO{ID: 1, Name: "If-else Statement"}

	serviceMock.EXPECT().
		GetByID(uint(1)).
		Return(expected, nil).
		Once()

	h := handler.NewLessonHandler(serviceMock)
	router := setupLessonRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/lessons/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestLessonHandler_GetByID_NotFound(t *testing.T) {
	serviceMock := mocks.NewMockLessonService(t)

	serviceMock.EXPECT().
		GetByID(uint(999)).
		Return(dto.LessonResponseDTO{}, apperrors.ErrLessonNotFound).
		Once()

	h := handler.NewLessonHandler(serviceMock)
	router := setupLessonRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/lessons/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)

	var resp lessonResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "lesson not found", resp.Error)
}

func TestLessonHandler_GetByChapterID_Success(t *testing.T) {
	serviceMock := mocks.NewMockLessonService(t)

	expected := []dto.LessonResponseDTO{
		{ID: 1, Name: "If-else Statement", ChapterID: 1},
		{ID: 2, Name: "Switch Statement", ChapterID: 1},
	}

	serviceMock.EXPECT().
		GetByChapterID(uint(1)).
		Return(expected, nil).
		Once()

	h := handler.NewLessonHandler(serviceMock)
	router := setupLessonRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/chapters/1/lessons", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp lessonListResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp.Data, 2)
}

func TestLessonHandler_Update_Success(t *testing.T) {
	serviceMock := mocks.NewMockLessonService(t)

	input := dto.UpdateLessonDTO{Name: "New Name", Content: "новый текст", Order: 2}
	expected := dto.LessonResponseDTO{ID: 1, Name: "New Name", Order: 2}

	serviceMock.EXPECT().
		Update(uint(1), input).
		Return(expected, nil).
		Once()

	h := handler.NewLessonHandler(serviceMock)
	router := setupLessonRouter(h)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPut, "/lessons/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestLessonHandler_Update_NotFound(t *testing.T) {
	serviceMock := mocks.NewMockLessonService(t)

	input := dto.UpdateLessonDTO{Name: "New Name", Content: "новый текст", Order: 2}

	serviceMock.EXPECT().
		Update(uint(999), input).
		Return(dto.LessonResponseDTO{}, apperrors.ErrLessonNotFound).
		Once()

	h := handler.NewLessonHandler(serviceMock)
	router := setupLessonRouter(h)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPut, "/lessons/999", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestLessonHandler_Delete_Success(t *testing.T) {
	serviceMock := mocks.NewMockLessonService(t)

	serviceMock.EXPECT().
		Delete(uint(1)).
		Return(nil).
		Once()

	h := handler.NewLessonHandler(serviceMock)
	router := setupLessonRouter(h)

	req := httptest.NewRequest(http.MethodDelete, "/lessons/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestLessonHandler_Delete_NotFound(t *testing.T) {
	serviceMock := mocks.NewMockLessonService(t)

	serviceMock.EXPECT().
		Delete(uint(999)).
		Return(apperrors.ErrLessonNotFound).
		Once()

	h := handler.NewLessonHandler(serviceMock)
	router := setupLessonRouter(h)

	req := httptest.NewRequest(http.MethodDelete, "/lessons/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}
