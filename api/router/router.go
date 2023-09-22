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

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:id", userHandler.GetUser)
		userRouter.PATCH("/:id", userHandler.UpdateUser)
		userRouter.DELETE("/:id", userHandler.DeleteUser)
		userRouter.POST("/login", userHandler.LoginUser)
		userRouter.POST("/signup", userHandler.CreateUser)
		userRouter.POST("/logout", userHandler.LogoutUser)
	}

	cardRouter := router.Group("/card")
	{
		cardRouter.GET("/list", cardHandler.ListCards)
		cardRouter.POST("/new", cardHandler.CreateCard)
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
