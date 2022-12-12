package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// url := "postgres://wvbxvjvp:8iC9SoaN_pUBkXWIcVw5ODwtVl6t9unQ@tiny.db.elephantsql.com/wvbxvjvp"
	url := os.Getenv("DATABASE_URL")
	// DATABASE_URL=postgres://wvbxvjvp:8iC9SoaN_pUBkXWIcVw5ODwtVl6t9unQ@tiny.db.elephantsql.com/wvbxvjvp go run insert.go

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("connect to database error", err)
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO users (name, age) values ($1, $2) RETURNING id", "Ewka", 18)
	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("can not insert data", err)
	}
	log.Println("insert todo success id:", id)
}
