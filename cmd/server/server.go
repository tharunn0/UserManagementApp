package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tharunn0/gin-server-gorm/internal/handlers"
	"github.com/tharunn0/gin-server-gorm/internal/middleware"
	"github.com/tharunn0/gin-server-gorm/internal/routes"
	"go.uber.org/zap"
)

type Server struct {
	handler *handlers.Handler
	logger  *zap.Logger
}

func NewServer(h *handlers.Handler, l *zap.Logger) *Server {
	return &Server{
		handler: h,
		logger:  l,
	}
}

func (s *Server) StartServer() error {
	g := gin.New()
	g.Use(middleware.ZapLogger(s.logger))

	routes.RegisterRoutes(g, s.handler, s.logger)

	fmt.Println("App started !")
	if er := g.Run(":3001"); er != nil {
		return er
	}
	return nil
}
