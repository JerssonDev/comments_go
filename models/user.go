package models

import (
	"github.com/jinzhu/gorm"
)

//User ...
type User struct {
	gorm.Model
	Username        string    `json:"username" gorm:"not null;unique"`
	Email           string    `json:"email" gorm:"not null;unique"`
	Fullname        string    `json:"fullname" gorm:"not null"`
	Password        string    `json:"password,omitempty" gorm:"not null;type:varchar(256)"`
	ConfirmPassword string    `json:"confirmPassword,omitempty" gorm:"-"` // (-) le dice a gorm que lo omita y no lo cree ni nada por el estilo
	Pickture        string    `json:"picture"`
	Comments        []Comment `json:"comments,omitempty"`
}
