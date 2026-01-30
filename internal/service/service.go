package service

import (
	"errors"
	"log"

	"github.com/AliasgharHeidari/wallet-v1/internal/model"
	"github.com/AliasgharHeidari/wallet-v1/internal/repository/postgres"
	"gorm.io/gorm"
)

var (
	ErrNotFound      error
	ErrInternal      error
 	ErrDuplicatedKey string = "ErrDuplicatedKey"

)

func GetWalletInfo(number string) (model.Wallet, error) {
	wallet, err := postgres.GetWalletInfo(number)
	if err != nil {
		return model.Wallet{}, err
	}
	return wallet, nil
}

func Transaction(number string) (model.Transaction, error) {
	DB := postgres.GetDB()

	var tx model.Transaction

	err := DB.Where("phone_number = ?", number).First(&tx).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Transaction{}, ErrNotFound
		} else {
			return model.Transaction{}, ErrInternal
		}
	}

	return tx, nil
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
		errorKey := errors.New("DuplicatedKey")
		return errorKey
	}

	err := DB.Create(&newWal).Error
	if err != nil {
		log.Println(err)
		return err
	}

	/* 	err := DB.Create(&newWal).Error
	   	if err != nil {
	   		if errors.Is(err, gorm.ErrDuplicatedKey) {
	   			return ErrDuplicatedKey
	   		}
	   		return err
	   	} */
	return nil
}
