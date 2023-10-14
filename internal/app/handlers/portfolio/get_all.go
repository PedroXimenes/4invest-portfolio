package portfolio

import (
	"net/http"

	models "github.com/PedroXimenes/4invest/internal/pkg/models/portfolio"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func GetAll(c *fiber.Ctx) error {
	users, err := models.GetAll()
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(http.StatusOK).JSON(users)
}
