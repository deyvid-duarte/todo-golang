package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"todo-golang/helpers"
	"todo-golang/usecases/user"
)
type UserInput struct {
	Name string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=8"`
}

func ListUsers (w http.ResponseWriter, _ *http.Request)  {
	io.WriteString(w, "Listando todos os usuários!")
}

func FindOneUser (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Listando um usuário!")
}

func CreateUser (w http.ResponseWriter, r *http.Request) {
	var userInput UserInput
	json.NewDecoder(r.Body).Decode(&userInput)
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")
	if helpers.ValidateRequest(userInput, w) {
		return
	}
	userRegistered, err := user.Create(userInput.Name, userInput.Email, userInput.Password)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(userRegistered)
	} else {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Description})
	}
}

func UpdateUser (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Atualizando um usuário!")
}

func DeleteUser (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Atualizando um usuário!")
}