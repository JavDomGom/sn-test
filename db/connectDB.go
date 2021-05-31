package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	/* MongoCN is the connection object to the database.*/
	MongoCN = ConnectDB()
	DBname  = "XXXXXXXXX"
	DBuser  = "XXXXXXXXX"
	DBpass  = "XXXXXXXXX"
)

/* ConnectDB Function to connect to the database.*/
func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(
			"mongodb+srv://"+DBuser+":"+DBpass+"@cluster0.9emsq.mongodb.net/"+DBname+"?retryWrites=true&w=majority",
		),
	)
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
	return MongoCN.Ping(context.TODO(), nil) == nil
}
