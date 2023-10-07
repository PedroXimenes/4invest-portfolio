package router

import (
	portfolioHandler "github.com/PedroXimenes/4invest/internal/app/handlers/portfolio"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/carteiras", portfolioHandler.GetAll)
	app.Post("/carteiras", portfolioHandler.Create)
}
