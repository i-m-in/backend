package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/melnk300/CalendarBack/internal/models"
	"net/http"
	"strconv"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var task models.Task = models.Task{}
	tasks, err := task.GetTasksService()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(tasks)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	task := models.Task{}
	task.Id, _ = strconv.Atoi(mux.Vars(r)["taskId"]) // 5
	if task.Id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := task.GetTaskByIdService()
	if err != nil {
		return
	}
}
