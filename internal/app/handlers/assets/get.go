package assets

import (
	"fmt"
	"net/http"
	"strconv"

	models "github.com/PedroXimenes/4invest/internal/pkg/models/assets"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func Get(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		log.Error("The path param 'id' is required")
		return c.Status(http.StatusBadRequest).SendString("The path param 'id' is required")
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).SendString("The path param 'id' must be an integer")
	}

	assets, err := models.Get(id)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	if len(assets) == 0 {
		errMsg := fmt.Sprintf("No results for user: %d", id)
		return c.Status(http.StatusNotFound).JSON(errMsg)
	}

	return c.Status(http.StatusCreated).JSON(assets)
}
