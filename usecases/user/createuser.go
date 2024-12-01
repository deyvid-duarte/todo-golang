package user

import (
	"todo-golang/helpers"
	"todo-golang/models"

	"golang.org/x/crypto/bcrypt"
)


func Create(name string, email string, password string) (*models.User, *helpers.CustomError) {
	hashedPassword, err := cryptPassword(password)
	if err != nil {
		panic(err.Error())
	}
	user := models.User{Name: name, Email: email, Password: hashedPassword}
	if user.EmailAlreadyExists() {
		return nil, helpers.NewError("email is already in use", 422)
	}
	if !user.Save() {
		return nil, helpers.NewError("error to create the user", 500)
	}
	return &user, nil
}

func cryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}