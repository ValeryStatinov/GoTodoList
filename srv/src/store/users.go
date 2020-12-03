package store

import (
	"database/sql"
	"todolist/src/models"
	"todolist/src/systemlogger"

	"github.com/dgrijalva/jwt-go"
)

type UsersManager struct {
	db *sql.DB
}

type JWTPayload struct {
	UserId int
	jwt.StandardClaims
	Time int64
}

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (um *UsersManager) GetByName(name string) (*models.User, bool) {
	user := &models.User{}

	query := "select * from users where name=$1"

	row := um.db.QueryRow(query, name)
	err := row.Scan(&user.Id, &user.Name, &user.Password)
	ok := true
	if err != nil {
		systemlogger.Log(err.Error(), query, name)
		ok = false
	}

	return user, ok
}

func (um *UsersManager) GetById(id int) (*models.User, bool) {
	user := &models.User{}

	query := "select * from users where id=$1"

	row := um.db.QueryRow(query, id)
	err := row.Scan(&user.Id, &user.Name, &user.Password)
	ok := true
	if err != nil {
		systemlogger.Log(err.Error(), query, string(id))
		ok = false
	}

	return user, ok
}

func (um *UsersManager) Create(name string) (*models.User, bool) {
	var id int
	ok := true
	query := "insert into users (name, password) values ($1, $2) returning id"

	err := um.db.QueryRow(query, name, "pass").Scan(&id)
	if err != nil {
		systemlogger.Log(err.Error(), query, name)
		ok = false
	}

	user := &models.User{Id: id, Name: name, Password: "pass"}

	return user, ok
}
