package router

import (
	assetsHandler "github.com/PedroXimenes/4invest/internal/app/handlers/assets"
	portfolioHandler "github.com/PedroXimenes/4invest/internal/app/handlers/portfolio"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/carteiras", portfolioHandler.GetAll)
	app.Get("/carteiras/users/:id", portfolioHandler.Get)
	//app.Get("/carteiras/users/:id/", portfolioHandler.Get)
	app.Post("/carteiras", portfolioHandler.Create)
	app.Post("/assets", assetsHandler.Create)
}
