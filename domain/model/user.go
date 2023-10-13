package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"varchar(100);not null"`
	Email string `json:"email" gorm:"varchar(320);not null;uniqueIndex"`
	Tasks []Task `json:"tasks"`
}

type CreateUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
