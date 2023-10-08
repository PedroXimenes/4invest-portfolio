package assets

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/PedroXimenes/4invest/internal/pkg/models/assets"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func Create(c *fiber.Ctx) error {
	assets := []*models.Asset{}

	if err := json.Unmarshal(c.Body(), &assets); err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).SendString("Could not decode request body")
	}
	fmt.Println(assets)
	// key, err := portfolio.ValidateInput()
	// if err != nil {
	// 	errMsg := fmt.Sprintf("Missing key: %s", key)
	// 	return c.Status(http.StatusBadRequest).SendString(errMsg)
	// }

	err := models.Insert(assets)
	if err != nil {
		log.Error(err)
		if err.Error() == `pq: duplicate key value violates unique constraint "unique_name"` {
			return c.Status(http.StatusConflict).SendString("Cannot create multiple portfolios with the same name for the same user. Please choose a diferent name.")
		}
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	message := fmt.Sprintf("Succesfully inserted %d rows.", len(assets))
	return c.Status(http.StatusCreated).SendString(message)
}
