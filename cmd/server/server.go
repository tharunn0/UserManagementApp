package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tharunn0/gin-server-gorm/internal/handlers"
	"github.com/tharunn0/gin-server-gorm/internal/middleware"
	"github.com/tharunn0/gin-server-gorm/internal/routes"
)

type Server struct {
	handler *handlers.Handler
}

func NewServer(h *handlers.Handler) *Server {
	return &Server{
		handler: h,
	}
}

func (s *Server) StartServer() error {
	g := gin.New()
	g.Use(middleware.ZapLogger())

	routes.RegisterRoutes(g, s.handler)

	fmt.Println("App started !")
	if er := g.Run(":3001"); er != nil {
		return er
	}
	return nil
}
