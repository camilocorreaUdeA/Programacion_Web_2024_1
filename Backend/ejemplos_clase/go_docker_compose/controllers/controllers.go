package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/camilocorreaUdeA/learn_docker_compose/models"
	repositorio "github.com/camilocorreaUdeA/learn_docker_compose/repository"
)

var (
	updateQuery = "UPDATE comentarios SET %s WHERE id=:id;"
	deleteQuery = "DELETE FROM comentarios WHERE id=$1;"
	selectQuery = "SELECT id, time, comment, reactions FROM comentarios WHERE id=$1;"
	listQuery   = "SELECT id, time, comment, reactions FROM comentarios limit $1 offset $2"
	createQuery = "INSERT INTO comentarios (time, comment, reactions) VALUES (:time, :comment, :reactions) returning id;"
)

type Controller struct {
	repo repositorio.Repository[models.Comentario]
}

func NewController(repo repositorio.Repository[models.Comentario]) (*Controller, error) {
	if repo == nil {
		return nil, fmt.Errorf("para instanciar un controlador se necesita un repositorio no nulo")
	}
	return &Controller{
		repo: repo,
	}, nil
}

func (c *Controller) ActualizarUnComentario(reqBody []byte, id string) error {
	nuevosValoresComentario := make(map[string]any)
	err := json.Unmarshal(reqBody, &nuevosValoresComentario)
	if err != nil {
		log.Printf("fallo al actualizar un comentario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un comentario, con error: %s", err.Error())
	}

	if len(nuevosValoresComentario) == 0 {
		log.Printf("fallo al actualizar un comentario, con error: no hay datos")
		return fmt.Errorf("fallo al actualizar un comentario, con error: no hay datos")
	}

	query := construirUpdateQuery(nuevosValoresComentario)
	nuevosValoresComentario["id"] = id
	err = c.repo.Update(context.TODO(), query, nuevosValoresComentario)
	if err != nil {
		log.Printf("fallo al actualizar un comentario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un comentario, con error: %s", err.Error())
	}
	return nil
}

func construirUpdateQuery(nuevosValores map[string]any) string {
	columnas := []string{}
	for key := range nuevosValores {
		columnas = append(columnas, fmt.Sprintf("%s=:%s", key, key))
	}
	columnasString := strings.Join(columnas, ",")
	return fmt.Sprintf(updateQuery, columnasString)
}

func (c *Controller) EliminarUnComentario(id string) error {
	err := c.repo.Delete(context.TODO(), deleteQuery, id)
	if err != nil {
		log.Printf("fallo al eliminar un comentario, con error: %s", err.Error())
		return fmt.Errorf("fallo al eliminar un comentario, con error: %s", err.Error())
	}
	return nil
}

func (c *Controller) LeerUnComentario(id string) ([]byte, error) {
	comentario, err := c.repo.Read(context.TODO(), selectQuery, id)
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

func (c *Controller) ListarComentarios(limit, offset int) ([]byte, error) {
	comentarios, _, err := c.repo.List(context.TODO(), listQuery, limit, offset)
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

func (c *Controller) CrearComentario(reqBody []byte) (int64, error) {
	nuevoComentario := &models.Comentario{}
	err := json.Unmarshal(reqBody, nuevoComentario)
	if err != nil {
		log.Printf("fallo al crear un nuevo comentario, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo comentario, con error: %s", err.Error())
	}

	valoresColumnasNuevoComentario := map[string]any{
		"time":      nuevoComentario.Timestamp,
		"comment":   nuevoComentario.Comment,
		"reactions": nuevoComentario.Reactions,
	}

	nuevoId, err := c.repo.Create(context.TODO(), createQuery, valoresColumnasNuevoComentario)
	if err != nil {
		log.Printf("fallo al crear un nuevo comentario, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo comentario, con error: %s", err.Error())
	}
	return nuevoId, nil
}
