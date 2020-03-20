package store

import (
	"database/sql"
	"todolist/src/models"
	"todolist/src/systemlogger"
)

type UsersManager struct {
	db *sql.DB
}

func (um *UsersManager) GetByName(name string) *models.User {
	user := &models.User{}

	row := um.db.QueryRow("select * from users where name=$1", name)
	err := row.Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		systemlogger.Log(err.Error())
	}

	return user
}
