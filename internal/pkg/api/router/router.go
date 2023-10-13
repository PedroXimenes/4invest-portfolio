package router

import (
	assetsHandler "github.com/PedroXimenes/4invest/internal/app/handlers/assets"
	portfolioHandler "github.com/PedroXimenes/4invest/internal/app/handlers/portfolio"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router(app *fiber.App) {
	app.Use(cors.New())

	app.Get("/carteiras", portfolioHandler.GetAll)
	app.Get("/carteiras/users/:id", portfolioHandler.Get)
	//app.Get("/carteiras/users/:id/", portfolioHandler.Get)
	app.Post("/carteiras", portfolioHandler.Create)
	app.Post("/assets", assetsHandler.Create)
	app.Get("/assets/carteiras/:id", assetsHandler.Get)
}
