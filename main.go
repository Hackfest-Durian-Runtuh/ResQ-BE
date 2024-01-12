package main

import (
	"log"
	"os"
	"resq-be/config"
	"resq-be/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}
	db := config.Init()

	if db == nil {
		log.Fatal("init connection db failed")
	}
	err = config.Migrate()
	if err != nil {
		log.Fatalln(err.Error())
	}
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.Use(middlewares.Timeout(60))

	app.Use(middlewares.Error())

	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}
	app.Run(port)
}
