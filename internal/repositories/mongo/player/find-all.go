package player

import (
	"context"
	"fmt"
	"log"

	"github.com/HexArq/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (r Repository) FindAll() (players []domain.Player, err error) {
	collection := r.Client.Database("hex-db").Collection("players")
	cursor, errFind := collection.Find(context.Background(), bson.D{})
	if errFind != nil {
		log.Println(errFind)
		return nil, fmt.Errorf("error finding players: %w", errFind)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		if errCursor := cursor.Err(); errCursor != nil {
			return nil, errCursor
		}
		var player domain.Player
		errDecode := cursor.Decode(&player)
		if errDecode != nil {
			return nil, errDecode
		}
		players = append(players, player)
	}
	return players, nil
}
