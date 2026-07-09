package mapper

import (
	"LMSBitLab/internal/dto"
	"LMSBitLab/internal/model"
)

func ToLessonResponseDTO(lesson *model.Lesson) dto.LessonResponseDTO {
	return dto.LessonResponseDTO{
		ID:          lesson.ID,
		Name:        lesson.Name,
		Description: lesson.Description,
		Content:     lesson.Content,
		Order:       lesson.Order,
		ChapterID:   lesson.ChapterID,
	}
}

func ToLessonResponseDTOList(lessons []model.Lesson) []dto.LessonResponseDTO {
	result := make([]dto.LessonResponseDTO, 0, len(lessons))
	for _, lesson := range lessons {
		result = append(result, ToLessonResponseDTO(&lesson))
	}
	return result
}
