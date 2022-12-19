package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	db := user.InitDb()
	defer db.Close()

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

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed { // Start server
			e.Logger.Fatal("shutting down the server")
		}
		log.Println("Server started at : 8080")
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Info("Server stopped")
}
