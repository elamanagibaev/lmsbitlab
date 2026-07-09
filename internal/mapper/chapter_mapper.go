package mapper

import (
	"LMSBitLab/internal/dto"
	"LMSBitLab/internal/model"
)

func ToChapterResponseDTO(chapter *model.Chapter) dto.ChapterResponseDTO {
	return dto.ChapterResponseDTO{
		ID:          chapter.ID,
		Name:        chapter.Name,
		Description: chapter.Description,
		Order:       chapter.Order,
		CourseID:    chapter.CourseID,
	}
}

func ToChapterResponseDTOList(chapters []model.Chapter) []dto.ChapterResponseDTO {
	result := make([]dto.ChapterResponseDTO, 0, len(chapters))
	for _, chapter := range chapters {
		result = append(result, ToChapterResponseDTO(&chapter))
	}
	return result
}
