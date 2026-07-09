package dto

type CreateLessonDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Content     string `json:"content" binding:"required"`
	Order       int    `json:"order" binding:"required"`
	ChapterID   uint   `json:"chapter_id" binding:"required"`
}

type UpdateLessonDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Content     string `json:"content" binding:"required"`
	Order       int    `json:"order" binding:"required"`
}

type LessonResponseDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Order       int    `json:"order"`
	ChapterID   uint   `json:"chapter_id"`
}
