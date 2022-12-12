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

	createTb := `CREATE TABLE IF NOT EXISTS users ( id SERIAL PRIMARY KEY, name TEXT, age INT );`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}
	log.Println("create table success")
}
