package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal("connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE users SET age=$2 WHERE id=$1")
	if err != nil {
		log.Fatal("can't prepare query all users statment", err)
	}

	updateId := 2
	newAge := 22
	if _, err := stmt.Exec(updateId, newAge); err != nil {
		log.Fatal("can't update user", err)
	}

	log.Println("updated")
}
