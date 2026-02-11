package server

import (
	"log"
	"os"

	"github.com/AliasgharHeidari/wallet-v1/internal/api/handler"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func Start() {

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

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("faild to load .env, error: ", err)
	}
	port := os.Getenv("SERVER_PORT")

	app.Listen(port)

}
