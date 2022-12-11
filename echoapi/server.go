package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Error struct {
	Message string `json:"message"`
}

var users = []User{
	{ID: 1, Name: "Kaew", Age: 18},
	{ID: 2, Name: "Ewka", Age: 22},
}

func healthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func getUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func createUserHandler(c echo.Context) error {
	var u User
	err := c.Bind(&u)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	}

	users = append(users, u)
	return c.JSON(http.StatusCreated, u)
}

func authValidator(u, p string, c echo.Context) (bool, error) {
	if u == "apidesign" && p == "45678" {
		return true, nil

	}
	return false, nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)

	g := e.Group("api")
	g.Use(middleware.BasicAuth(authValidator))

	g.GET("/users", getUsersHandler)
	g.POST("/users", createUserHandler)

	log.Println("Server started ad : 8080")
	log.Fatal(e.Start(":8080"))
}
