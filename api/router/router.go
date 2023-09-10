package router

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/card"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func New(userHandler user.Handler, cardHandler card.Handler) *gin.Engine {
	router := gin.Default()

	//setCookieManager(router)
	setCors(router)

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:id", userHandler.GetUser)
		userRouter.PATCH("/:id", userHandler.UpdateUser)
		userRouter.DELETE("/:id", userHandler.DeleteUser)
		userRouter.POST("/signup", userHandler.CreateUser)
		userRouter.POST("/login", userHandler.LoginUser)
		userRouter.POST("/logout", userHandler.LogoutUser)
	}

	cardRouter := router.Group("/card")
	{
		cardRouter.GET("/:id", cardHandler.GetCard)
		cardRouter.GET("/", cardHandler.ListCards)
		cardRouter.PATCH("/:id", cardHandler.UpdateCard)
		cardRouter.DELETE("/:id", cardHandler.DeleteCard)
		cardRouter.POST("/new", cardHandler.CreateCard)
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

//func setCookieManager(router *gin.Engine) {
//	store := cookie.NewStore([]byte("secret"))
//	router.Use(sessions.Sessions("mysession", store))
//}

//func setCsrf(router *gin.Engine) {
//	router.Use(csrf.Middleware(csrf.Options{
//		CookiePath:     "/",
//		CookieDomain:   os.Getenv("API_DOMAIN"),
//		CookieHTTPOnly: true,
//		CookieSameSite: http.SameSiteNoneMode,
//		// CookieSameSite: http.SameSiteDefaultMode,
//		// CookieMaxAge: 60,
//	}))
//}
