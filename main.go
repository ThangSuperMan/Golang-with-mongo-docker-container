package main

import (
	"GoAndMongo/helper"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"html/template"
	"net/http"
	"time"
)

var (
	mongoURI = "mongodb://admin:password@localhost:27017"
)

func renderHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("renderHomePage")
	t, err := template.ParseFiles("index.html")
	helper.HaltOn(err)
	t.Execute(w, nil)
}

func handler() {
	http.HandleFunc("/", renderHomePage)
}

func main() {
	handler()
	port := "3002"
	fmt.Println("Listenning on the port :>> ", port)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	helper.HaltOn(err)

	// context := context.Background()
	deadlineInsertTime := 1000 * time.Millisecond
	context, _ := context.WithTimeout(context.Background(), deadlineInsertTime)
	err = client.Connect(context)
	helper.HaltOn(err)

	fmt.Println("Connected to the mongoDB succesfully!")

	defer client.Disconnect(context)

	demoDB := client.Database("demo")
	// err = demoDB.CreateCollection(context, "cats")
	// helper.HaltOn(err)

	catsCollection := demoDB.Collection("cats")

	fmt.Println("catsCollection  :>> ", catsCollection)
	result, err := catsCollection.InsertOne(context, bson.D{
		{Key: "name", Value: "Bich"},
		{Key: "age", Value: 20},
	})
	helper.HaltOn(err)

	fmt.Println("Result :>> ", result)
	http.ListenAndServe(":"+port, nil)
}
