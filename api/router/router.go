package router

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/card"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/health"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func New(healthHandler health.Handler, userHandler user.Handler, cardHandler card.Handler) *gin.Engine {
	router := setupRouter()

	router.GET("/health", healthHandler.HealthCheck)

	userRouter := router.Group("/users")
	{
		userRouter.PATCH("/:id", userHandler.UpdateUser)
		//userRouter.DELETE("/:id", userHandler.DeleteUser)
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
		cardRouter.PATCH("/analysis/:id", cardHandler.UpdateAnalysisResult)
	}

	return router
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("ORIGIN_URL")},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	return router
}
