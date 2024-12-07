package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"todo-golang/helpers"
	"todo-golang/models"
	"todo-golang/usecases/user"
)
type UserInput struct {
	Name string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=8"`
}

func ListUsers (w http.ResponseWriter, _ *http.Request)  {
	users, err := models.GetAllUsers()
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

func FindOneUser (w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	user, err := models.FindOneById(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Description})
	} else {
		json.NewEncoder(w).Encode(user)
	}
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