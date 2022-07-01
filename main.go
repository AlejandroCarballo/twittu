package main

import (
	db "github/twittu/bd"
	"github/twittu/handlers"
	"log"
)

func main() {

	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}

	handlers.Handlers()

}
