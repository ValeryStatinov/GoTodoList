package models

type Project struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	UserId int    `json:"userId,omitempty"`
}

type ProjectRequest struct {
	Name string `json:"name"`
}
