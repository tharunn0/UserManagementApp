package handlers

import (
	"fmt"

	"github.com/tharunn0/gin-server-gorm/internal/middleware/jwt"
	"github.com/tharunn0/gin-server-gorm/internal/models"
	"github.com/tharunn0/gin-server-gorm/internal/services"
	log "github.com/tharunn0/gin-server-gorm/pkg/logger"

	"github.com/gin-gonic/gin"
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
	val, ok := c.Get("username")
	if !ok {
		return
	}
	username, _ := val.(string)
	msg := "Hi " + username
	c.JSON(200, gin.H{"message": msg})
	log.Info("INFO",
		zap.String("from", c.ClientIP()),
		zap.String("to", c.Request.URL.Path),
	)
}

func (h *Handler) CreateUser(c *gin.Context) {

	var RegReq models.RegisterReq
	er := c.ShouldBindJSON(&RegReq)
	if er != nil {
		log.Error(
			"JSON Parse Error",
			zap.Error(er),
		)
		return
	}

	er = h.Service.Create(&RegReq)
	if er != nil {
		log.Error(
			"Service User Creation Error",
			zap.Error(er),
		)
		return
	}

	c.JSON(200, gin.H{
		"message": "Registration Successful",
	})

}

func (h *Handler) LogInUser(c *gin.Context) {

	role, ok := jwt.ValidateToken(c.GetHeader("Authorization"))
	if ok || role == "user" {
		c.JSON(301, gin.H{"redirect": "/home"})
		return
	}

	var LoginReq models.LoginReq

	er := c.ShouldBindJSON(&LoginReq)
	if er != nil {
		log.Error("JSON Parsing", zap.Error(er))
		return
	}

	userprofile, er := h.Service.LoginUser(&LoginReq)
	if er != nil {
		log.Error("LoginUser Error", zap.Error(er))
		c.JSON(301, gin.H{"error": fmt.Sprintf("%s", er)})
		return
	}

	token := jwt.Issue(userprofile.Username, "user")
	log.Info("JWT Issued")

	c.JSON(200, gin.H{"login": "Login successful", "jwt_token": token})

}
