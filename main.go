package main

import (
	"log"
	"os"
	"resq-be/config"
	"resq-be/middlewares"
	"resq-be/model"
	"resq-be/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}
	db := config.InitDB()

	if db == nil {
		log.Fatal("init connection db failed")
	}
	err = config.MigrateDB(model.User{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	gin.SetMode(gin.DebugMode)
	app := gin.New()
	app.NoRoute(func(c *gin.Context) {
		middlewares.NotFound()(c)
	})
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.Use(middlewares.Timeout(60))
	v1 := app.Group("/api/v1")
	routes.User(db, v1)
	port := ":8081"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}
	app.Run(port)
}
