package player

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/HexArq/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreatePLayerService(player domain.Player) (interface{}, error) {
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
		return nil, errConnectDB
	}

	// Call the ping method to verify that the connection has been established successfully
	if errPing := client.Ping(ctx, nil); errPing != nil {
		log.Fatal("errPing: ", errPing)
		return nil, errPing
	}

	collection := client.Database("hex-db").Collection("players")
	insertResult, errInsert := collection.InsertOne(ctx, player)
	if errInsert != nil {
		log.Fatal("errInsert: ", errInsert)
		return nil, errInsert
	}

	return insertResult.InsertedID, nil
}
