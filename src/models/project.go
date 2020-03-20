package models

type Project struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	UserId string `json:"userId,omitempty"`
}
