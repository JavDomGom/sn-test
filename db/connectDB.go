package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	/* MongoCN is the connection object to the database.*/
	MongoCN       = ConnectDB()
	DBname        = "XXX"
	DBuser        = "XXX"
	DBpass        = "XXX"
	clientOptions = options.Client().ApplyURI("mongodb+srv://" + DBuser + ":" + DBpass + "@cluster0.9emsq.mongodb.net/" + DBname + "?retryWrites=true&w=majority")
)

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
