package models

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    uint8  `json:"priority"`
	ProjectId   int    `json:"projectId"`
	Completed   bool   `json:"completed"`
}
