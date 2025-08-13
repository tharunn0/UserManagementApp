package main

import (
	"github.com/tharunn0/gin-server-gorm/cmd/server"
	"github.com/tharunn0/gin-server-gorm/internal/database"
	"github.com/tharunn0/gin-server-gorm/internal/handlers"
	"github.com/tharunn0/gin-server-gorm/internal/repository"
	"github.com/tharunn0/gin-server-gorm/internal/services"
	log "github.com/tharunn0/gin-server-gorm/pkg/logger"

	"go.uber.org/zap"
)

func main() {

	log.InitLogger()
	db, er := database.ConnectToDB()
	if er != nil {
		log.Error("Could'nt establish db connecton", zap.Error(er))
	}
	log.Info("Connection established with db")

	userRepo := repository.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewHandler(userService)

	server := server.NewServer(userHandler)

	log.Info("INFO", zap.String("message", "this is a info log"))
	log.Warn("WARN", zap.String("message", "this is a warn log"))
	log.Error("ERROR", zap.String("message", "this is a error log"))

	// logr.Fatal("FATAL", zap.String("message", "this is a fatal log"))
	// logr.Panic("PANIC", zap.String("message", "this is a panic log"))

	er = server.StartServer()
	if er != nil {
		log.Error("Error", zap.Error(er))
		return
	}

}
