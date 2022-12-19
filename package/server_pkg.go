package main

import (
	"log"
	"net/http"

	"github.com/brown-kaew/api-design-lecture/package/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func healthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func authValidator(u, p string, c echo.Context) (bool, error) {
	if u == "apidesign" && p == "45678" {
		return true, nil

	}
	return false, nil
}

func main() {
	user.InitDb()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)

	g := e.Group("api")
	g.Use(middleware.BasicAuth(authValidator))

	g.GET("/users", user.GetUsersHandler)
	g.GET("/users/:id", user.GetUserByIdHandler)
	g.POST("/users", user.CreateUserHandler)
	g.PUT("/users/:id", user.UpdateUserByIdHandler)
	g.DELETE("/users/:id", user.DeleteUserByIdHandler)

	log.Println("Server started ad : 8080")
	log.Fatal(e.Start(":8080"))
}
