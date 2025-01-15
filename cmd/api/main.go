package main

import (
	"gorestapi/internal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Rutas
	e.POST("/users", handlers.CreateUser)
	e.GET("/users", handlers.GetUsers)

	e.Start(":8080")
}
