package router

import (
	assetsHandler "github.com/PedroXimenes/4invest/internal/app/handlers/assets"
	portfolioHandler "github.com/PedroXimenes/4invest/internal/app/handlers/portfolio"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, https://four-invest-front-p3xh7jp6wa-uc.a.run.app",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/carteiras", portfolioHandler.GetAll)
	app.Get("/carteiras/usuarios/:id", portfolioHandler.Get)
	//app.Get("/carteiras/users/:id/", portfolioHandler.Get)
	app.Post("/carteiras", portfolioHandler.Create)
	app.Post("/ativos", assetsHandler.Create)
	app.Get("/ativos/carteiras/:id", assetsHandler.Get)
}
