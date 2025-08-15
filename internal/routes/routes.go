package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tharunn0/gin-server-gorm/internal/handlers"
	"github.com/tharunn0/gin-server-gorm/internal/middleware/jwt"
)

func RegisterRoutes(r *gin.Engine, handler *handlers.Handler) {

	r.POST("/register", handler.CreateUser)
	r.POST("/login", handler.LogInUser)
	r.POST("admin/login", handler.AdminLogin)

	userRoute := r.Group("/")
	userRoute.Use(jwt.ValidateMiddleware())
	{
		userRoute.GET("/home", handler.GetHomePage)
		userRoute.GET("/profile", handler.GetUserProfile)
		// userRoute.GET("/uploads", handler.GetUploads)
		// userRoute.POST("/logout", handler.LogOutUser)
		// userRoute.POST("/reset-password", handler.ResetUserPassword)
	}

	adminRoute := r.Group("/admin")
	adminRoute.Use(jwt.ValidateMiddlewareAdmin())
	// {
	// 	adminRoute.POST("/create", handler.CreateAdmin)
	adminRoute.GET("/dashboard", handler.GetAdminDashboard)
	adminRoute.GET("/getusers", handler.GetAllUsers)
	adminRoute.DELETE("/delete/:username", handler.DeleteUser)
	// }

}
