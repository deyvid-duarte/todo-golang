package controllers

import (
	"io"
	"net/http"
)

func ListTasks (w http.ResponseWriter, _ *http.Request)  {
	io.WriteString(w, "Listando todas as tarefa!")
}

func FindOneTask (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Listando uma tarefa!")
}

func CreateTask (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Criando uma tarefa!")
}

func UpdateTask (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Atualizando uma tarefa!")
}

func DeleteTask (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Atualizando uma tarefa!")
}