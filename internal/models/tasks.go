package models

import "fmt"

type Task struct {
	ID          int    `json:"task_id,omitempty" db:"task_id"`
	Title       string `json:"task_title,omitempty" db:"task_title"`
	Description string `json:"description,omitempty" db:"description"`
	ProjectID   int    `json:"project_id_t,omitempty" db:"project_id_t"`
	StartTime   string `json:"start_time,omitempty" db:"start_time"`
	EndTime     string `json:"end_time,omitempty" db:"end_time"`
	Type        int    `json:"type,omitempty" db:"type"`
	IsDone      bool   `json:"is_done,omitempty" db:"is_done"`
	AuthorID    int    `json:"task_author_id,omitempty" db:"task_author_id"`
	IsPublic    bool   `json:"is_public,omitempty" db:"is_public"`
}

type TaskExtended struct {
	Task
	Members []struct {
		FirstName string `json:"first_name" db:"first_name"`
		LastName  string `json:"last_name" db:"last_name"`
	} `json:"members"`
	Color string `json:"color" db:"color"`
}

func (t *Task) GetTasks(userId int) ([]TaskExtended, error) {
	var Tasks []TaskExtended
	err := db.Select(&Tasks, "SELECT t.*, p.color FROM tasks AS t JOIN project p on p.project_id = t.project_id_t WHERE task_author_id=$1", userId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for task := range Tasks {
		fmt.Println(Tasks[task].ID)
		err = db.Select(&Tasks[task].Members, "SELECT u.user_name, u.first_name, u.last_name FROM tasks_users JOIN users u on tasks_users.user_id_ttu = u.user_id WHERE task_id_ttu=$1", Tasks[task].ID)
		if err != nil {
			fmt.Println(err)
		}
	}

	return Tasks, nil
}
