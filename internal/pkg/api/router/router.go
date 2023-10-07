package router

import (
	"github.com/PedroXimenes/4invest/internal/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/carteiras", handlers.GetAll)
	app.Post("/carteiras", handlers.Create)
}
