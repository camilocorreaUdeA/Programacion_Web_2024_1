package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type CustomHandler struct {
	numConn int
	msg     string
}

/* Handler implementando el método ServeHTTP de la interfaz http.Handler */
func (ch *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ch.numConn++
	ch.msg = fmt.Sprintf("Se ha conectado al servidor. Conexion numero: %d\n", ch.numConn)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, ch.msg)
}

/* Handler implementado a traves de una funcion http.HandlerFunc */
func (ch *CustomHandler) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ch.numConn++
		ch.msg = fmt.Sprintf("Se ha conectado al servidor. Conexion numero: %d\n", ch.numConn)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, ch.msg)
	}
}

func main() {
	handler := &CustomHandler{}
	/* http.Handle para ejecutar handlers implementados como http.Handler */
	//http.Handle("/hello", handler)
	/* http.HandleFunc para ejecutar handlers implementados como http.HandlerFunc */
	http.HandleFunc("/hello", handler.Handler())
	/* servidor web escuchando en localhost en el puerto 8080 */
	log.Fatal(http.ListenAndServe(":8080", nil))
}
