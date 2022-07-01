package db

import (
	"context"
	"github/twittu/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 18*time.Second)
	defer cancel()

	db := MongoCN.Database("twittu")
	col := db.Collection("user")

	condicion := bson.M{"email": email}

	var result models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&result)

	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID

}
