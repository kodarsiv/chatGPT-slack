package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"kodarsiv/chatGPT-slack/internal/storage"
	"kodarsiv/chatGPT-slack/internal/types"
	"os"
)

type Mongo struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func NewMongo() (storage.Storage, error) {

	credential := options.Credential{
		Username: os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_CON_URI")).SetAuth(credential))
	//client, err := mongo.NewClient(options.Client().ApplyURI(generateConnectionURI())) // for remote database connection
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	database := client.Database(os.Getenv("DATABASE_NAME"))
	collection := database.Collection(os.Getenv("COLLECTION_NAME"))

	return &Mongo{
		Client:     client,
		Database:   database,
		Collection: collection,
	}, nil
}

func (m *Mongo) PutMessage(message types.Message) (bool, error) {
	// TODO:  will updated
	return false, nil
}
