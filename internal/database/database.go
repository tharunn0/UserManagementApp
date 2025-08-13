package database

import (
	"os"

	"github.com/tharunn0/gin-server-gorm/internal/models"
	log "github.com/tharunn0/gin-server-gorm/pkg/logger"
	"go.uber.org/zap"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {

	er := godotenv.Load()
	if er != nil {
		log.Fatal("Couldnt load environment varialbes :", zap.Error(er))
	}

	dsn := os.Getenv("DB_CRED")
	db, er := gorm.Open(postgres.Open(dsn))
	if er != nil {
		return nil, er
	}

	er = db.AutoMigrate(&models.User{})
	if er != nil {
		log.Error("Failed to migrate database", zap.Error(er))
	}
	log.Info("Db migration successful")

	return db, nil
}
