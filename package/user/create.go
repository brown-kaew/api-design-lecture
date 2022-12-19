package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

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
