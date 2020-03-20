package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *User) HaveAccesToProject(pr *Project) bool {
	return pr.UserId == u.Id
}
