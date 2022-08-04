package app

import (
	"database/sql"
	"time"

	"github.com/mproyyan/go-simple-restful/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:sussyballs@tcp(localhost:3306)/go_simple_restful")
	helper.CheckErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
