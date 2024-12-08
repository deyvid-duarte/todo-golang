package user

import (
	"todo-golang/helpers"
	"todo-golang/models"
)

func Update(id int, name string, email string) *helpers.CustomError {
	_, err := models.FindOneById(id)
	if err != nil {
		return err
	}
	user := models.User{Name: name, Email: email}
	if user.EmailAlreadyExists(&id) {
		return helpers.NewError("email is already in use", 422)
	}
	errUpdate := models.UpdateById(id, map[string]interface{}{"name": name, "email": email})
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}