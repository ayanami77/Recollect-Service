package router

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/user"
	"github.com/gin-gonic/gin"
)

func setupRouter(userHandler user.Handler) *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:id", userHandler.GetUser)
		userRouter.POST("/", userHandler.CreateUser)
		userRouter.DELETE("/:id", userHandler.UpdateUser)
		userRouter.PATCH("/:id", userHandler.DeleteUser)
	}

	return router
}
