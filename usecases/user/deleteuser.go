package user

import (
	"todo-golang/helpers"
	"todo-golang/models"
)

func Delete(id int) *helpers.CustomError {
	_, err := models.FindOneById(id)
	if err != nil {
		return err
	}
	return models.DeleteById(id)
}