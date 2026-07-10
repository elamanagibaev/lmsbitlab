package apperrors

import "errors"

var ErrCourseNotFound = errors.New("course not found")
var ErrChapterNotFound = errors.New("chapter not found")
var ErrLessonNotFound = errors.New("lesson not found")
