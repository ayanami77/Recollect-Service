package router

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/card"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/health"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func New(healthHandler health.Handler, userHandler user.Handler, cardHandler card.Handler) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	// TODO: このようにグルーピングしないと、どういうわけかcorsエラーになってしまう。。
	apiRouter := router.Group("/api")
	{
		// health check
		apiRouter.GET("/health", healthHandler.Check)

		// user
		apiRouter.GET("/user/:id", userHandler.GetUser)
		apiRouter.PATCH("/user/:id", userHandler.UpdateUser)
		apiRouter.DELETE("/user/:id", userHandler.DeleteUser)
		apiRouter.POST("/user/login", userHandler.LoginUser)
		apiRouter.POST("/user/signup", userHandler.CreateUser)
		apiRouter.POST("/user/logout", userHandler.LogoutUser)

		// card
		apiRouter.GET("/card", cardHandler.ListCards)
		apiRouter.POST("/card/new", cardHandler.CreateCard)
		apiRouter.PATCH("/card/:id", cardHandler.UpdateCard)
		apiRouter.DELETE("/card/:id", cardHandler.DeleteCard)
	}

	return router
}
