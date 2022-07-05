package routers

import (
	db "github/twittu/bd"
	"net/http"
)

/*EliminarTweet*/

func EliminarTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	err := db.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un erro al intentar borrar un tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
