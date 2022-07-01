package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN object to coneccion to DB*/
var MongoCN = ConectionDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://acarballo:4WrBCnT1WpE8TXN8@twittu.ahajp.mongodb.net/?retryWrites=true&w=majority")

/*ConectionDB is a function to allow connect to DB*/
func ConectionDB() *mongo.Client {
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
	log.Println("conexion success")
	return client

}

/* CheckConnection is a ping to DB*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1

}
