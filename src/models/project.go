package models

type Project struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	UserId int    `json:"userId,omitempty"`
}

func (p *Project) Validate() bool {
	if p.Name == "" {
		return false
	}

	return true
}
