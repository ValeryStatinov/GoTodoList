package database

import (
	"database/sql"
)

func PrepareSQL(query string) *sql.Stmt {
	db := GetDBInstance()
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}

	return stmt
}
