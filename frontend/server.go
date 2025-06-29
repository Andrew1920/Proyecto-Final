package main

import (
	"log"
	"net/http"
)

func main() {
	// Sirve los archivos estáticos (html, css, js) del directorio actual.
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	log.Println("Servidor de frontend iniciado en http://localhost:8001")
	// Escucha en un puerto diferente para no chocar con la API.
	http.ListenAndServe(":8001", nil)
}
