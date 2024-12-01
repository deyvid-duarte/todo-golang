package routes

import (
	"net/http"
	"todo-golang/controllers"
)


func HandleRouteUser(router *http.ServeMux) {
	// router.HandleFunc("GET /users", controllers.ListUsers)
	router.HandleFunc("GET /users/{id}", controllers.FindOneUser)
	router.HandleFunc("GET /users", controllers.CreateUser)
	router.HandleFunc("PUT /users", controllers.UpdateUser)
	router.HandleFunc("DELETE /users/{id}", controllers.DeleteUser)
}