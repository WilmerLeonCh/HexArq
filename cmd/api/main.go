package main

import (
	"log"
	"os"

	playerHandler "github.com/HexArq/cmd/api/handlers/player"
	"github.com/HexArq/internal/repositories/mongo"
	playerRepository "github.com/HexArq/internal/repositories/mongo/player"
	playerService "github.com/HexArq/internal/services/player"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	ginEngine := gin.Default()

	uriDb := os.Getenv("MONGO_URI")
	client, errConnect := mongo.ConnectClient(uriDb)
	if errConnect != nil {
		log.Fatal(errConnect.Error())
	}

	playerRepo := playerRepository.Repository{
		Client: client,
	}

	playerServ := playerService.Service{
		Repo: playerRepo,
	}

	playerHand := playerHandler.Handler{
		PlayerService: playerServ,
	}

	ginEngine.POST("/players", playerHand.Create)

	log.Fatalln(ginEngine.Run(":8081"))
}
