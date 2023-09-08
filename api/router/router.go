package router

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func New(userHandler user.Handler) *gin.Engine {
	router := gin.Default()

	setCors(router)

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:id", userHandler.GetUser)
		//userRouter.POST("/", userHandler.CreateUser)
		userRouter.PATCH("/:id", userHandler.UpdateUser)
		userRouter.DELETE("/:id", userHandler.DeleteUser)
		userRouter.POST("/signup", userHandler.CreateUser)
		userRouter.POST("/login", userHandler.LoginUser)
		//userRouter.GET("/logout", userHandler.LogoutUser)
	}

	return router
}

func setCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
