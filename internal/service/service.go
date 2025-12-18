package service

import (
	"github.com/AliasgharHeidari/wallet-v1/internal/model"
	"github.com/AliasgharHeidari/wallet-v1/internal/repository/postgres"
)

func GetWalletInfo(number int)(model.Wallet, error) {
	wallet, err := postgres.GetWalletInfo(number)
	if err != nil {
		return model.Wallet{}, err
	}
	return wallet, nil
}
