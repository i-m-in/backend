package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/melnk300/CalendarBack/internal/models"
	"net/http"
	"strconv"
)

func getTasks(w http.ResponseWriter, r *http.Request) {
	var Task models.Task

	userId, _ := strconv.Atoi(mux.Vars(r)["user_id"])

	tasks, err := Task.GetTasks(userId)

	if err != nil {
		switch err.Error() {
		case "tasks not found":
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tasks)
}
