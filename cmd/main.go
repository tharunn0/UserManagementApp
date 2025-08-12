package main

import (
	"github.com/tharunn0/gin-server-gorm/cmd/server"
	"github.com/tharunn0/gin-server-gorm/internal/database"
	"github.com/tharunn0/gin-server-gorm/internal/handlers"
	"github.com/tharunn0/gin-server-gorm/internal/models"
	"github.com/tharunn0/gin-server-gorm/internal/repository"
	"github.com/tharunn0/gin-server-gorm/internal/services"
	"github.com/tharunn0/gin-server-gorm/pkg/logger"

	"go.uber.org/zap"
)

func main() {

	logr := logger.InitLogger()
	db, er := database.ConnectToDB()
	if er != nil {
		logr.Error("Could'nt establish db connecton", zap.Error(er))
	}
	logr.Info("Connection established with db")

	er = db.AutoMigrate(&models.User{})
	if er != nil {
		logr.Error("Failed to migrate database", zap.Error(er))
	}
	logr.Info("Db migration successful")

	userRepo := repository.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewHandler(userService)

	server := server.NewServer(userHandler, logr)

	logr.Info("INFO", zap.String("message", "this is a info log"))
	logr.Warn("WARN", zap.String("message", "this is a warn log"))
	logr.Error("ERROR", zap.String("message", "this is a error log"))
	logr.Fatal("FATAL", zap.String("message", "this is a fatal log"))
	logr.Panic("PANIC", zap.String("message", "this is a panic log"))

	er = server.StartServer()
	if er != nil {
		logr.Error("Error", zap.Error(er))
		return
	}

}
