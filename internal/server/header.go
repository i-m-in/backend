package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initRoute() *mux.Router {
	r := mux.NewRouter()

	tasks := r.PathPrefix("/tasks").Subrouter()
	tasks.HandleFunc("", GetAllTasks).Methods(http.MethodGet)

	return r
}
