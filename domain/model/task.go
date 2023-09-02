package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string `json:"name" gorm:"varchar(30);not null"`
	Description string `json:"description" gorm:"varchar(400);not null"`
}

type CreateTask struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
