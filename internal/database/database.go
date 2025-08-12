package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {

	er := godotenv.Load()
	if er != nil {
		log.Fatalln("Couldnt load environment varialbes :", er)
	}

	dsn := os.Getenv("DB_CRED")
	db, er := gorm.Open(postgres.Open(dsn))
	if er != nil {
		return nil, er
	}

	return db, nil
}
