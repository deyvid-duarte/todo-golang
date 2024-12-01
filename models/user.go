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

func (u *User) EmailAlreadyExists() bool {
	var count int64
	database.Connection.Model(u).Where("email = ?", u.Email).Count(&count)
	return count > 0
}

func GetAllUsers() ([]map[string]interface{}, error) {
	var user User
	var users []map[string]interface{}
	result := database.Connection.Model(&user).Select("id", "name", "email").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(users) == 0 {
		return []map[string]interface{}{}, nil
	}
	return users, nil
}