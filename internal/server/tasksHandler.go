package server

import (
	"encoding/json"
	"fmt"
	"github.com/melnk300/CalendarBack/internal/models"
	"net/http"
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
