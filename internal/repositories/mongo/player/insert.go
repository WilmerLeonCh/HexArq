package player

import (
	"context"
	"fmt"
	"log"

	"github.com/HexArq/internal/domain"
)

func (r Repository) Insert(player domain.Player) (id interface{}, err error) {
	collection := r.Client.Database("hex-db").Collection("players")
	insertResult, errInsert := collection.InsertOne(context.Background(), player)
	if errInsert != nil {
		log.Println(errInsert)
		return nil, fmt.Errorf("error inserting player: %w", errInsert)
	}
	return insertResult.InsertedID, nil
}
