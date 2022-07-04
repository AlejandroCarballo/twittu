package routers

import (
	"encoding/json"
	db "github/twittu/bd"
	"github/twittu/models"
	"net/http"
)

/*ModificarPerfil*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}
	var status bool
	status, err = db.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrión un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se pudo modificar el registro del usuario"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
