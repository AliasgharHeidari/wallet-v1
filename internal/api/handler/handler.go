package handler

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/AliasgharHeidari/wallet-v1/internal/model"
	"github.com/AliasgharHeidari/wallet-v1/internal/repository/postgres"
	"github.com/AliasgharHeidari/wallet-v1/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateAccount(c *fiber.Ctx) error {
	StringNumber := c.Params("number")

	if len(StringNumber) > 16 || len(StringNumber) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "the Number most be between 8-16 characters",
		})
	}

	Number, err := strconv.Atoi(StringNumber)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request url, only numbers are allowed",
		})
	}

	err = service.CreateAccount(Number)
	if errors.Is(err, service.) {
		log.Println(err)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "this ID is used by anoher account",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal error, please try again later",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "account has been created successfuly",
		"createdAt": time.Now(),
	})
}

func GetWalletInfo(c *fiber.Ctx) error {
	number := c.Params("number")

	wallet, err := service.GetWalletInfo(number)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "failed to get wallet information",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"balance": wallet.Balance,
	})

}

func Transaction(c *fiber.Ctx) error {
	number := c.Params("number")

	Transactions, err := service.Transaction(number)
	if errors.Is(err, service.ErrNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "wallet does not exist",
		})
	}
	if errors.Is(err, service.ErrInternal) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"transaction-ID": Transactions.ID,
		"value":          Transactions.Value,
		"date":           Transactions.CreatedAt,
	})

}

func AddCredit(c *fiber.Ctx) error {

	var wallet model.Wallet
	var wal model.Wallet
	var amount float64 = 1000000
	err := c.BodyParser(&wallet)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	DB := postgres.GetDB()

	err = DB.Where("mobile_number = ?", wallet.MobileNumber).First(&wal).Error
	if err == gorm.ErrRecordNotFound {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "wallet does not exist",
		})
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
		"current balance": wal.Balance,
	})
}
