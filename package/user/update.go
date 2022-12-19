package user

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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
