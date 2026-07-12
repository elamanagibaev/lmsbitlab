package service_test

import (
	"LMSBitLab/internal/apperrors"
	"errors"
	"testing"

	"LMSBitLab/internal/dto"
	"LMSBitLab/internal/model"
	"LMSBitLab/internal/repository/mocks"
	"LMSBitLab/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCourseService_Create_Success(t *testing.T) {
	repoMock := mocks.NewMockCourseRepository(t)

	repoMock.EXPECT().
		Create(mock.Anything).
		RunAndReturn(func(course *model.Course) error {
			course.ID = 1
			return nil
		}).
		Once()

	courseService := service.NewCourseService(repoMock)

	input := dto.CreateCourseDTO{
		Name:        "Go Developer",
		Description: "Курс по Go",
	}

	result, err := courseService.Create(input)

	require.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "Go Developer", result.Name)
	assert.Equal(t, "Курс по Go", result.Description)
}

func TestCourseService_Create_RepositoryError(t *testing.T) {
	repoMock := mocks.NewMockCourseRepository(t)

	expectedErr := errors.New("database connection failed")

	repoMock.EXPECT().
		Create(mock.Anything).
		Return(expectedErr).
		Once()

	courseService := service.NewCourseService(repoMock)

	input := dto.CreateCourseDTO{
		Name:        "Go Developer",
		Description: "Курс по Go",
	}

	result, err := courseService.Create(input)

	require.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Equal(t, dto.CourseResponseDTO{}, result)
}

func TestCourseService_GetByID_Success(t *testing.T) {
	repoMock := mocks.NewMockCourseRepository(t)

	expectedCourse := &model.Course{
		ID:          1,
		Name:        "Go Developer",
		Description: "Курс по Go",
	}

	repoMock.EXPECT().
		GetByID(uint(1)).
		Return(expectedCourse, nil).
		Once()

	courseService := service.NewCourseService(repoMock)

	result, err := courseService.GetByID(1)

	require.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, "Go Developer", result.Name)
}

func TestCourseService_GetByID_NotFound(t *testing.T) {
	repoMock := mocks.NewMockCourseRepository(t)

	repoMock.EXPECT().
		GetByID(uint(999)).
		Return(nil, apperrors.ErrCourseNotFound).
		Once()

	courseService := service.NewCourseService(repoMock)

	result, err := courseService.GetByID(999)

	require.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrCourseNotFound))
	assert.Equal(t, dto.CourseResponseDTO{}, result)
}

func TestCourseService_GetAll_Success(t *testing.T) {
	repoMock := mocks.NewMockCourseRepository(t)

	expectedCourses := []model.Course{
		{ID: 1, Name: "Go Developer", Description: "Курс по Go"},
		{ID: 2, Name: "Python Developer", Description: "Курс по Python"},
	}

	repoMock.EXPECT().
		GetAll().
		Return(expectedCourses, nil).
		Once()

	courseService := service.NewCourseService(repoMock)

	result, err := courseService.GetAll()

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Go Developer", result[0].Name)
	assert.Equal(t, "Python Developer", result[1].Name)
}

func TestCourseService_GetAll_RepositoryError(t *testing.T) {
	repoMock := mocks.NewMockCourseRepository(t)

	expectedErr := errors.New("database connection failed")

	repoMock.EXPECT().
		GetAll().
		Return(nil, expectedErr).
		Once()

	courseService := service.NewCourseService(repoMock)

	result, err := courseService.GetAll()

	require.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, result)
}

func TestCourseService_Update_Success(t *testing.T) {
	repoMock := mocks.NewMockCourseRepository(t)

	existingCourse := &model.Course{
		ID:          1,
		Name:        "Old Name",
		Description: "Old Description",
	}

	repoMock.EXPECT().
		GetByID(uint(1)).
		Return(existingCourse, nil).
		Once()

	repoMock.EXPECT().
		Update(mock.Anything).
		Return(nil).
		Once()

	courseService := service.NewCourseService(repoMock)

	input := dto.UpdateCourseDTO{
		Name:        "New Name",
		Description: "New Description",
	}

	result, err := courseService.Update(1, input)

	require.NoError(t, err)
	assert.Equal(t, "New Name", result.Name)
	assert.Equal(t, "New Description", result.Description)
}

func TestCourseService_Update_NotFound(t *testing.T) {
	repoMock := mocks.NewMockCourseRepository(t)

	repoMock.EXPECT().
		GetByID(uint(999)).
		Return(nil, apperrors.ErrCourseNotFound).
		Once()

	courseService := service.NewCourseService(repoMock)

	input := dto.UpdateCourseDTO{
		Name:        "New Name",
		Description: "New Description",
	}

	result, err := courseService.Update(999, input)

	require.Error(t, err)
	assert.True(t, errors.Is(err, apperrors.ErrCourseNotFound))
	assert.Equal(t, dto.CourseResponseDTO{}, result)
}

func TestCourseService_Delete_Success(t *testing.T) {
	repoMock := mocks.NewMockCourseRepository(t)

	repoMock.EXPECT().
		Delete(uint(1)).
		Return(nil).
		Once()

	courseService := service.NewCourseService(repoMock)

	err := courseService.Delete(1)

	require.NoError(t, err)
}

func TestCourseService_Delete_RepositoryError(t *testing.T) {
	repoMock := mocks.NewMockCourseRepository(t)

	expectedErr := errors.New("database connection failed")

	repoMock.EXPECT().
		Delete(uint(1)).
		Return(expectedErr).
		Once()

	courseService := service.NewCourseService(repoMock)

	err := courseService.Delete(1)

	require.Error(t, err)
	assert.Equal(t, expectedErr, err)
}
