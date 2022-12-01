package main

import (
	"finalproject4/config"
	"finalproject4/handler"
	"finalproject4/helper"
	"finalproject4/repository"
	"finalproject4/route"
	"finalproject4/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	cfg := config.LoadConfig()
	db, err := config.ConnectDB(cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	if err != nil {
		panic(err)
	}

	auth := helper.NewService()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, auth)

	route.UserRouter(router, userHandler, db)

	router.Run(":" + cfg.ServerPort)
}
