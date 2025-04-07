package routes

import (
	"word-wolf-backend/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "EchoとWebSocketを使用したサーバー")
	})

	e.GET("/ws", handlers.WebSocketHandler)

	e.POST("/rooms", handlers.CreateRoom)
}
