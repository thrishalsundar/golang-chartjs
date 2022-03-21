package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetClient() (*mongo.Client, context.Context) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading")
	}

	ctx := context.TODO()
	peru := os.Getenv("MONGO_URI")
	mongoconn := options.Client().ApplyURI(peru)
	client, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodone")
	return client, ctx
}
