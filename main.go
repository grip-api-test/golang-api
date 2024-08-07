package main

import (
	"gin-gonic-api/app/router"
	"gin-gonic-api/config"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	init := config.Init()
	app := router.Init(init)

	app.Run(":8080")
}
