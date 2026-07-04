package model

import "time"

type Lesson struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Content     string `gorm:"type:text"`
	Order       int    `gorm:"not null"`
	ChapterID   uint   `gorm:"not null;index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Chapter Chapter `gorm:"foreignKey:ChapterID;references:ID"`
}

func (l Lesson) TableName() string {
	return "lessons"
}
