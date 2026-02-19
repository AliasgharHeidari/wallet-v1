package server

import (
	"github.com/AliasgharHeidari/wallet-v1/config"
	"github.com/AliasgharHeidari/wallet-v1/internal/api/handler"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start(cfg config.Server) {

	app := fiber.New()
	app.Use(logger.New())

	//apply giftCode to wallet (only used by giftcode service)
	app.Post("wallet/topup", handler.AddCredit)

	//create wallet account
	app.Post("/wallet/:number", handler.CreateAccount)

	//get wallet list
	app.Get("/wallet", handler.GetWalletList)

	//get wallet transactions list
	app.Get("/wallet/transaction", handler.Transaction)

	// get wallet info
	app.Get("/wallet/:number", handler.GetWalletInfo)

	// Delete wallet account
	app.Delete("/wallet/", handler.DeleteWallet)

	app.Listen(cfg.Port)

}
