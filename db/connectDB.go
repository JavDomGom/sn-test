package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN is the connection object to the database.*/
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://<userName>:<password>@cluster0.9emsq.mongodb.net/<databaseName>?retryWrites=true&w=majority")

/* ConnectDB Function to connect to the database.*/
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Successful connection to the database.")
	return client
}

/* CheckConnection Function to check database connection.*/
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
