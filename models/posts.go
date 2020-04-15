package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Date struct {
	day   int
	month int
	year  int
}

type PostUser struct {
	Firstname string
	Lastname  string
}

type Post struct {
	title string
	body  string
}

var (
	posts      []*Post
	client     *mongo.Client
	connString = "removed for github"
	ctx        context.Context
)

func Connect() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Printf("there was an error: %v", err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("there was an error: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Printf("there was an error: %v", err)
	} else {
		fmt.Println("Connection was established")
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fmt.Printf("there was an error: %v", err)
	}
	fmt.Println(databases)
}

func GetPosts() []*Post {

	collections, err := client.Database("db").ListCollectionNames(ctx, bson.D{})
	if err != nil {
		fmt.Printf("there was an error: %v", err)
	}
	fmt.Println(collections)
	return nil
}
