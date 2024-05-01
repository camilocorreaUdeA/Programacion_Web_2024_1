package models

import "time"

type Comentario struct {
	Id        int       `json:"id"`
	Timestamp time.Time `json:"time"`
	Comment   string    `json:"comment"`
	Reactions uint      `json:"reactions"`
}
