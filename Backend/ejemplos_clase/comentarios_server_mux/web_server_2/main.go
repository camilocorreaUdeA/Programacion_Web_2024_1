package main

import (
	"log"
	"net/http"

	"github.com/camilocorreaUdeA/web_server_2/handlers"
	"github.com/camilocorreaUdeA/web_server_2/repository"
)

/*
vamos a implementar un servidor para una aplicacion ficticia que registra los
comentarios que se hacen en una red social
*/
func main() {
	bd := repository.NewBaseDatos()
	handler := handlers.NewHandlerComentarios(bd)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /posts", handler.ListarComentarios())
	mux.HandleFunc("POST /posts", handler.CrearComentario())

	mux.HandleFunc("GET /posts/{id}", handler.TraerComentario()) // parametro de ruta o path parameter

	log.Fatal(http.ListenAndServe(":8080", mux))
}
