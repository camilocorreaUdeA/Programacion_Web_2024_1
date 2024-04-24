package main

import "fmt"

/* Ejemplo:
Haga una funcion que determina la interseccion entre dos slices
La funcion debe retornar un slice con los elementos comunes a los dos
slices y las veces que se repiten en ambos

Ejm: x = {23, 3, 1, 2}, y = {6, 2, 4, 23}
Resultado = {2, 23}

Ejem: x = {1, 1, 1}, y = {1, 1, 1, 1}
Resultado = {1, 1, 1}
*/

func interseccion(x, y []int) []int {
	mapaDeX, mapaDeY := make(map[int]int), make(map[int]int)
	resultado := []int{}

	for _, num := range x {
		veces, ok := mapaDeX[num]
		if ok {
			mapaDeX[num] = veces + 1
			continue
		}
		mapaDeX[num] = 1
	}

	for _, num := range y {
		veces, ok := mapaDeY[num]
		if ok {
			mapaDeY[num] = veces + 1
			continue
		}
		mapaDeY[num] = 1
	}

	for c, v := range mapaDeX {
		veces, ok := mapaDeY[c]
		if ok {
			minimo := min(v, veces)
			for _ = range minimo {
				resultado = append(resultado, c)
			}
		}
	}
	return resultado
}

func main() {
	x := []int{23, 3, 1, 2}
	y := []int{6, 2, 4, 23}

	res := interseccion(x, y)
	fmt.Printf("%#v\n", res)

	a := []int{1, 1, 1}
	b := []int{1, 1, 1, 1}
	res = interseccion(a, b)
	fmt.Printf("%#v\n", res)
	/*slice := []string{"hola", "mundo"}
	slice[1] = "bye"
	slice = append(slice, "web")

	for i, elemento := range slice {
		fmt.Printf("slice en la pos: %d es %s\n", i, elemento)
	}

	ss := make([]int, 0, 10)
	for _, elemento := range ss {
		fmt.Println(elemento)
	}

	mapa := make(map[int]string) // mapa != nil
	mapa[1] = "uno"
	mapa[2] = "dos"
	mapa[3] = "tres"

	for clave := range mapa {
		//for clave, valor := range mapa {
		fmt.Println("clave:", clave, "valor:", mapa[clave])
		//fmt.Println("clave:", clave, "valor:", valor)
	}*/

	estudiante := &Estudiante{
		Nombre:      "Laura",
		Edad:        20,
		Universidad: "UdeA",
	}

	/*fmt.Println("Nombre estudiante:", estudiante.Nombre)
	fmt.Println("Edad estudiante:", estudiante.Edad)
	fmt.Println("Universidad estudiante:", estudiante.Universidad)

	estudiante.CambiarNombreEstudiante("Carolina")
	fmt.Println("Nombre estudiante:", estudiante.Nombre)
	estudiante.CambiarEdadEstudiante(35)
	fmt.Println("Edad estudiante:", estudiante.Edad)*/

	DecirHola(estudiante)
	arbol := Arbol{}
	DecirHola(arbol)

	profe := &Profesional{
		Estudiante: Estudiante{
			Nombre:      "Emiliano",
			Edad:        45,
			Universidad: "UdeM",
		},
	}

	fmt.Println(profe.GetNombreProfesional())
	profe.CambiarNombreEstudiante("Pedro")
	fmt.Println(profe.GetNombreProfesional())

	resultado, err := AdivinaElNumero(10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resultado)

	_, err = AdivinaElNumero(1)
	if err != nil {
		fmt.Println(err.Error())
	}
}

type Estudiante struct {
	Nombre      string
	Edad        uint
	Universidad string
}

type Profesional struct {
	Estudiante
}

type Arbol struct{}

func (p Profesional) GetNombreProfesional() string {
	return p.Nombre
}

func (e *Estudiante) CambiarNombreEstudiante(nombre string) {
	e.Nombre = nombre
}

func (e *Estudiante) CambiarEdadEstudiante(edad uint) {
	e.Edad = edad
}

func (e Estudiante) VerEdadEstudiante() uint {
	return e.Edad
}

type Interfaz interface {
	SayHello(nombre string) string
}

func (e Estudiante) SayHello(nombre string) string {
	return fmt.Sprintf("Hello %s\n", nombre)
}

func (a Arbol) SayHello(nombre string) string {
	return fmt.Sprintf("Hola %s\n", nombre)
}

func DecirHola(obj Interfaz) {
	fmt.Println(obj.SayHello("Luis"))
}

func AdivinaElNumero(numero int) (string, error) {
	if numero == 10 {
		return "ganaste", nil
	}
	return "", fmt.Errorf("error: perdiste, ese no es el numero")
}
