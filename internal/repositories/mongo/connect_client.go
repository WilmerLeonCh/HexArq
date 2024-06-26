package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectClient(uriDb string) (client *mongo.Client, err error) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()

	client, errConnectDB := mongo.Connect(ctx, options.Client().ApplyURI(uriDb))
	if errConnectDB != nil {
		return nil, err
	}

	if errPing := client.Ping(ctx, nil); errPing != nil {
		return nil, err
	}
	return client, nil
}
