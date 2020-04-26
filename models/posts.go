package models

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

type Result struct {
	ConnString string `json:"connString"`
}

type MongoClient struct {
	client *mongo.Client
	ctx    context.Context
}

var (
	posts      []*Post
	jsonFile   *os.File
	connString string
	err        error
	result     Result
	db         MongoClient
)

func Connect() {
	initializeClient(&db)
	err = db.client.Ping(db.ctx, readpref.Primary())
	if err != nil {
		fmt.Printf("there was an error: %v", err)
	} else {
		fmt.Println("Connection was established")
	}
}
func initializeClient(db *MongoClient) {
	//extract connection string from config.json
	fmt.Println("Initialize struct has been reached")
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)

	//create client
	db.client, err = mongo.NewClient(options.Client().ApplyURI(result.ConnString))
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	} else {
		fmt.Println("Client has been created")
	}

	//define context
	db.ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = db.client.Connect(db.ctx)
	if err != nil {
		fmt.Printf("there was an error: %v", err)
	} else {
		fmt.Println("Context has been defined")
	}

}

func GetPosts() []*Post {
	Connect()
	databases, err := db.client.ListDatabaseNames(db.ctx, bson.M{})
	if err != nil {
		fmt.Printf("there was an error: %v", err)
	}
	fmt.Println(databases)
	collections, err := db.client.Database("db").ListCollectionNames(db.ctx, bson.D{})
	if err != nil {
		fmt.Printf("there was an error: %v", err)
	}
	fmt.Println(collections)
	defer db.client.Disconnect(db.ctx)
	return nil
}
