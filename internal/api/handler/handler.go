package handler

import (
	"strconv"

	"github.com/AliasgharHeidari/wallet-v1/internal/model"
	"github.com/AliasgharHeidari/wallet-v1/internal/repository/postgres"
	"github.com/AliasgharHeidari/wallet-v1/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetWalletInfo(c *fiber.Ctx) error {
	mobileNumber := c.Params("number")
	number, err := strconv.Atoi(mobileNumber)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid number",
		})
	}

	wallet, err := service.GetWalletInfo(number)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "failed to get wallet information",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"balance":          wallet.Balance,
		"last transaction": wallet.UpdatedAt,
	})

}

func AddCredit(c *fiber.Ctx) error {

	var wallet model.Wallet
	var wal model.Wallet
	amount := 10000
	err := c.BodyParser(&wallet)
	if err != nil {
		return  c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	DB := postgres.GetDB()

	err = DB.Where("mobile_number = ?", wallet.MobileNumber).First(&wal).Error
	if err == gorm.ErrRecordNotFound {

		wal = model.Wallet{
			MobileNumber: wallet.MobileNumber,
			Balance:      0,
		}
		wal.Balance += amount
		err := DB.Save(&wal).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal server error",
			})
		}
	} else if err == nil {

		wal.Balance += amount
		err := DB.Save(&wal).Error
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "internal server error",
			})
		}

	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":         "balance updated",
		"current balance": wallet.Balance,
	})
}
