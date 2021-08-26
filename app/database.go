package app

import (
	"database/sql"
	"fmt"
	"go-restful-api/helper"
	"time"
)

func NewDB(DBPort, DBHost, DBUser, DBPassword, DBName string) *sql.DB {
	// urlDatabase := "postgres://mapple@localhost:5432/go_rest?sslmode=disable"
	psqlconn := fmt.Sprintf(`dbname=%s user=%s password=%s host=%s port=%s  sslmode=disable`, DBName, DBUser, DBPassword, DBHost, DBPort)
	db, err := sql.Open("postgres", psqlconn)
	err = db.Ping()
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
