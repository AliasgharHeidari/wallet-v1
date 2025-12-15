package postgres

import (
	"log"
	"os"
	"github.com/AliasgharHeidari/wallet-v1/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	err error
	DB  *gorm.DB
)

func InitDB() {

	if err := godotenv.Load("./.env"); err != nil {
		panic(err)
	}
	dsn := os.Getenv("DSN")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database, error:", err)
	}

}

func AutoMigrate() {
	err := DB.AutoMigrate(&model.Wallet{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return DB
}
