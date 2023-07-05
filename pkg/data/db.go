package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"loggerservice/pkg/config"
	"loggerservice/pkg/utils"
)

func createDSN() string {
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		config.EnvConfigs.MongoUser,
		config.EnvConfigs.MongoPass,
		config.EnvConfigs.MongoHost,
		config.EnvConfigs.MongoPort)

	return dsn
}

func InitDB() *mongo.Client {

	log.Default().Println("Connecting to MongoDB...")

	dbURI := createDSN()

	clientOptions := options.Client().ApplyURI(dbURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	utils.FatalErr("Error while connecting to MongoDB", err)

	err = client.Ping(context.TODO(), nil)
	utils.FatalErr("Error while connecting to MongoDB", err)

	log.Default().Println("Connected to MongoDB.")

	return client
}

func CloseDB(client *mongo.Client) {

	log.Default().Println("Closing MongoDB connection...")

	err := client.Disconnect(context.TODO())
	utils.FatalErr("Error while closing the MongoDB connection", err)

	log.Default().Println("MongoDB connection closed.")
}
