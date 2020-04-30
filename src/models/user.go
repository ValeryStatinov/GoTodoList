package models

import (
	"bytes"
	"crypto/sha256"
)

type User struct {
	Id       int
	Name     string
	Password []byte
}

func (u *User) HaveAccesToProject(pr *Project) bool {
	return pr.UserId == u.Id
}

func CheckPassword(user *User, login string, password string) bool {
	h := sha256.New()
	_, _ = h.Write([]byte(password))
	hashPassword := h.Sum(nil)
	return bytes.Equal(hashPassword[:], user.Password[:])
}
