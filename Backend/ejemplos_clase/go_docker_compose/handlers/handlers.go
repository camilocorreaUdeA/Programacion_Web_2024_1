package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/camilocorreaUdeA/learn_docker_compose/controllers"
	"github.com/gorilla/mux"
)

type Handler struct {
	controller *controllers.Controller
}

func NewHandler(controller *controllers.Controller) (*Handler, error) {
	if controller == nil {
		return nil, fmt.Errorf("para instanciar un handler se necesita un controlador no nulo")
	}
	return &Handler{
		controller: controller,
	}, nil
}

func (h *Handler) ActualizarComentario(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al actualizar un comentario, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un comentario, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	err = h.controller.ActualizarUnComentario(body, id)
	if err != nil {
		log.Printf("fallo al actualizar un comentario, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un comentario, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) EliminarComentario(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	err := h.controller.EliminarUnComentario(id)
	if err != nil {
		log.Printf("fallo al eliminar un comentario, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al eliminar un comentario, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) TraerComentario(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	comentario, err := h.controller.LeerUnComentario(id)
	if err != nil {
		log.Printf("fallo al leer un comentario, con error: %s", err.Error())
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(fmt.Sprintf("el comentario con id %s no se pudo encontrar", id)))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(comentario)
}

func (h *Handler) ListarComentarios(writer http.ResponseWriter, req *http.Request) {
	comentarios, err := h.controller.ListarComentarios(100, 0)
	if err != nil {
		log.Printf("fallo al leer comentarios, con error: %s", err.Error())
		http.Error(writer, "fallo al leer los comentarios", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(comentarios)
}

func (h *Handler) CrearComentario(writer http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al crear un nuevo comentario, con error: %s", err.Error())
		http.Error(writer, "fallo al crear un nuevo comentario", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	nuevoId, err := h.controller.CrearComentario(body)
	if err != nil {
		log.Println("fallo al crear un nuevo comentario, con error:", err.Error())
		http.Error(writer, "fallo al crear un nuevo comentario", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("id nuevo comentario: %d", nuevoId)))
}
