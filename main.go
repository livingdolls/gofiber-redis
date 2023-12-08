package main

import (
	"fmt"
	"gofiber-redis/config"
	"gofiber-redis/controller"
	"gofiber-redis/database"
	"gofiber-redis/model"
	"gofiber-redis/repository"
	"gofiber-redis/router"
	"gofiber-redis/usecase"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello World")

	// Connection database

	loadConfig, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("cannto load env variables", err)
	}

	db 	:= database.ConnectionDB(&loadConfig)

	db.AutoMigrate(&model.Novel{})


	rdb := database.ConnectionRedis(&loadConfig)

	startServer(db,rdb)
}

func startServer(db *gorm.DB, rdb *redis.Client){
	app := fiber.New()

	novelRepo := repository.NewNovelRepo(db, rdb)
	novelUseCase := usecase.NewNovelUseCase(novelRepo)
	novelController := controller.NewNovelController(novelUseCase)

	routes := router.NewRouter(app, novelController)

	err := routes.Listen(":3400")

	if err != nil {
		panic(err)
	}
}