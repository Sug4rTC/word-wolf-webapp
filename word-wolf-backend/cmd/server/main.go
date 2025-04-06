package main

import (
	"net/http"

	"word-wolf-backend/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.RegisterRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "EchoとWebSocketを使用したサーバー")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
