package routes

import (
	"resq-be/controllers"
	"resq-be/repositories"
	"resq-be/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(db *gorm.DB, r *gin.RouterGroup) {
	userRepo := repositories.NewUser(db)
	userUsecase := usecases.NewUser(userRepo)
	userController := controllers.NewUser(userUsecase)
	user := r.Group("/user")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
		user.PUT("/:id", userController.Update)
	}
}
