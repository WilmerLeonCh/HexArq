package player

import (
	"github.com/HexArq/internal/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Client *mongo.Client
}

var _ ports.PlayerRepository = &Repository{}
