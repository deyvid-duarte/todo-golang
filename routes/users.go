package routes

import (
	"net/http"
	"todo-golang/controllers"
)


func HandleRouteUser(router *http.ServeMux) {
	router.HandleFunc("GET /users", controllers.ListUsers)
	router.HandleFunc("GET /users/{id}", controllers.FindOneUser)
	router.HandleFunc("POST /users", controllers.CreateUser)
	router.HandleFunc("PUT /users/{id}", controllers.UpdateUser)
	router.HandleFunc("DELETE /users/{id}", controllers.DeleteUser)
}