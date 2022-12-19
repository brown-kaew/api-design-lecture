package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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
