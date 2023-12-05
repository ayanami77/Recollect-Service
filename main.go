package main

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/db"
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/card"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/health"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/user"
	cardRepository "github.com/Seiya-Tagami/Recollect-Service/api/infra/repository/card"
	userRepository "github.com/Seiya-Tagami/Recollect-Service/api/infra/repository/user"
	"github.com/Seiya-Tagami/Recollect-Service/api/router"
	cardUsecase "github.com/Seiya-Tagami/Recollect-Service/api/usecase/card"
	userUsecase "github.com/Seiya-Tagami/Recollect-Service/api/usecase/user"
)

func main() {
	dbConn := db.New()
	defer db.Close(dbConn)
	dbConn.AutoMigrate(&entity.User{})
	dbConn.AutoMigrate(&entity.Card{})

	healthHandler := health.New()
	userRepository := userRepository.New(dbConn)
	userUsecase := userUsecase.New(userRepository)
	userHandler := user.New(userUsecase)
	cardRepository := cardRepository.New(dbConn)
	cardUsecase := cardUsecase.New(cardRepository)
	cardHandler := card.New(cardUsecase)

	router := router.New(healthHandler, userHandler, cardHandler)

	router.Run()
}
