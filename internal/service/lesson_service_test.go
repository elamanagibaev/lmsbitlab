package service_test

import (
	"errors"
	"testing"

	"LMSBitLab/internal/apperrors"
	"LMSBitLab/internal/dto"
	"LMSBitLab/internal/model"
	"LMSBitLab/internal/repository/mocks"
	"LMSBitLab/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLessonService_Create_Success(t *testing.T) {
	repoMock := mocks.NewMockLessonRepository(t)

	repoMock.EXPECT().
		Create(mock.Anything).
		RunAndReturn(func(lesson *model.Lesson) error {
			lesson.ID = 1
			return nil
		}).
		Once()

	lessonService := service.NewLessonService(repoMock)

	input := dto.CreateLessonDTO{
		Name:      "If-else Statement",
		Content:   "текст урока",
		Order:     1,
		ChapterID: 1,
	}

	result, err := lessonService.Create(input)

	require.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "If-else Statement", result.Name)
	assert.Equal(t, uint(1), result.ChapterID)
}

func TestLessonService_Create_RepositoryError(t *testing.T) {
	repoMock := mocks.NewMockLessonRepository(t)

	expectedErr := errors.New("database connection failed")

	repoMock.EXPECT().
		Create(mock.Anything).
		Return(expectedErr).
		Once()

	lessonService := service.NewLessonService(repoMock)

	input := dto.CreateLessonDTO{
		Name:      "If-else Statement",
		Order:     1,
		ChapterID: 1,
	}

	result, err := lessonService.Create(input)

	require.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, dto.LessonResponseDTO{}, result)
}

func TestLessonService_GetByID_Success(t *testing.T) {
	repoMock := mocks.NewMockLessonRepository(t)

	expectedLesson := &model.Lesson{
		ID:        1,
		Name:      "If-else Statement",
		Order:     1,
		ChapterID: 1,
	}

	repoMock.EXPECT().
		GetByID(uint(1)).
		Return(expectedLesson, nil).
		Once()

	lessonService := service.NewLessonService(repoMock)

	result, err := lessonService.GetByID(1)

	require.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "If-else Statement", result.Name)
}

func TestLessonService_GetByID_NotFound(t *testing.T) {
	repoMock := mocks.NewMockLessonRepository(t)

	repoMock.EXPECT().
		GetByID(uint(999)).
		Return(nil, apperrors.ErrLessonNotFound).
		Once()

	lessonService := service.NewLessonService(repoMock)

	result, err := lessonService.GetByID(999)

	require.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrLessonNotFound))
	assert.Equal(t, dto.LessonResponseDTO{}, result)
}

func TestLessonService_GetByChapterID_Success(t *testing.T) {
	repoMock := mocks.NewMockLessonRepository(t)

	expectedLessons := []model.Lesson{
		{ID: 1, Name: "If-else Statement", Order: 1, ChapterID: 1},
		{ID: 2, Name: "Switch Statement", Order: 2, ChapterID: 1},
	}

	repoMock.EXPECT().
		GetAllByChapterID(uint(1)).
		Return(expectedLessons, nil).
		Once()

	lessonService := service.NewLessonService(repoMock)

	result, err := lessonService.GetByChapterID(1)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "If-else Statement", result[0].Name)
}

func TestLessonService_GetByChapterID_RepositoryError(t *testing.T) {
	repoMock := mocks.NewMockLessonRepository(t)

	expectedErr := errors.New("database connection failed")

	repoMock.EXPECT().
		GetAllByChapterID(uint(1)).
		Return(nil, expectedErr).
		Once()

	lessonService := service.NewLessonService(repoMock)

	result, err := lessonService.GetByChapterID(1)

	require.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, result)
}

func TestLessonService_Update_Success(t *testing.T) {
	repoMock := mocks.NewMockLessonRepository(t)

	existingLesson := &model.Lesson{
		ID:        1,
		Name:      "Old Name",
		Order:     1,
		ChapterID: 1,
	}

	repoMock.EXPECT().
		GetByID(uint(1)).
		Return(existingLesson, nil).
		Once()

	repoMock.EXPECT().
		Update(mock.Anything).
		Return(nil).
		Once()

	lessonService := service.NewLessonService(repoMock)

	input := dto.UpdateLessonDTO{
		Name:    "New Name",
		Content: "новый текст",
		Order:   2,
	}

	result, err := lessonService.Update(1, input)

	require.NoError(t, err)
	assert.Equal(t, "New Name", result.Name)
	assert.Equal(t, 2, result.Order)
}

func TestLessonService_Update_NotFound(t *testing.T) {
	repoMock := mocks.NewMockLessonRepository(t)

	repoMock.EXPECT().
		GetByID(uint(999)).
		Return(nil, apperrors.ErrLessonNotFound).
		Once()

	lessonService := service.NewLessonService(repoMock)

	input := dto.UpdateLessonDTO{
		Name:  "New Name",
		Order: 2,
	}

	result, err := lessonService.Update(999, input)

	require.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrLessonNotFound))
	assert.Equal(t, dto.LessonResponseDTO{}, result)
}

func TestLessonService_Delete_Success(t *testing.T) {
	repoMock := mocks.NewMockLessonRepository(t)

	repoMock.EXPECT().
		Delete(uint(1)).
		Return(nil).
		Once()

	lessonService := service.NewLessonService(repoMock)

	err := lessonService.Delete(1)

	require.NoError(t, err)
}

func TestLessonService_Delete_RepositoryError(t *testing.T) {
	repoMock := mocks.NewMockLessonRepository(t)

	expectedErr := errors.New("database connection failed")

	repoMock.EXPECT().
		Delete(uint(1)).
		Return(expectedErr).
		Once()

	lessonService := service.NewLessonService(repoMock)

	err := lessonService.Delete(1)

	require.Error(t, err)
	assert.Equal(t, expectedErr, err)
}
