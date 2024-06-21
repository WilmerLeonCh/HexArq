package main

import (
	"log"

	"github.com/HexArq/cmd/api/handlers/player"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	ginEngine := gin.Default()

	ginEngine.POST("/players", player.CreatePlayer)

	log.Fatalln(ginEngine.Run(":8081"))
}
