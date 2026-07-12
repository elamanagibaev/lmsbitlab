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

func TestChapterService_Create_Success(t *testing.T) {
	repoMock := mocks.NewMockChapterRepository(t)

	repoMock.EXPECT().
		Create(mock.Anything).
		RunAndReturn(func(chapter *model.Chapter) error {
			chapter.ID = 1
			return nil
		}).
		Once()

	chapterService := service.NewChapterService(repoMock)

	input := dto.CreateChapterDTO{
		Name:        "Control Structures",
		Description: "Управляющие конструкции",
		Order:       1,
		CourseID:    1,
	}

	result, err := chapterService.Create(input)

	require.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "Control Structures", result.Name)
	assert.Equal(t, uint(1), result.CourseID)
}

func TestChapterService_Create_RepositoryError(t *testing.T) {
	repoMock := mocks.NewMockChapterRepository(t)

	expectedErr := errors.New("database connection failed")

	repoMock.EXPECT().
		Create(mock.Anything).
		Return(expectedErr).
		Once()

	chapterService := service.NewChapterService(repoMock)

	input := dto.CreateChapterDTO{
		Name:     "Control Structures",
		Order:    1,
		CourseID: 1,
	}

	result, err := chapterService.Create(input)

	require.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, dto.ChapterResponseDTO{}, result)
}

func TestChapterService_GetByID_Success(t *testing.T) {
	repoMock := mocks.NewMockChapterRepository(t)

	expectedChapter := &model.Chapter{
		ID:       1,
		Name:     "Control Structures",
		Order:    1,
		CourseID: 1,
	}

	repoMock.EXPECT().
		GetByID(uint(1)).
		Return(expectedChapter, nil).
		Once()

	chapterService := service.NewChapterService(repoMock)

	result, err := chapterService.GetByID(1)

	require.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "Control Structures", result.Name)
}

func TestChapterService_GetByID_NotFound(t *testing.T) {
	repoMock := mocks.NewMockChapterRepository(t)

	repoMock.EXPECT().
		GetByID(uint(999)).
		Return(nil, apperrors.ErrChapterNotFound).
		Once()

	chapterService := service.NewChapterService(repoMock)

	result, err := chapterService.GetByID(999)

	require.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrChapterNotFound))
	assert.Equal(t, dto.ChapterResponseDTO{}, result)
}

func TestChapterService_GetByCourseID_Success(t *testing.T) {
	repoMock := mocks.NewMockChapterRepository(t)

	expectedChapters := []model.Chapter{
		{ID: 1, Name: "Control Structures", Order: 1, CourseID: 1},
		{ID: 2, Name: "Data Types", Order: 2, CourseID: 1},
	}

	repoMock.EXPECT().
		GetAllByCourseID(uint(1)).
		Return(expectedChapters, nil).
		Once()

	chapterService := service.NewChapterService(repoMock)

	result, err := chapterService.GetByCourseID(1)

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Control Structures", result[0].Name)
}

func TestChapterService_GetByCourseID_RepositoryError(t *testing.T) {
	repoMock := mocks.NewMockChapterRepository(t)

	expectedErr := errors.New("database connection failed")

	repoMock.EXPECT().
		GetAllByCourseID(uint(1)).
		Return(nil, expectedErr).
		Once()

	chapterService := service.NewChapterService(repoMock)

	result, err := chapterService.GetByCourseID(1)

	require.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, result)
}

func TestChapterService_Update_Success(t *testing.T) {
	repoMock := mocks.NewMockChapterRepository(t)

	existingChapter := &model.Chapter{
		ID:       1,
		Name:     "Old Name",
		Order:    1,
		CourseID: 1,
	}

	repoMock.EXPECT().
		GetByID(uint(1)).
		Return(existingChapter, nil).
		Once()

	repoMock.EXPECT().
		Update(mock.Anything).
		Return(nil).
		Once()

	chapterService := service.NewChapterService(repoMock)

	input := dto.UpdateChapterDTO{
		Name:  "New Name",
		Order: 2,
	}

	result, err := chapterService.Update(1, input)

	require.NoError(t, err)
	assert.Equal(t, "New Name", result.Name)
	assert.Equal(t, 2, result.Order)
}

func TestChapterService_Update_NotFound(t *testing.T) {
	repoMock := mocks.NewMockChapterRepository(t)

	repoMock.EXPECT().
		GetByID(uint(999)).
		Return(nil, apperrors.ErrChapterNotFound).
		Once()

	chapterService := service.NewChapterService(repoMock)

	input := dto.UpdateChapterDTO{
		Name:  "New Name",
		Order: 2,
	}

	result, err := chapterService.Update(999, input)

	require.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrChapterNotFound))
	assert.Equal(t, dto.ChapterResponseDTO{}, result)
}

func TestChapterService_Delete_Success(t *testing.T) {
	repoMock := mocks.NewMockChapterRepository(t)

	repoMock.EXPECT().
		Delete(uint(1)).
		Return(nil).
		Once()

	chapterService := service.NewChapterService(repoMock)

	err := chapterService.Delete(1)

	require.NoError(t, err)
}

func TestChapterService_Delete_RepositoryError(t *testing.T) {
	repoMock := mocks.NewMockChapterRepository(t)

	expectedErr := errors.New("database connection failed")

	repoMock.EXPECT().
		Delete(uint(1)).
		Return(expectedErr).
		Once()

	chapterService := service.NewChapterService(repoMock)

	err := chapterService.Delete(1)

	require.Error(t, err)
	assert.Equal(t, expectedErr, err)
}
