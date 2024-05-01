package repository

import "github.com/camilocorreaUdeA/web_server_2/models"

type BaseDatos struct {
	Memoria   map[int]models.Comentario
	ProximoId int
}

func NewBaseDatos() *BaseDatos {
	return &BaseDatos{
		Memoria:   make(map[int]models.Comentario),
		ProximoId: 1,
	}
}
