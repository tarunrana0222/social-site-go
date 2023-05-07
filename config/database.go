package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitiateClient() *mongo.Client {
	if !isEnvLoaded {
		LoadEnv()
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DB_Url))
	if err != nil {
		log.Fatal("While connecting DB", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("While ping", err)
	}
	fmt.Println("DB initiated")
	return client
}

var (
	Client    = InitiateClient()
	UsersColl = Client.Database(DB_Name).Collection("users")
	PostsColl = Client.Database(DB_Name).Collection("posts")
)
