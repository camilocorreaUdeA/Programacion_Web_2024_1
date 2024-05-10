package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/camilocorreaUdeA/web_server_2/controllers"
)

type HandlerComentarios struct {
	controller *controllers.Controller
}

func NewHandlerComentarios(controller *controllers.Controller) (*HandlerComentarios, error) {
	if controller == nil {
		return nil, fmt.Errorf("para instanciar un handler se necesita un controlador no nulo")
	}
	return &HandlerComentarios{
		controller: controller,
	}, nil
}

func (hc *HandlerComentarios) ListarComentarios() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		comentarios, err := hc.controller.ListarComentarios()
		if err != nil {
			log.Printf("fallo al leer comentarios, con error: %s", err.Error())
			http.Error(w, "fallo al leer los comentarios", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(comentarios)
	})
}

func (hc *HandlerComentarios) CrearComentario() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("fallo al crear un nuevo comentario, con error: %s", err.Error())
			http.Error(w, "fallo al crear un nuevo comentario", http.StatusBadRequest)
			return
		}

		nuevoId, err := hc.controller.CrearComentario(body)
		if err != nil {
			log.Println("fallo al crear un nuevo comentario, con error:", err.Error())
			http.Error(w, "fallo al crear un nuevo comentario", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, fmt.Sprintf("id del nuevo comentario: %d", nuevoId))
	})
}

func (hc *HandlerComentarios) TraerComentario() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "no se encontro un id valido", http.StatusBadRequest)
			return
		}
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "no se encontro un id valido", http.StatusBadRequest)
			return
		}

		comentario, err := hc.controller.TraerComentario(id)
		if err != nil {
			log.Printf("fallo al leer un comentario, con error: %s", err.Error())
			http.Error(w, "fallo al leer un comentario", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(comentario)
	})
}

func (hc *HandlerComentarios) ActualizarComentario() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "no se encontro un id valido", http.StatusBadRequest)
			return
		}
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "no se encontro un id valido", http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("fallo al actualizar un comentario, con error: %s", err.Error())
			http.Error(w, "fallo al actualizar un comentario", http.StatusInternalServerError)
			return
		}

		err = hc.controller.ActualizarComentario(body, id)
		if err != nil {
			log.Printf("fallo al actualizar un comentario, con error: %s", err.Error())
			http.Error(w, "fallo al actualizar un comentario", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func (hc *HandlerComentarios) EliminarComentario() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "no se encontro un id valido", http.StatusBadRequest)
			return
		}
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "no se encontro un id valido", http.StatusBadRequest)
			return
		}

		err = hc.controller.EliminarComentario(id)
		if err != nil {
			log.Printf("fallo al eliminar un comentario, con error: %s", err.Error())
			http.Error(w, "fallo al eliminar un comentario", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

}
