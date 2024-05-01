package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/camilocorreaUdeA/web_server_2/models"
	"github.com/camilocorreaUdeA/web_server_2/repository"
)

type HandlerComentarios struct {
	BD *repository.BaseDatos
}

func NewHandlerComentarios(bd *repository.BaseDatos) *HandlerComentarios {
	return &HandlerComentarios{
		BD: bd,
	}
}

func (hc *HandlerComentarios) ListarComentarios() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		comentarios := []models.Comentario{}
		for _, comentario := range hc.BD.Memoria {
			comentarios = append(comentarios, comentario)
		}
		jsonComent, err := json.Marshal(comentarios)
		if err != nil {
			http.Error(w, "fallo al codificar en json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonComent)
	})
}

func (hc *HandlerComentarios) CrearComentario() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "fallo la peticion POST", http.StatusBadRequest)
		}
		comentario := models.Comentario{}
		err = json.Unmarshal(body, &comentario)
		if err != nil {
			http.Error(w, "fallo al codificar en json", http.StatusInternalServerError)
		}
		comentario.Id = hc.BD.ProximoId
		hc.BD.Memoria[hc.BD.ProximoId] = comentario
		hc.BD.ProximoId++
		w.WriteHeader(http.StatusCreated)
	})
}

func (hc *HandlerComentarios) TraerComentario() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "no se encontro un id valido", http.StatusBadRequest)
			return
		}
		idNum, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "no se encontro un id valido", http.StatusBadRequest)
			return
		}

		comentario, ok := hc.BD.Memoria[idNum]
		if !ok {
			http.Error(w, "no se encontro un comentario para ese id", http.StatusNotFound)
			return
		}

		payload, err := json.Marshal(comentario)
		if err != nil {
			http.Error(w, "fallo la codificacion a JSON del comentario", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	})
}
