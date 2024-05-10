package main

import (
	"log"
	"net/http"

	"github.com/camilocorreaUdeA/web_server_2/controllers"
	"github.com/camilocorreaUdeA/web_server_2/handlers"
	"github.com/camilocorreaUdeA/web_server_2/models"
	"github.com/camilocorreaUdeA/web_server_2/repository"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

/*
vamos a implementar un servidor para una aplicacion ficticia que registra los
comentarios que se hacen en una red social
*/
func main() {

	conn, err := ConectarDB("url_de_conexion_a_la_BD_SQL", "postgres")
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
	}

	bd, err := repository.NewRepository[models.Comentario](conn)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de repositorio", err.Error())
	}

	controller, err := controllers.NewController(bd)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de controller", err.Error())
	}

	handler, err := handlers.NewHandlerComentarios(controller)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de handler", err.Error())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /posts", handler.ListarComentarios())
	mux.HandleFunc("POST /posts", handler.CrearComentario())
	mux.HandleFunc("GET /posts/{id}", handler.TraerComentario())        // parametro de ruta o path parameter
	mux.HandleFunc("PATCH /posts/{id}", handler.ActualizarComentario()) // parametro de ruta o path parameter
	mux.HandleFunc("DELETE /posts/{id}", handler.EliminarComentario())  // parametro de ruta o path parameter

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func ConectarDB(url, driver string) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(url)
	db, err := sqlx.Connect(driver, pgUrl) // driver: postgres
	if err != nil {
		log.Printf("fallo la conexion a PostgreSQL, error: %s", err.Error())
		return nil, err
	}

	log.Printf("Nos conectamos bien a la base de datos db: %#v", db)
	return db, nil
}
