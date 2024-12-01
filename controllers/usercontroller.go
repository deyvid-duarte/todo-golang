package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"todo-golang/models"
)

func ListUsers (w http.ResponseWriter, _ *http.Request)  {
	io.WriteString(w, "Listando todos os usuários!")
}

func FindOneUser (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Listando um usuário!")
}

func CreateUser (w http.ResponseWriter, _ *http.Request) {
	user := models.User{Name: "Deyvid Duarte", Email: "deyviddc@gmail.com", Password: "abc@123"}
	w.Header().Add("Content-Type", "application/json")
	if user.Save() {
		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Não foi possivel criar o usuário!"})
	}
}

func UpdateUser (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Atualizando um usuário!")
}

func DeleteUser (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Atualizando um usuário!")
}