package user

import (
	"database/sql"
	"log"
	"os"
)

func InitDb() {
	url := os.Getenv("DATABASE_URL")
	var err error
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("connect to database error", err)
	}

	createTb := `CREATE TABLE IF NOT EXISTS users ( id SERIAL PRIMARY KEY, name TEXT, age INT );`
	if _, err := db.Exec(createTb); err != nil {
		log.Fatal("can't create table", err)
	}
}
