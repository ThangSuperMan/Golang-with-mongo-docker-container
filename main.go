package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI = "mongodb://admin:password@localhost:27017"
)

func haltOn(err error) {
	if err != nil {
		log.Fatal("Error here: ", err)
	}
}

func main() {
	fmt.Println("hello")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	haltOn(err)

	// context := context.Background()
	deadlineInsertTime := 1000 * time.Millisecond
	context, _ := context.WithTimeout(context.Background(), deadlineInsertTime)
	err = client.Connect(context)
	haltOn(err)

	fmt.Println("Connected to the mongoDB succesfully!")

	defer client.Disconnect(context)

	demoDB := client.Database("demo")
	// err = demoDB.CreateCollection(context, "cats")
	// haltOn(err)

	catsCollection := demoDB.Collection("cats")

	fmt.Println("catsCollection  :>> ", catsCollection)
	result, err := catsCollection.InsertOne(context, bson.D{
		{Key: "name", Value: "Thang"},
		{Key: "age", Value: 20},
	})
	haltOn(err)

	fmt.Println("Result :>> ", result)
}
