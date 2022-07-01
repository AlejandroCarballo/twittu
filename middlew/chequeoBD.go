package middlew

import (
	db "github/twittu/bd"
	"net/http"
)

/*ChequeoDB es el middlew que me permite conocer el estado de la BD */
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return

		}
		next.ServeHTTP(w, r)

	}
}
