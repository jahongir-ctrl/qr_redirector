package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func InitDB() *sql.DB {

	connStr := "host = localhost user=postgres password=jhg_3399 dbname=scan_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД", err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("БД недоступна:", err)
	}

	query := `
CREATE TABLE IF NOT EXISTS scans (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT now(),
		user_agent TEXT
);`

	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("Ошибка при создании таблицы: ", err)
	}

	return db
}
