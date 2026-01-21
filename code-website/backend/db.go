package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "appuser:app123@tcp(127.0.0.1:3306)/employee_db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
