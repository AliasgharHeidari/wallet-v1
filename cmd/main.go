package main

import (
	"github.com/AliasgharHeidari/wallet-v1/internal/api/server"
	/* onmemory "github.com/AliasgharHeidari/wallet-v1/internal/repository/on-memory" */
	"github.com/AliasgharHeidari/wallet-v1/internal/repository/postgres"
)

func main() {
	postgres.InitDB()
	postgres.AutoMigrate()
/* 	onmemory.InitWallet() */
	server.Start()
}