package db

import (
	"github/twittu/models"

	"golang.org/x/crypto/bcrypt"
)

/*Intento login*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {

	us, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return us, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(us.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return us, false
	}

	return us, true
}
