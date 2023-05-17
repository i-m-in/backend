package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initRoute() *mux.Router {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	//r.Path("/tasks/{user_id}").Queries("page").HandlerFunc(getTasks).Methods(http.MethodGet)
	//r.Path("/tasks/{user_id}").Queries("date").HandlerFunc(getTasks).Methods(http.MethodGet)
	r.Path("/tasks/{user_id}").HandlerFunc(getTasks).Methods(http.MethodGet)

	r.HandleFunc("/users/reg", registerUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{user_id}", getUser).Methods(http.MethodGet)
	r.HandleFunc("/users/{user_id}/friends", getFriends).Methods(http.MethodGet)

	return r
}
