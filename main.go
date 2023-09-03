package main

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/db"
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/user"
	userRepository "github.com/Seiya-Tagami/Recollect-Service/api/infra/repository/user"
	"github.com/Seiya-Tagami/Recollect-Service/api/router"
	userUsecase "github.com/Seiya-Tagami/Recollect-Service/api/usecase/user"
)

func main() {
	dbConn := db.New()
	defer db.Close(dbConn)
	dbConn.AutoMigrate(&entity.User{})

	userRepository := userRepository.New(dbConn)
	userUsecase := userUsecase.New(userRepository)
	userHandler := user.New(userUsecase)
	router := router.New(userHandler)

	router.Run()
}
