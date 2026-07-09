package dto

type CreateChapterDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Order       int    `json:"order" binding:"required"`
	CourseID    uint   `json:"course_id" binding:"required"`
}

type UpdateChapterDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Order       int    `json:"order" binding:"required"`
}

type ChapterResponseDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	CourseID    uint   `json:"course_id"`
}
