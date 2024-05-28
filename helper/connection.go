package helper

import (
	"database/sql"
	"time"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_rest_study")

	PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
