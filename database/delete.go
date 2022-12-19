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

	stmt, err := db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		log.Fatal("can't prepare query all users statment", err)
	}

	deleteId := 3
	if _, err := stmt.Exec(deleteId); err != nil {
		log.Fatal("can't exec delete stmt", err)
	}

	log.Println("deleted")
}
