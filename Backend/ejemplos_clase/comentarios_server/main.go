package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

/*
vamos a implementar un servidor para una aplicacion ficticia que registra los
comentarios que se hacen en una red social
*/
type Comentario struct {
	Id        int       `json:"id"`
	Timestamp time.Time `json:"time"`
	Comment   string    `json:"comment"`
	Reactions uint      `json:"reactions"`
}

type BaseDatos struct {
	memoria   map[int]Comentario
	proximoId int
}

func (bd *BaseDatos) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		comentarios := []Comentario{}
		for _, comentario := range bd.memoria {
			comentarios = append(comentarios, comentario)
		}
		jsonComent, err := json.Marshal(comentarios)
		if err != nil {
			http.Error(w, "fallo al codificar en json", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonComent)
	} else if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "fallo la peticion POST", http.StatusBadRequest)
		}
		comentario := Comentario{}
		err = json.Unmarshal(body, &comentario)
		if err != nil {
			http.Error(w, "fallo al codificar en json", http.StatusInternalServerError)
		}
		bd.memoria[comentario.Id] = comentario
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	handler := &BaseDatos{
		memoria: map[int]Comentario{
			1: {
				Id:        1,
				Timestamp: time.Now(),
				Comment:   "Quedaste muy bien en la foto",
				Reactions: 4,
			},
			2: {
				Id:        2,
				Timestamp: time.Now(),
				Comment:   "Estabas paseando?",
				Reactions: 1,
			},
		},
	}
	http.Handle("/posts", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
