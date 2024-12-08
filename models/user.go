package models

import (
	"todo-golang/database"
	"todo-golang/helpers"

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
	var users []map[string]interface{}
	result := database.Connection.Model(&User{}).Select("id", "name", "email").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(users) == 0 {
		return []map[string]interface{}{}, nil
	}
	return users, nil
}

func FindOneById(id int) (*map[string]interface{}, *helpers.CustomError) {
	var user map[string]interface{}
	result := database.Connection.Model(&User{}).Select("id", "name", "email").First(&user, id)
	if result.Error != nil {
		return nil, helpers.NewError(result.Error.Error(), 404)
	}
	return &user, nil
}

func DeleteById(id int) *helpers.CustomError {
	result := database.Connection.Delete(&User{}, id)
	if result.Error != nil {
		return helpers.NewError(result.Error.Error(), 500)
	}
	return nil
}