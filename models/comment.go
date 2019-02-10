package models

import (
	"github.com/jinzhu/gorm"
)

// Comment ... comentario de sistema
type Comment struct {
	gorm.Model
	UserID   uint      `json:"userID"`
	ParentID uint      `json:"parentID"`
	Votes    int32     `json:"votes"`
	Content  string    `json:"content"`
	HasVote  int8      `json:"hasVote" gorm:"-"`
	User     []User    `json:"user,omitempty" gorm:"-"`
	Childen  []Comment `json:"children,omitempty"`
}
