package routes

import (
	"net/http"
	"todo-golang/controllers"
)


func HandleRouteTask(router *http.ServeMux) {
	router.HandleFunc("GET /tasks", controllers.ListTasks)
	router.HandleFunc("GET /tasks/{id}", controllers.FindOneTask)
	router.HandleFunc("POST /tasks", controllers.CreateTask)
	router.HandleFunc("PUT /tasks", controllers.UpdateTask)
	router.HandleFunc("DELETE /tasks/{id}", controllers.DeleteTask)
}