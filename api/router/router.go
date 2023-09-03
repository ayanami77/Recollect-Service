package router

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/user"
	"github.com/gin-gonic/gin"
)

func New(userHandler user.Handler) *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:id", userHandler.GetUser)
		//userRouter.POST("/", userHandler.CreateUser)
		userRouter.PATCH("/:id", userHandler.UpdateUser)
		userRouter.DELETE("/:id", userHandler.DeleteUser)
		userRouter.POST("/signup", userHandler.CreateUser)
		userRouter.GET("/login", userHandler.LoginUser)
		//userRouter.GET("/logout", userHandler.LogoutUser)
	}

	return router
}
