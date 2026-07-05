package model

import "time"

type Chapter struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Order       int    `gorm:"not null"`
	CourseID    uint   `gorm:"not null;index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Course  Course   `gorm:"foreignKey:CourseID;references:ID"`
	Lessons []Lesson `gorm:"foreignKey:ChapterID"`
}
