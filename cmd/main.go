package main

import (
	"log"

	"github.com/AliasgharHeidari/wallet-v1/config"
	"github.com/AliasgharHeidari/wallet-v1/internal/api/server"

	/* onmemory "github.com/AliasgharHeidari/wallet-v1/internal/repository/on-memory" */
	"github.com/AliasgharHeidari/wallet-v1/internal/repository/postgres"
)

func main() {
	cfg, err := config.Load("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("running on Port :", cfg.Server.Port)
	log.Println("server Host :", cfg.Server.Host)
	postgres.InitDB(cfg.Database)
	postgres.AutoMigrate()
	/* onmemory.InitWallet()  */
	server.Start(cfg.Server)
}
