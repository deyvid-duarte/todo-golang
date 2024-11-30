package main

import (
	"log"
	"net/http"
	"todo-golang/routes"
)

func main() {
	router := http.NewServeMux()
	routes.StartRoutes(router)
	log.Fatal(http.ListenAndServe(":9000", router))
}