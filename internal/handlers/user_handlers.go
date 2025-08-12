package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tharunn0/gin-server-gorm/internal/models"
	"github.com/tharunn0/gin-server-gorm/internal/services"
	"github.com/tharunn0/gin-server-gorm/pkg/logger"
	"go.uber.org/zap"
)

type Handler struct {
	Service *services.UserService
}

func NewHandler(h *services.UserService) *Handler {
	return &Handler{
		Service: h,
	}
}

func (h *Handler) GetHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "In the homepage",
	})
	logger.Logr.Info(
		"INFO",
		zap.String("from", c.ClientIP()),
		zap.String("to", c.Request.URL.Path),
	)

}

func (h *Handler) CreateUser(c *gin.Context) {

	var RegReq models.RegisterReq

	// Email     string `json:"email"`
	// 	Username  string `json:"username"`
	// 	Firstname string `json:"first_name"`
	// 	Lastname  string `json:"last_name"`
	// 	Password  string `json:"password"`

	er := c.ShouldBindJSON(&RegReq)
	if er != nil {
		logger.Logr.Error(
			"ERROR",
			zap.Error(er),
		)
		return
	}

	er = h.Service.Create(&RegReq)
	if er != nil {
		logger.Logr.Error(
			"ERROR",
			zap.Error(er),
		)
		return
	}

	c.JSON(200, gin.H{
		"message": "Registration Successful",
	})

}
