package app

import (
	"Golang-RestAPI/helper"
	"database/sql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:3306)/golang-migration")
	helper.Panic(err)

	return db
}
