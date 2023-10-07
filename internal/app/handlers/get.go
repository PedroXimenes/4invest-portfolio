package handlers

// func Get(c *fiber.Ctx) error {
// 	id, userId, err := parseIds(c)
// 	if err != nil {
// 		return err
// 	}
// 	user, err := models.Get(id, userId)
// 	if err != nil {
// 		fmt.Printf("%s\n", err)
// 		if err.Error() == "sql: no rows in result set" {
// 			errMsg := fmt.Sprintf("No results for user: %d", id)
// 			return c.Status(http.StatusNotFound).SendString(errMsg)
// 		}
// 		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
// 	}

// 	return c.Status(http.StatusOK).JSON(user)
// }

// func parseIds(c *fiber.Ctx) (int64, int64, error) {
// 	idStr := c.Params("id")
// 	if idStr == "" {
// 		return 0, 0, c.Status(http.StatusBadRequest).SendString("The path param 'id' is required")
// 	}

// 	id, err := strconv.ParseInt(idStr, 10, 64)
// 	if err != nil {
// 		return 0, 0, c.Status(http.StatusBadRequest).SendString("The path param 'id' must be an integer")
// 	}
// 	userIdStr := c.Params("userid")
// 	if idStr == "" {
// 		return 0, 0, c.Status(http.StatusBadRequest).SendString("The path param 'id' is required")
// 	}

// 	userId, err := strconv.ParseInt(userIdStr, 10, 64)
// 	if err != nil {
// 		return 0, 0, c.Status(http.StatusBadRequest).SendString("The path param 'id' must be an integer")
// 	}

// 	return id, userId, nil
// }
