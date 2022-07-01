package routers

import (
	"encoding/json"
	db "github/twittu/bd"
	"github/twittu/models"
	"net/http"
)

/*Registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El mail vino vacio ", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres ", 400)
		return
	}

	_, encontrado, _ := db.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario con ese Email", 400)
		return
	}
	_, status, err := db.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario", 400)
		return

	}

	w.WriteHeader(http.StatusCreated)
}
