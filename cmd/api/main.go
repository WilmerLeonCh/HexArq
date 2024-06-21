package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Player struct {
	Name         string    `json:"name" binding:"required"`
	Age          int       `json:"age" binding:"required"`
	CreationTime time.Time `json:"-"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	ginEngine := gin.Default()

	ginEngine.POST("/players", func(ginCtx *gin.Context) {
		var player Player

		if errBindJSON := ginCtx.BindJSON(&player); errBindJSON != nil {
			ginCtx.JSON(400, gin.H{"error": errBindJSON.Error()})
			return
		}

		// Set creation time
		player.CreationTime = time.Now().UTC()

		// Set a timeout to allow the connection process to abort if it takes too long
		ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelCtx()

		// Connect to the MongoDB server
		uriDb := os.Getenv("MONGO_URI")
		client, errConnectDB := mongo.Connect(ctx, options.Client().ApplyURI(uriDb))
		if errConnectDB != nil {
			log.Fatal("errConnectDB: ", errConnectDB)
		}

		// Call the ping method to verify that the connection has been established successfully
		if errPing := client.Ping(ctx, nil); errPing != nil {
			log.Fatal("errPing: ", errPing)
		}

		collection := client.Database("hex-db").Collection("players")
		insertResult, errInsert := collection.InsertOne(ctx, player)
		if errInsert != nil {
			log.Fatal("errInsert: ", errInsert)
		}

		ginCtx.JSON(200, gin.H{"player_id": insertResult.InsertedID})
	})

	log.Fatalln(ginEngine.Run(":8081"))
}
