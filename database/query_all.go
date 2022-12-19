package main

import (
	"database/sql"
	"fmt"
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

	stmt, err := db.Prepare("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal("can't prepare query all users statment", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all users", err)
	}

	for rows.Next() {
		var id, age int
		var name string
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}
		fmt.Println(id, name, age)
	}

}
