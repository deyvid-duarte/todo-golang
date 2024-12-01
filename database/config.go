package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func InitDB() {
	var err error
	Connection, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("Erro ao abrir a conexão com o banco de dados. Erro: " + err.Error())
	}
}
