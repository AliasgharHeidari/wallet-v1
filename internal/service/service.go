package service

import (
	"errors"
	"log"

	"github.com/AliasgharHeidari/wallet-v1/internal/model"
	"github.com/AliasgharHeidari/wallet-v1/internal/repository/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrNotFound      = errors.New("ErrNotFound")
	ErrInternal      = errors.New("ErrInternal")
	ErrDuplicatedKey = errors.New("DuplicatedKey")
)

func GetWalletInfo(number string) (model.Wallet, error) {
	wallet, err := postgres.GetWalletInfo(number)
	if err != nil {
		return model.Wallet{}, err
	}
	return wallet, nil
}

func Transaction() ([]model.Transaction, error) {
	DB := postgres.GetDB()

	var tx []model.Transaction

	result := DB.Model(&model.Transaction{}).Find(&tx)

	if result.Error != nil {
		return []model.Transaction{}, ErrInternal
	}

	if result.RowsAffected == 0 {
		return []model.Transaction{}, ErrNotFound
	}

	return tx, nil
}

func GetWalletList() ([]model.Wallet, error) {
	DB := postgres.GetDB()

	var Wallets []model.Wallet

	result := DB.Model(&model.Wallet{}).Find(&Wallets)

	if result.Error != nil {
		return []model.Wallet{}, ErrInternal
	}

	if result.RowsAffected == 0 {
		return []model.Wallet{}, ErrNotFound
	}

	return Wallets, nil

}

func CreateAccount(number int) error {
	newWal := model.Wallet{
		MobileNumber: number,
		Balance:      0,
	}
	DB := postgres.GetDB()
	var count int64

	if err := DB.Model(&model.Wallet{}).Where("mobile_number = ?", number).Count(&count).Error; err != nil {
		log.Println(err)
		return err
	}
	log.Println(count)
	if count > 0 {
		return ErrDuplicatedKey
	}

	err := DB.Create(&newWal).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeleteWallet(input model.Wallet) error {
	DB := postgres.GetDB()

	tx := DB.Begin()
	var count model.Wallet

	result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("mobile_number = ?", input.MobileNumber).First(&count)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return ErrNotFound
	}

	if result.Error != nil {
		tx.Rollback()
		return ErrInternal
	}
	
	result = tx.Model(&model.Wallet{}).Where("mobile_number = ?", input.MobileNumber).Delete(&model.Wallet{})
	if result.Error != nil {
		tx.Rollback()
		return ErrInternal
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return ErrInternal
	}

	return nil
}
