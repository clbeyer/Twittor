package routers

import (
	"encoding/json"
	"net/http"

	"github.com/clbeyer/Twittor/bd"
	"github.com/clbeyer/Twittor/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos:"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password minimo de 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado {
		http.Error(w, "Usuario ya registrado", 400)
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar usuario:"+err.Error(), 400)
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el usuario", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
