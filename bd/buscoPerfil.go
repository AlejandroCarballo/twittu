package db

import (
	"context"
	"fmt"
	"github/twittu/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil*/

func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 18*time.Second)
	defer cancel()

	db := MongoCN.Database("twittu")
	col := db.Collection("user")

	var perfil models.Usuario
	ObjID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": ObjID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return perfil, err
	}
	return perfil, nil

}
