package dto

type CreateCourseDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateCourseDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type CourseResponseDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
