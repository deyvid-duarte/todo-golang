package models

import (
	"todo-golang/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Password string
}

func (u *User) Save() bool {
	result := database.Connection.Create(u)
	return result.RowsAffected > 0
}