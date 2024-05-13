package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/camilocorreaUdeA/learn_docker_compose/controllers"
	"github.com/camilocorreaUdeA/learn_docker_compose/handlers"
	"github.com/camilocorreaUdeA/learn_docker_compose/models"
	repositorio "github.com/camilocorreaUdeA/learn_docker_compose/repository" /* importando el paquete de repositorio */
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

/*
función para conectarse a la instancia de PostgreSQL, en general sirve para cualquier base de datos SQL.
Necesita la URL del host donde está instalada la base de datos y el tipo de base datos (driver)
*/
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

func main() {
	/* creando un objeto de conexión a PostgreSQL */
	//db, err := ConectarDB("postgres://frtzcnqy:pYvsWxUKNQhG6xtFFqAj6sdTZdoc0lvB@chunee.db.elephantsql.com/frtzcnqy", "postgres")
	db, err := ConectarDB(fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")), "postgres")
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}

	/* creando una instancia del tipo Repository del paquete repository
	se debe especificar el tipo de struct que va a manejar la base de datos
	para este ejemplo es Amigo y se le pasa como parámetro el objeto de
	conexión a PostgreSQL */
	repo, err := repositorio.NewRepository[models.Comentario](db)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de repositorio", err.Error())
		return
	}

	controller, err := controllers.NewController(repo)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de controller", err.Error())
		return
	}

	handler, err := handlers.NewHandler(controller)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de handler", err.Error())
		return
	}

	/* router (multiplexador) a los endpoints de la API (implementado con el paquete gorilla/mux) */
	router := mux.NewRouter()

	/* rutas a los endpoints de la API */
	router.Handle("/posts", http.HandlerFunc(handler.ListarComentarios)).Methods(http.MethodGet)
	router.Handle("/posts", http.HandlerFunc(handler.CrearComentario)).Methods(http.MethodPost)
	router.Handle("/posts/{id}", http.HandlerFunc(handler.TraerComentario)).Methods(http.MethodGet)
	router.Handle("/posts/{id}", http.HandlerFunc(handler.ActualizarComentario)).Methods(http.MethodPatch)
	router.Handle("/posts/{id}", http.HandlerFunc(handler.EliminarComentario)).Methods(http.MethodDelete)

	/* servidor escuchando en localhost por el puerto 8080 y entrutando las peticiones con el router */
	log.Fatal(http.ListenAndServe(":8080", router))
}
