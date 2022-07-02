package handlers

import (
	"github/twittu/middlew"
	"github/twittu/routers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Handlers and listening to server*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoDB(routers.Login)).Methods("POST")
	/*router.HandleFunc("/verperfil", middlew.ChequeoDB(middlew.ChequeoDB(middlew.ValidoJWT(routers.Verperfil))).Methods("GET")*/

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"

	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
