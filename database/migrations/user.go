package migrations

import (
	"todo-golang/database"
	"todo-golang/models"
)

func Migrate() {
	database.Connection.AutoMigrate(&models.User{})
}