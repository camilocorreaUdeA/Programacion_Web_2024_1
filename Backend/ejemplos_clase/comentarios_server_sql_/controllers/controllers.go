package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/camilocorreaUdeA/web_server_2/models"
	"github.com/camilocorreaUdeA/web_server_2/repository"
)

var (
	listQuery   = "SELECT id, time, comment, reactions FROM comentarios limit $1 offset $2"
	readQuery   = "SELECT id, time, comment, reactions FROM comentarios WHERE id=$1"
	crearQuery  = "INSERT INTO comentarios (time, comment, reactions) VALUES (:time, :comment, :reactions) RETURNING id"
	updateQuery = "UPDATE comentarios SET %s WHERE id=:id"
	deleteQuery = "DELETE FROM comentarios WHERE id=$1"
)

type Controller struct {
	repo repository.Repository[models.Comentario]
}

func NewController(repo repository.Repository[models.Comentario]) (*Controller, error) {
	if repo == nil {
		return nil, fmt.Errorf("para un controlador es necesario un repositorio valido")
	}
	return &Controller{
		repo: repo,
	}, nil
}

func (c *Controller) ListarComentarios() ([]byte, error) {
	comentarios, _, err := c.repo.List(context.Background(), listQuery, 10, 0)
	if err != nil {
		log.Printf("fallo al leer comentarios, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer comentarios, con error: %s", err.Error())
	}

	jsonComentarios, err := json.Marshal(comentarios)
	if err != nil {
		log.Printf("fallo al leer comentarios, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer comentarios, con error: %s", err.Error())
	}
	return jsonComentarios, nil
}

func (c *Controller) TraerComentario(id string) ([]byte, error) {
	comentario, err := c.repo.Read(context.Background(), readQuery, id)
	if err != nil {
		log.Printf("fallo al leer un comentario, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un comentario, con error: %s", err.Error())
	}

	comentarioJson, err := json.Marshal(comentario)
	if err != nil {
		log.Printf("fallo al leer un comentario, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un comentario, con error: %s", err.Error())
	}

	return comentarioJson, nil
}

func (c *Controller) CrearComentario(body []byte) (int64, error) {
	nuevoComentario := &models.Comentario{}
	err := json.Unmarshal(body, nuevoComentario)
	if err != nil {
		log.Printf("fallo al crear un nuevo comentario, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo comentario, con error: %s", err.Error())
	}

	valoresColumnas := map[string]any{
		"time":      nuevoComentario.Timestamp,
		"comment":   nuevoComentario.Comment,
		"reactions": nuevoComentario.Reactions,
	}

	nuevoId, err := c.repo.Create(context.Background(), crearQuery, valoresColumnas)
	if err != nil {
		log.Printf("fallo al crear un nuevo comentario, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo comentario, con error: %s", err.Error())
	}

	return nuevoId, nil
}

func (c *Controller) ActualizarComentario(body []byte, id string) error {
	valoresActualizarBody := make(map[string]any)
	err := json.Unmarshal(body, &valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un comentario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar comentario, con error: %s", err.Error())
	}

	if len(valoresActualizarBody) == 0 {
		log.Printf("fallo al actualizar un comentario, con error: no hay datos en el body")
		return fmt.Errorf("fallo al actualizar comentario, con error: no hay datos en el body")
	}

	updtQuery := fmt.Sprintf(updateQuery, buildUpdateQuery(valoresActualizarBody))
	valoresActualizarBody["id"] = id
	log.Printf("query actualizar: %s", updtQuery)
	err = c.repo.Update(context.Background(), updtQuery, valoresActualizarBody)
	if err != nil {
		log.Printf("fallo al actualizar un comentario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar comentario, con error: %s", err.Error())
	}
	return nil
}

func buildUpdateQuery(columnasActualizar map[string]any) string {
	columnas := []string{}
	for key := range columnasActualizar {
		columnas = append(columnas, fmt.Sprintf("%s=:%s", key, key))
	}
	return strings.Join(columnas, ",")
}

func (c *Controller) EliminarComentario(id string) error {
	err := c.repo.Delete(context.Background(), deleteQuery, id)
	if err != nil {
		log.Printf("fallo al eliminar un comentario, con error: %s", err.Error())
		return fmt.Errorf("fallo al eliminar comentario, con error: %s", err.Error())
	}
	return nil
}
