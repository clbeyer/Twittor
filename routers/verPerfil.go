package routers

import (
	"encoding/json"
	"net/http"

	"github.com/clbeyer/Twittor/bd"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
