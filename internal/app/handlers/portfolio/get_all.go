package portfolio

import (
	"fmt"
	"net/http"

	models "github.com/PedroXimenes/4invest/internal/pkg/models/portfolio"
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	users, err := models.GetAll()
	if err != nil {
		fmt.Printf("%s\n", err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(http.StatusOK).JSON(users)
}
