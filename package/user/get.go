package user

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
