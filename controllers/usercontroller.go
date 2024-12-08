package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-golang/helpers"
	"todo-golang/models"
	"todo-golang/usecases/user"
)
type UserCreateInput struct {
	Name string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=8"`
}

type UserUpdateInput struct {
	Name string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
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
	var userInput UserCreateInput
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

func UpdateUser (w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	var userInput UserUpdateInput
	json.NewDecoder(r.Body).Decode(&userInput)
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")
	if helpers.ValidateRequest(userInput, w) {
		return
	}
	err := user.Update(id, userInput.Name, userInput.Email)
	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Description})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": "user successfuly updated!"})
	}
}

func DeleteUser (w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	err := user.Delete(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Description})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": "user successfuly removed!"})
	}
}