package main

import (
	"log"
	"net/http"
	"todo-golang/database"
	"todo-golang/database/migrations"
	"todo-golang/routes"
)

func main() {
	database.InitDB()
	migrations.Migrate()
	router := http.NewServeMux()
	routes.StartRoutes(router)
	log.Fatal(http.ListenAndServe(":9000", router))
}