package user

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
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

func GetUsersHandler(c echo.Context) error {
	stmt, err := db.Prepare("SELECT id, name, age FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{Message: "can't prepare query all users statment:" + err.Error()})
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{Message: "can't query all users:" + err.Error()})
	}

	var users = []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Error{Message: "can't scan user:" + err.Error()})
		}
		users = append(users, user)
	}
	return c.JSON(http.StatusOK, users)
}

func GetUserByIdHandler(c echo.Context) error {
	stmt, err := db.Prepare("SELECT id, name, age FROM users WHERE id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{Message: "can't prepare query all users statment"})
	}

	id := c.Param("id")
	row := stmt.QueryRow(id)

	var user User
	err = row.Scan(&user.ID, &user.Name, &user.Age)

	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Error{Message: "user not found"})
	case nil:
		return c.JSON(http.StatusOK, user)
	default:
		return c.JSON(http.StatusInternalServerError, Error{Message: "can't Scan row into variable"})
	}
}

func CreateUserHandler(c echo.Context) error {
	var u User
	err := c.Bind(&u)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}

	row := db.QueryRow("INSERT INTO users (name, age) values ($1, $2) RETURNING id", u.Name, u.Age)

	if err := row.Scan(&u.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, u)
}

func UpdateUserByIdHandler(c echo.Context) error {
	id := c.Param("id")
	var u User
	err := c.Bind(&u)
	u.ID, err = strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: "invalid id"})
	}

	stmt, err := db.Prepare(`
	UPDATE users 
	SET 
		name=$2,
		age=$3
	WHERE id=$1`)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{Message: "can't prepare query all users statment" + err.Error()})
	}
	var res sql.Result
	res, err = stmt.Exec(id, u.Name, u.Age)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{Message: "can't update user " + err.Error()})
	}

	row, err := res.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{Message: "can't update user " + err.Error()})
	}

	if row == 0 {
		return c.JSON(http.StatusNotFound, Error{Message: "not found user"})
	}
	return c.JSON(http.StatusOK, u)
}

func DeleteUserByIdHandler(c echo.Context) error {
	id := c.Param("id")
	deleteId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}

	stmt, err := db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Error{Message: "can't prepare query all users statment" + err.Error()})
	}

	if _, err := stmt.Exec(deleteId); err != nil {
		return c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, Error{Message: "deleted id: " + id})
}
