package models

type User struct {
	Id       int
	Name     string
	Password string
}

func (u *User) HaveAccesToProject(pr *Project) bool {
	return pr.UserId == u.Id
}
