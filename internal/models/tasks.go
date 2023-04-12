package models

import "fmt"

type Task struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

func (model *Task) GetTasksService() ([]Task, error) {
	var result []Task
	err := db.Select(&result, "SELECT * FROM tasks")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil
}

func (model *Task) GetTaskByIdService() error {
	// model := Task{id: 5, title:"", description:""}

}
