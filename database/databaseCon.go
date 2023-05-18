package database

import (
	"context"
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go/mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
	
)

func DBInstance() *mongo.Client{
	error := godotenv.Load(".env")
	if error != nil{
		log.Fatal("Error loading .env files")
	}

	MongoDB := os.Getenv("MONGO_URL")

	client, error := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if error != nil{
		log.Fatal(error)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	error := client.Connect(ctx)
	if error != nil{
		log.Fatal(error)
	}
	fmt.Println("Connected to mongDB")
}

var client *mongo.client = DBInstance()

func Opencollection(client *mongo.Client, collectionName string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("Cluster0").Collection(collectionName)
	return collection
}