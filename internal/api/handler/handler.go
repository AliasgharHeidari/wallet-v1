package handler

import (
	"strconv"

	"github.com/gofiber/fiber"
)

func GetWalletInfo(c *fiber.Ctx) error {
	mobileNumber := c.Params("number")
	number, err := strconv.Atoi(mobileNumber)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid number",
		})
	}

	return 
}
