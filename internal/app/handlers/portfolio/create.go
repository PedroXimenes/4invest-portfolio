package portfolio

import (
	"fmt"
	"net/http"

	models "github.com/PedroXimenes/4invest/internal/pkg/models/portfolio"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func Create(c *fiber.Ctx) error {
	portfolio := &models.Portfolio{}
	err := c.BodyParser(portfolio)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).SendString("Could not decode request body")
	}
	key, err := portfolio.ValidateInput()
	if err != nil {
		errMsg := fmt.Sprintf("Missing key: %s", key)
		return c.Status(http.StatusBadRequest).SendString(errMsg)
	}
	id, err := models.Insert(portfolio)
	if err != nil {
		log.Error(err)
		if err.Error() == `pq: duplicate key value violates unique constraint "unique_name"` {
			return c.Status(http.StatusConflict).SendString("Cannot create multiple portfolios with the same name for the same user. Please choose a diferent name.")
		} else if err.Error() == `pq: insert or update on table "portfolio" violates foreign key constraint "portfolio_user_id_fkey"` {
			log.Error(err)
			return c.Status(http.StatusNotAcceptable).SendString("Please provide a valid user_id.")

		}
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	message := fmt.Sprintf("Created with ID: %d.", id)
	return c.Status(http.StatusCreated).SendString(message)
}
