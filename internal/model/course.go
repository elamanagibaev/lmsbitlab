package model

import "time"

type Course struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Chapter []Chapter `gorm:"foreignKey:CourseID"`
}
