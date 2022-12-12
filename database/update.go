package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	url := "postgres://wvbxvjvp:8iC9SoaN_pUBkXWIcVw5ODwtVl6t9unQ@tiny.db.elephantsql.com/wvbxvjvp"

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
		log.Fatal("can't query all users", err)
	}

	log.Println("updated")
}
