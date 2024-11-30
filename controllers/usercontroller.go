package controllers

import (
	"io"
	"net/http"
)

func ListUsers (w http.ResponseWriter, _ *http.Request)  {
	io.WriteString(w, "Listando todos os usuários!")
}

func FindOneUser (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Listando um usuário!")
}

func CreateUser (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Criando um usuário!")
}

func UpdateUser (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Atualizando um usuário!")
}

func DeleteUser (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Atualizando um usuário!")
}