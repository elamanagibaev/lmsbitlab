package apperrors

import (
	"errors"
	"net/http"
)

var (
	ErrCourseNotFound  = errors.New("course not found")
	ErrChapterNotFound = errors.New("chapter not found")
	ErrLessonNotFound  = errors.New("lesson not found")
)

func StatusCode(err error) int {
	switch {
	case errors.Is(err, ErrCourseNotFound),
		errors.Is(err, ErrChapterNotFound),
		errors.Is(err, ErrLessonNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
