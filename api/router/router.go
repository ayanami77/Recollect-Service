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
		userRouter.POST("/", userHandler.CreateUser)
		userRouter.PATCH("/:id", userHandler.UpdateUser)
		userRouter.DELETE("/:id", userHandler.DeleteUser)
	}

	return router
}
