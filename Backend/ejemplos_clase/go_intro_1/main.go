package main

import (
	format "fmt"

	"github.com/camilocorreaUdeA/go_demo/mipaquete"

	carbon "github.com/golang-module/carbon/v2"
)

var array [5]string

func main() {
	var numero int = 100
	numero *= 2
	Saludo = "hola"
	format.Println("Desarrollo de aplicaciones web")
	format.Printf("%s\n", carbon.Yesterday())
	SayHello()
	format.Println(mipaquete.MyFunction("Juan Camilo"))

	/* Declaración de un puntero */
	format.Println("numero:", numero)
	ptr := &numero
	*ptr = 500
	format.Println("numero (por puntero):", numero)
	array = [5]string{"hello", "world", "dev", "apps", "web"}
	for _, valor := range array {
		format.Println(valor)
	}

	s := []int{1, 2, 3}
	format.Printf("%#v\n", s)
	s = append(s, []int{4, 5, 6}...)
	format.Printf("%#v\n", s)

	var mapa map[int]string
	format.Println("mapa[1]:", mapa[1])
	mapa = make(map[int]string)
	mapa[1] = "mapa en la uno"
	format.Println("mapa[1]:", mapa[1])

	for key, value := range mapa {
		format.Printf("mapa[%d]: %s\n", key, value)
	}

}

func SayHello() {
	format.Println("Hello")
}

var Saludo string //variable global
