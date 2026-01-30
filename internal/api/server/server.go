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

	//create wallet account
	app.Post("/wallet/create-account/:number",handler.CreateAccount)
	
	// get balance info
	app.Get("/wallet/:number", handler.GetWalletInfo)

	//get wallet transactions list
	app.Get("/wallet/transactions/:number", handler.Transaction)

	//apply giftCode to wallet (only used by giftcode service)
	app.Post("wallet/gift", handler.AddCredit)

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("faild to load .env, error: ", err)
	}
	port := os.Getenv("SERVER_PORT")

	app.Listen(port)

}
