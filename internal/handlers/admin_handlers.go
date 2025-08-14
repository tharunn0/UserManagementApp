package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tharunn0/gin-server-gorm/internal/middleware/jwt"
	"github.com/tharunn0/gin-server-gorm/internal/models"
	log "github.com/tharunn0/gin-server-gorm/pkg/logger"
	"go.uber.org/zap"
)

func (h *Handler) AdminLogin(c *gin.Context) {

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

	if !userprofile.IsAdmin {
		c.JSON(401, gin.H{"error": "Not an admin"})
		return
	}

	token := jwt.Issue(userprofile.Username, "admin")
	log.Info("JWT Issued")

	c.JSON(200, gin.H{"login": "Login successful", "jwt_token": token, "redirect": "/admin/dashboard"})

}

func (h *Handler) GetAdminDashboard(c *gin.Context) {
	c.JSON(200, gin.H{"message": "In the admin dash"})
}

func (h *Handler) GetAllUsers(c *gin.Context) {

	var userprofiles []*models.UserProfile
	var res []gin.H
	userprofiles, er := h.Service.GetAllUsers()
	if er != nil {
	}

	for _, up := range userprofiles {
		res = append(res, gin.H{
			"email":     up.Email,
			"username":  up.Username,
			"firstname": up.Firstname,
			"lastname":  up.Lastname,
		})
	}

	c.JSON(200, res)

}

func (h *Handler) DeleteUser(c *gin.Context) {

}
