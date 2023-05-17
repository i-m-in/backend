package models

type TaskToUsers struct {
	ID     int    `json:"id" db:"id"`
	TaskID int    `json:"-" db:"task_id_ttu"`
	UserID int    `json:"-" db:"user_id_ttu"`
	Role   string `json:"role" db:"role"`
}

type TUP struct {
	Task
	User
	Project
	TaskToUsers
}

func (tu *TaskToUsers) Create() error {
	_, err := db.Exec("INSERT INTO tasks_users (task_id, user_id, role) VALUES ($1, $2, $3)",
		tu.TaskID, tu.UserID, tu.Role)
	if err != nil {
		return err
	}
	return nil
}

func (tu *TaskToUsers) Update() error {
	_, err := db.Exec("UPDATE tasks_users SET task_id=$1, user_id=$2, role=$3 WHERE id=$4",
		tu.TaskID, tu.UserID, tu.Role, tu.ID)
	if err != nil {
		return err
	}
	return nil
}

func (tu *TaskToUsers) Delete() error {
	_, err := db.Exec("DELETE FROM tasks_users WHERE id=$1", tu.ID)
	if err != nil {
		return err
	}
	return nil
}
