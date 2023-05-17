package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/melnk300/CalendarBack/internal/models"
	"github.com/melnk300/CalendarBack/utils"
	"net/http"
	"strconv"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	var User models.User

	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	User.Password, _ = utils.HashPassword(User.Password)

	err = User.Register()
	if err != nil {
		switch err.Error() {
		case "email already exists":
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(User)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var User models.User

	User.ID, _ = strconv.Atoi(mux.Vars(r)["user_id"])

	err := User.GetUser()
	if err != nil {
		switch err.Error() {
		case "user not found":
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(User)
}

func getFriends(w http.ResponseWriter, r *http.Request) {
	var User models.User

	User.ID, _ = strconv.Atoi(mux.Vars(r)["user_id"])

	friends, err := User.GetFriends()
	if err != nil {
		switch err.Error() {
		case "user not found":
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(friends)
}
