package user

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Error struct {
	Message string `json:"message"`
}
