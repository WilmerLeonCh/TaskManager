package tasks

import "gorm.io/gorm"

type MTask struct {
	gorm.Model
	ID          int `gorm:"primaryKey,autoIncrement"`
	Name        string
	Description string
	Completed   bool
}
