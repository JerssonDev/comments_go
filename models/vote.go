package models

import (
	"github.com/jinzhu/gorm"
)

// Vote ... votos realizados por comentario
type Vote struct {
	gorm.Model
	CommentID uint `json:"commentID" gorm:"not null"`
	UserID    uint `json:"userID" gorm:"not null"`
	Value     bool `json:"value" gorm:"not null"`
}
