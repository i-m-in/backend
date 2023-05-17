package models

type Project struct {
	ID           int    `json:"project_id" db:"project_id"`
	Title        string `json:"project_title" db:"project_title"`
	MaintainerID int    `json:"maintainer_id" db:"maintainer_id"`
	Color        string `json:"color" db:"color"`
	Description  string `json:"description" db:"description"`
}
