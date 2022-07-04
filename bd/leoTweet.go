package db

import (
	"context"
	"github/twittu/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoTweet*/

func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var resultado []*models.DevuelvoTweets

	cond := bson.M{
		"userid": ID,
	}

	opc := options.Find()
	opc.SetLimit(20)
	opc.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opc.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, cond, opc)
	if err != nil {
		log.Fatal(err.Error())
		return resultado, false
	}

	for cursor.Next(context.TODO()) {
		var reg models.DevuelvoTweets

		err := cursor.Decode(&reg)

		if err != nil {
			return resultado, false
		}

		resultado = append(resultado, &reg)
	}
	return resultado, true

}
