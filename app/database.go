package app

import (
	"database/sql"
	"fmt"
	"go-restful-api/helper"
	"time"
)

func NewDB(DBPort, DBHost, DBUser, DBPassword, DBName string) *sql.DB {
	// urlDatabase := "postgres://pqgotest@localhost:5432/go_rest?sslmode=disable"
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DBHost, DBPort, DBUser, DBPassword, DBName)
	db, err := sql.Open("postgres", psqlconn)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
