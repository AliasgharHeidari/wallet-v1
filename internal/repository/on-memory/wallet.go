package onmemory

import (
	"log"
	"github.com/AliasgharHeidari/wallet-v1/internal/model"
	"github.com/AliasgharHeidari/wallet-v1/internal/repository/postgres"
)

var wallet []model.Wallet

func InitWallet() {
	DB := postgres.GetDB()

	for i := 0; i < 100; i++ {
		InitWal := model.Wallet{
			MobileNumber: 900000 + i,
			Balance:      0,
		}
		log.Println(i)
		log.Println(InitWal.MobileNumber)
		wallet = append(wallet, InitWal)
	}

	for i := range wallet {
		wal := wallet[i]
		if err := DB.Create(&wallet[i]).Error; err != nil {
			log.Printf("failed to add wallet %d to database, %s", wal.MobileNumber, err)
		}
	}

}
