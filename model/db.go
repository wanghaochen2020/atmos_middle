package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Db *mongo.Database

func InitDb() {
	applyURI := fmt.Sprintf("mongodb://admin:bupt2022@43.138.78.252:27017")
	clientOptions := options.Client().ApplyURI(applyURI)
	var ctx = context.TODO()
	// Connect to MongoDB
	DbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = DbClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	Db = DbClient.Database("sanitation")
}
