package mapper

import (
	"LMSBitLab/internal/dto"
	"LMSBitLab/internal/model"
)

func ToCourseResponseDTO(course *model.Course) dto.CourseResponseDTO {
	return dto.CourseResponseDTO{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
	}
}

func ToCourseResponseDTOList(courses []model.Course) []dto.CourseResponseDTO {
	result := make([]dto.CourseResponseDTO, 0, len(courses))
	for _, course := range courses {
		result = append(result, ToCourseResponseDTO(&course))
	}
	return result
}
