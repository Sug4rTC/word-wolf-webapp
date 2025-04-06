package routes

import (
	"word-wolf-backend/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/ws", handlers.WebSocketHandler)
}
