package main

import (
	"log"

	"github.com/clbeyer/Twittor/bd"
	"github.com/clbeyer/Twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
