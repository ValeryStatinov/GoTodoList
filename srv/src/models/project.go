package models

type Project struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	UserId int    `json:"userId,omitempty"`
}

type ProjectWithTasksCount struct {
	TasksCount int `json:"tasksCount"`
	Project
}

func (p *Project) Validate() bool {
	return p.Name != ""
}
