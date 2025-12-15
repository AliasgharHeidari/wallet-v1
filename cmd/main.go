package main

import (
	onmemory "github.com/AliasgharHeidari/wallet-v1/internal/repository/on-memory"
	"github.com/AliasgharHeidari/wallet-v1/internal/repository/postgres"
)

func main() {
	postgres.InitDB()
	postgres.AutoMigrate()
	onmemory.InitWallet()
}