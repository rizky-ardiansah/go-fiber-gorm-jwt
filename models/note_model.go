package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title   string `json:"title" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
	UserID  uint   `json:"user_id" gorm:"not null"`
	User    User   `json:"-" gorm:"foreignKey:UserID"`
}
