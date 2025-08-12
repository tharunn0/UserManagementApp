package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tharunn0/gin-server-gorm/internal/handlers"
	"go.uber.org/zap"
)

func RegisterRoutes(r *gin.Engine, handler *handlers.Handler, logger *zap.Logger) {

	r.POST("/register", handler.CreateUser)
	// r.POST("/login", handler.LogInUser)
	// r.POST("admin/login", handler.AdminLogIn)

	userRoute := r.Group("/")
	// userRoute.Use(middleware.ValidateSession())
	{
		userRoute.GET("/home", handler.GetHomePage)
		// userRoute.GET("/profile", handler.GetUserProfile)
		// userRoute.GET("/uploads", handler.GetUploads)
		// userRoute.POST("/logout", handler.LogOutUser)
		// userRoute.POST("/reset-password", handler.ResetUserPassword)
	}

	// adminRoute := r.Group("/admin")
	// // adminRoute.Use(middlewares.ValidateJWT())
	// {
	// 	adminRoute.POST("/create", handler.CreateAdmin)
	// 	adminRoute.GET("/dashboard", handler.GetAdminDashboard)
	// 	adminRoute.GET("/getusers", handler.GetUsersByName)
	// 	adminRoute.DELETE("/users", handler.DeleteUser)
	// }

}
