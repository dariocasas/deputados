package mongodb

import (
	"context"
	"log"

	"github.com/dariocasas/deputados3/src/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DisconnectMongoDB(client *mongo.Client) error {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Println(err)
		return err
	}
	log.Println("Successfully disconnected from MongoDB.")
	return nil
}

func ConnectMongoDB(cfg *configs.Config) (*mongo.Client, error) {

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.MongodbHost).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").
		RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).
		Decode(&result); err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Ping! Successfully connected to MongoDB.")

	return client, nil
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CollectionExists(db *mongo.Database, collectionName string) bool {

	collNames, err := db.ListCollectionNames(
		context.TODO(),
		bson.D{},
	)
	if err != nil {
		panic(err)
	}

	res := Contains(collNames, collectionName)
	log.Printf("collection %s exists: %t", collectionName, res)
	return res
}

func DbExists(client *mongo.Client, dbName string) bool {

	dbNames, err := client.ListDatabaseNames(
		context.TODO(),
		bson.D{},
	)
	if err != nil {
		panic(err)
	}

	res := Contains(dbNames, dbName)
	log.Printf("db %s exists: %t", dbName, res)
	return res
}
