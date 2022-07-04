package routers

import (
	"encoding/json"
	db "github/twittu/bd"
	"github/twittu/models"
	"net/http"
	"time"
)

/*GraboTweet*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error para grabar tweet"+err.Error(), 400)
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet", 400)
	}

	w.WriteHeader(http.StatusCreated)

}
