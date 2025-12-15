package server

import (
	"log"
	"os"

	"github.com/AliasgharHeidari/wallet-v1/internal/api/handler"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func Start() {

	app := fiber.New()

	app.Get("/wallet/:number", handler.GetWalletInfo)

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("faild to load .env, error: ", err)
	}
	port := os.Getenv("SERVER_PORT")

	app.Listen(port)

}
