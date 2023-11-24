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
	router := setupRouter()

	router.GET("/health", healthHandler.HealthCheck)

	userRouter := router.Group("/users")
	{
		userRouter.PATCH("", userHandler.UpdateUser)
		//userRouter.DELETE("", userHandler.DeleteUser)
		userRouter.POST("/signup", userHandler.CreateUser)
		userRouter.POST("/email-duplicate-check", userHandler.CheckEmailDuplication)
		userRouter.POST("/id-duplicate-check", userHandler.CheckUserIDDuplication)
	}

	cardRouter := router.Group("/cards")
	{
		cardRouter.GET("", cardHandler.ListCards)
		cardRouter.POST("", cardHandler.CreateCard)
		cardRouter.POST("/batch", cardHandler.CreateCards)
		cardRouter.PATCH("/:id", cardHandler.UpdateCard)
		cardRouter.DELETE("/:id", cardHandler.DeleteCard)
	}

	return router
}

func setupRouter() *gin.Engine {
	//TODO: CORSの設定を適切にする
	//router := gin.Default()
	//router.Use(cors.New(cors.Config{
	//	AllowOrigins: []string{"http://localhost:3000"},
	//	AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	//	AllowHeaders: []string{"Origin", "Content-Type"},
	//	MaxAge: 24 * time.Hour,
	//}))
	//return router

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

	return router
}
