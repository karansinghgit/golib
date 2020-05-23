package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	handlers "github.com/karansinghgit/golib/handlers"
)

// Connect estabilishes connection with the Database
func Connect() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error while estabilishing connection to database", err)
	}

	log.Println("Connected")

	//create db newMongo
	database := client.Database("books")

	//db instance
	handlers.BookCollection(database)
	return
}
