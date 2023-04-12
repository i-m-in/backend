package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initRoute() *mux.Router {
	r := mux.NewRouter() // localhost:8080

	tasks := r.PathPrefix("/tasks").Subrouter()                                // localhost:8080/tasks
	tasks.HandleFunc("", GetAllTasks).Methods(http.MethodGet)                  // localhost:8080/tasks/
	tasks.HandleFunc("{taskId:-?[0-9]+}", GetTaskById).Methods(http.MethodGet) // localhost:8080/tasks/1, localhost:8080/tasks/42, localhost:8080/tasks/0, localhost:8080/tasks/-42

	return r
}
