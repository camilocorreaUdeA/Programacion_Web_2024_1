# Go
![2](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/e131e81c-9a0d-4db2-8d40-99119ebe006e)
## El lenguaje de programación Go (Golang)

Go inició como un proyecto de bajo perfil que rapidamente escalo hasta convertirse en un proyecto ambicioso patrocinado por Google. A finales de 2011 se liberó Go como un proyecto de código abierto y luego el 28 de Marzo de 2012 se lanzó la primera versión estable del lenguaje y a partir de ese momento se procura liberar una nueva versión cada medio año.

"...El objetivo del proyecto Go era eliminar la lentitud y la falta de coordinación en el proceso de desarrollo de software en Google, y por tanto hacer el proceso más productivo y escalable. El lenguaje fue diseñado por y para gente que escribe, lee, depura y mantiene grandes proyectos de software..." -Rob Pike-

Al crear el lenguaje de programación Go la idea principal era combinar las mejores características que están presentes en otros lenguajes, entre ellas:
<ul>
  <li>Fácil de integrar con las herramientas actuales para disparar la productividad</li>
  <li>Alta eficiencia y tipado estático</li>
  <li>Alto desempeño para aplicaciones de networking y un aprovechamiento total de las arquitecturas multi-core</li>
</ul>

Go es un lenguaje especialmente apto para las siguientes aplicaciones y sistemas:
<ul>
  <li>Servicios distribuidos interconectados: Las aplicaciones de red dependen fuertemente de la concurrencia, la cual es manejada de forma simple y eficiente por característcas nativas del lenguaje como los canales y las go-rutinas</li>
  <li>Desarrollo nativo en la nube: Las apliaciones nativas de la nube son más fáciles de construir con las características que Go ofrece para concurrencia y networking, además de la portabilidad entre sistemas operativos y arquitecturas que ofrece de forma nativa el lenguaje.</li>
  <li>Remplazo para infraestructura existente: Go es un lenguaje ideal para re-escribir aplicaciones, permite mejorar la seguridad en el manejo de memoria, permite el despliegue multi-plataforma y una base de código limpia que es de más fácil mantenimiento en el futuro.</li>
  <li>Utilidades y herramientas independientes: Los programas en Go complilan en binarios que practicamente no requieren de dependencias externas. Esto hace a Go ideal para la creación de herramientas y utilidades que se pueden poner en producción con prontitud y que pueden ser empaquetadas para su redistribución.</li>
</ul>

Tal vez Go no es lo sufiecientemente apto en los siguientes casos:
<ul>
  <li>Interfaces gráficas o GUIs (aunque hay paquetes como Fyne y Wails que tienen un desempeño de aceptable a bueno)</li>
  <li>Desarrollo de sistemas embebidos y desarrollo de bajo nivel en general: Existe un branch de Go conocido como TinyGO que permite programar dispositivos de hardware. Go no está diseñado para la creación de drivers o interfaces de propósito general de entrada/salida de datos (comunicaión con periféricos y hardware externo)</li>
</ul>

Ventajas de Go con respecto a otros lenguajes de programación:
<ul>
<li>Fácil de usar y de leer: La sintaxis de Go es simple, entre 25 y 35 palabras clave conforman el lenguaje. Y además su curva de aprendizaje lo hace bastante accesible a programadores novatos.</li>
<li>Amplia librería estándar: Una amplia y variada librería estándar viene empaquetada en el lenguaje, lo cual evita el uso de librerías de terceros para similares propósitos. La librería es sofisticada pero no compleja ni complicada, y el manejo de dependencias de Go evita riesgo de conflictos entre nombres de fiunciones.</li>
<li>Seguridad: Gracias a lo simple del código, a que es un lenguaje de tipado estático y al garbage collector incluido en el runtime del lenguaje.</li>
<li>Soporte: Además del soporte de ingenieros de Google cuenta con una siempre creciente coumunidad.</li>
<li>Concurrencia y paralelismo: Incluidos en el lenguaje en librerías con funciones para estos propósitos.</li>
<li>Manejo de dependencias: El sistema de manejo de dependencias está incluido en el lenguaje facilitando la inclusión de paquetes de terceros y el posterior empaquetamiento en el binario de la aplicación.</li>  
</ul>

Utilidades de Go para los desarrolladores:
<ul>
  <li>Generación automática de documentación del código.</li>
  <li>Herramientas de testing y benchmarking nativas incluídas en el lenguaje.</li>
  <li>Detección de condiciones de carrera para apliaciones concurrentes.</li>
  <li>Un sistema de tipos liviano pero suficiente para clasificar y diferenciar datos y objetos.</li>
  <li>Concurrencia nativa para la ejecución de tareas en segundo plano.</li>
  <li>Garbage collector para el uso racional de memoria.</li>
  <li>Compilación multiplaforma que permite desplegar en cualquier arquitectura y sistema operativo.</li>
  <li>Buen desempeño y velocidad en comparación con otros lenguajes de programación similares.</li>
  <li>Amplia librería estándar, especializada en funciones de concurrencia y networking.</li>
  <li>Punteros para manejo eficiente de memoria.</li>
</ul>

## Instalación de Go

Puedes revisar el [Getting started](https://go.dev/doc/tutorial/getting-started) oficial.
También puedes seguir las instrucciones oficiales de [instalación](https://go.dev/doc/install)

En general el proceso de instalación de Go es bastante intuitivo:
<ul>
  <li>Instala el IDE de tu preferencia, recomiendo VSCode pero Goland también está bien.</li>
  <li>Descarga Go en https://go.dev/dl/. En ese sitio están los instaladores para los distintos sistemas operativos.</li>
  <li>Instalación manual de Go en tu máquina:
    <ul>
      <li>Linux: Ejecute en la terminal: <code>rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz</code>, luego setea la variable de ambiente <code>PATH</code> para que también apunte al directorio bin de Go: <code>export PATH=$PATH:/usr/local/go/bin</code></li>
      <li>Con MacOS y Windows descargue el instalador correspondiente, ejecutelo y verifique la instalación con el comando <code>go version</code> en la terminal.</li>
    </ul>
  </li>
  <li>Otras formas de instalar Go:
  <ul>
    <li>MacOS: Instale <b>homebrew</b> en su sistema y una vez instalado ejecute en la terminal el comando <code>brew install go</code></li>
    <li>Windows: Instale <b>chocolatey</b> y una vez instalado ejecute en la terminal el comando <code>choco install golang</code></li>
  </ul>
    </li>
  <li>Verifique la instalación ejecutando el comando <code>go version</code></li>
</ul>

Una vez hayas instalado Go en tu máquina y si ya tienes VSCode instalado también (¡si no, entonces instalalo!). Abre VSCode y luego presiona la combinación de teclas <code>Ctrl + Shif + P</code> y en el menú que se despliega escoge la opción <code>Go: Install/Update Tools</code>. Espere que se complete la instalación de las herramientas.

<img width="603" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/8a9b7c5d-b8b5-4dd4-82a0-c7926d2bd2c0">

Luego instale la extensión de Go para facilitar la codificación ya que trae funciones como autocompletado, indentación automática y verificación de errores de sintaxis, etc.

<img width="655" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/aa49bf57-b453-45e1-a61f-adb722405ca3">

## Go modules

<i>Go modules</i> es el sistema de administración de dependencias en Go, el cual facilita el acceso a la información de versiones y la inclusión de paquetes externos en una aplicación de Go.

Un modulo en Go es una colección de paquetes almacenados en un árbol de archivos (directorio) donde un archivo llamado <i>go.mod</i> es el nodo raíz del árbol.

El archivo <i>go.mod</i> define la ruta del modulo que se usa como la ruta de <i>import</i> del directorio raíz. También indica la versión de Go usada para compilar el proyecto y el conjunto de dependencias (otros modulos) que se importan y se utilizan en el proyecto.

## Creación de un proyecto en Go

Para crear un proyecto de Go:
<ol>
  <li>Crear un directorio para alojar el proyecto.</li>
  <li>Acceda al directorio con VSCode.</li>
  <li>Abra una Terminal (Terminal -> New Terminal)</li>
  <li>En la terminal ejecute el siguiente comando <code>go mod init &lt;<i>ruta-al-modulo</i>&gt;</code> Donde <i>ruta-al-modulo</i> es el nombre del modulo del proyecto, por convención se acostumbra que esta ruta sea la misma URL al repositorio en github.com donde se alojará el proyecto, por ejemplo: <i>github.com/usuario123/repo-proyecto</i></li>

  ![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/3cd3a8b7-cafd-431f-a3f9-7dbebdc04846)
  
  <li>Verifique que un archivo <i>go.mod</i> haya sido agregado al directorio raíz del proyecto.</li>  
</ol>

## Anatomía básica de un archivo fuente de Go (archivos con extensión <i>.go</i>)

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/196c1dd0-032a-4e7b-9cda-def7e0ca9d15)

<ul>
  <li>Línea package &lt;<i>nombre-package</i>&gt;: Esta línea es la única parte obligatoria de la anatomía y debe ir en todos los archivos fuente de una aplicación de Go, ella indica a que paquete (package) pertenece el archivo fuente. Si el archivo fuente contiene una función <i>main</i> (<i>entry point</i> a la aplicación) entonces el nombre del paquete debe ser <i>main</i>, y por tanto la línea debe quedar <code>package main</code></li>
  <li>Conjunto de <i>imports</i>: En esta sección se hace referencia a los paquetes de los cuales se importan variables, tipos o funciones en el archivo de fuente actual. Se puede importar tanto paquetes locales (es decir, paquetes que están dentro del mismo proyecto) o paquetes externos (son paquetes creados por terceros que se agregan como modulos al proyecto) que se descargan direactmente del repositorio donde están alojados. La sintaxis incluye la palabra <i>import</i> seguida por el conjunto de los nombres y/o las rutas a los paquetes que se van a importar, encerrados entre paréntesis. Dichos paréntesis no son necesarios cuando es un solo paquete el que se importa.</li>
  <li>Declaraciones de variables, constantes y tipos: Por convención se sugiere hacer declaraciones de variables y constantes y definiciones de tipos justo después de la sección de imports y antes de la definición de cualquier función.</li>
  <li>Definición de funciones (lógica): A continuación de las declaraciones de variables y constantes y de las definiciones de tipos se codifican las funciones que hacen parte del archivo fuente. Si el paquete es main, entonces se sugiere que la función main se defina antes de cualquier otra función.</li>
</ul>

## Jerarquía dentro de un módulo en Go

Hasta el momento hemos hablado de los conceptos de módulo, paquete y archivo fuente pero sin dar una definción precisa de cada uno y sin explicar cuál sería la relación entre estos conceptos, pues bien, a continuación dejamos las definiciones pertinentes:
<ul>
  <li><b>Archivo fuente:</b> Es un archivo amoldado por el patrón de anatomía básica que vimos anteriormente. Todos los archivos fuente en Go tienen la extensión .go y son simplemente el lugar en el que se agrupan un conjunto de funciones, declaraciones de variables y definiciones de tipos que están inter-relacionados y que tienen el objetivo de implementar unas funcionalidades que componen una lógica de negocio o algoritmo en específico.</li>
  <li><b>Paquete (package):</b> Un paquete es una agrupación de archivos fuente dentro de un mismo directorio, dichos archivos deben compartir el mismo nombre de paquete en la línea package. Un paquete permite exportar a otros paquetes (o módulos) las variables, constantes, tipos y funciones definidos en los archivos fuente que lo conforman.</li>
  <li><b>Módulo:</b> Es un conjunto o colección formado por uno o más paquetes almacenados en un árbol de archivos y en cuyo directorio raíz existe un archivo go.mod que contiene el nombre del módulo (ruta al repositorio), la versión del lenguaje con la que fue creado y las posibles dependencias que pueda tener con otros módulos externos. Un módulo en Go es en cierto grado lo que comunmente se denomina librería o biblioteca en otros lenguajes de programación.</li>
</ul>

## Tipos de datos en Go, declaración de variables, constantes, funciones y definición de tipos

### Tipos nativos de Go

Numéricos

```go
/* enteros */
uint8 (uint, byte), uint16, uint32, uint64
int8 (int, rune), int16, int32, int64

/* decimales */
float32, float64

/* complejos */
complex64, complex128
```
Booleano

```go
bool
```
Texto y cadenas de caracteres

```go
string
```
Punteros: Los punteros son variables que permiten almacenar direcciones de memoria o referencias de otras variables (no punteros) del mismo tipo. Los punteros cuentan con los operadores <code>*</code> y <code>&</code> para averiguar por el valor almacenado en la referencia y para acceder a una referencia.

```go
/* puntero a una variable de tipo uint64 */
var ptr *uint64
```

Colecciones de datos

Arrays: Son colecciones de datos de un mismo tipo que se almacenan de forma consecutiva en memoria y para las que se define el tamaño en el momento mismo de la declaración. Un tipo array está compuesto tanto por el tipo de dato de los elementos como por la longitud de arreglo mismo. Un array es estático en memoria, esto quiere decir que no pueden cambiar de tamaño en tiempo de ejecución, conservan el tamaño asignado en la declaración durante el ciclo de vida de la aplicación.

```go
/* arreglo de 10 elementos de tipo int */
var array [10]int
```
Slices: Son colecciones similares a los arrays pero con la diferencia de que son dinámicos en memoria, es decir, pueden cambiar de tamaño en tiempo de ejecución, se les puede retirar o agregar elementos. Hay que tener cuidado cuando al agregar elementos se sobrepasa la capacidad del slice ya que esto implica una relocalización del slice en memoria, por tanto invalidando los punteros y copias del slice.

```go
/* slice */
var slice []string
```
Mapas: Un mapa en Go es una colección no ordenada de elementos de cierto tipo a los que se puede acceder o indexar mediante un conjunto de claves que también son de algún tipo. Son similares a los mapas o diccionarios de otros lenguajes de programación, es decir, son colecciones de pares clave-valor, donde el par puede ser del mismo tipo o de tipos distintos, y una clave es única e irrepetible en el mapa. Las claves pueden ser de cualquier tipo a excepción de los tipos array, slice, mapa o función.

```go
/* mapa cuyas claves son de tipo string y sus valores de tipo int */
var mapa map[string]int
```
Funciones, interfaces y estructuras

Funciones: Las funciones en Go son tratadas como tipos de datos. Un tipo función describe el conjunto de funciones que comparten el mismo conjunto de parámetros de entrada (cantidad y tipos de datos) y tipos de retorno.

```go
/* variable de tipo funcion para almacenar funciones que reciben un dato string y retornan un slice de int64 y un error */
var fun func(string)([]int64, error)
```
Interfaces: Las interfaces son contratos que definen un conjunto de métodos (funciones) que deden ser implementados por cualquier tipo que quiera cumplir el contrato de la interfaz (implementar la interfaz). Un tipo que implementa una interfaz automáticamente adquiere el tipo definido por la interfaz. Por lo tanto, una variable del tipo de una interfaz es capaz de almacenar un valor de cualquier otro tipo que implemente la interfaz (que implemente los métodos definidos en la interfaz). Todos los tipos en Go implementan una interfaz en común que es la interfaz vacía (empty interface) definida como <code>interface{}</code> y también con el alias <code>any</code>.

```go
/* una interfaz que define unos métodos para ser implementados por otros tipos */
type MiInterfaz interface {
   MetodoUno()
   MetodoDos(string)error
}
```
Estructuras: Las estructuras son secuencias de datos de distintos tipos (inclusive de otros tipos de estructuras o de interfaces), llamados campos de la estructura. Cada campo tiene un nombre y un tipo. Las estructuras permiten que se les asocien funciones, conocidas como métodos de la estructura, lo cual las hace similares a las clases de otros lenguajes de programación. Se puede acceder a los campos y métodos de una estructura mediante una variable o instancia del tipo de la estructura. Se puede decir que las estructuras son tipos personalizados por el usuario en Go.

```go
/* una estructura en Go con sus campos */
type MiEstructura struct {
   Nombre          string
   Edad            uint
   Email           string
   Caracteristicas []string
}
```
### Declaración de variables en Go

Variables globales de un paquete

Las variables globales de un paquete son aquellas visibles desde cualquier función definida en el paquete. Siempre se declaran por fuera de cualquier función y en la sección del ar archivo fuente dedicado a la declaración de variables y definición de tipos y funciones.

Las variables globales se declaran siempre con la palabra clave <code>var</code> seguida del nombre de la variable y opcionalmentedel tipo de la variable o de una inicialización dejando al compilador la tarea de deducir el tipo.

```go
/* distintas formas de declarar variables globales en Go */
var cadenaDeTexto string /* indicando el tipo y sin inicialización explícita, el valor inicial de cadenaDeTexto será "" */
var contadorDeEventos uint = 10 /* indicando el tipo e inicializando de manera explícita */
var dataSource = "digital" /* sin indicar el tipo y dejando su deducción al compilador a partir del valor asignado */
var x, y, z float32 /* declarando un conjunto de variables del mismo tipo */
var a, b, c = false, 100, "hello"  /* declarando un conjunto de variables de distinto tipo */
/* declarando un conjunto de variables de distintos tipos */
var (
  isSafe = true
  numArray []float64
  times int32 = 1
)
```
Variables locales de una función

Las variables locales son aquellas que se declaran al interior de las funciones y solo tienen validez dentro de la función en la que se declaran (scope). Se pueden declarar de la misma manera que se declaran las variables globales y adicionalmente se pueden declarar usando el operador walrus dejando que el compilador deduzca el tipo.

```go
func myFunc() {
  myLocalVar := "una_var_local"
}
```
Constantes

Las constantes son literales asociados a un valor que se conserva durante toda la ejecución de la aplicación. Las constantes pueden ser globales o locales según se requiera. Y para su declaración aplican las mismas reglas que para la declaración de variables globales y locales, a excepción de que se usa siempre la palabra clave <code>const</code>, siempre se deben inicializar en el momento de su declaración y no se pueden declarar utilizando el operador <code>:=</code> (operador walrus).

Nota: No se pueden declarar constantes de los tipos array, slice, map, struct ni interface.

```go
/* distintas formas de declarar constantes en Go */
const dataCode uint = 111 /* indicando el tipo */
const newLabelCode = "tag_label"  /* permitiendo deducción del tipo de la constante */
const x, y, z = 1, 2, 3 /* declarando un conjunto de constantes del mismo tipo */
/* declarando un conjunto de constantes de distintos tipos */
const (
  isAnalog = false
  numPossibleDataEvents = 15
  DigitalDatasource = "generic_digital"
)
```
Declaración de punteros

Para declarar un puntero se utiliza el operador <code>*</code> y para almacenar una referencia en el puntero se utiliza el operador <code>&</code>. No se pueden declarar punteros constantes ni punteros a constantes. La aritmética de punteros no existe en Go. Pueden declararse tanto punteros globales como locales y por lo tanto también aplica el operador walrus y la deducción de tipo para punteros.

Nota: Existe el operador <code>new</code> para inicializar un puntero al momento de su declaración. Inicializar se entiende como asignar un valor distinto a nulo (<code>nil</code>) al puntero.

```go
/* distintas formas de declarar punteros en Go */
var ptrOne *int /* puntero a variables de tipo int */
var ptrTwo = &someVariable  /* deducción del tipo puntero por inicialización */
var ptrThree = new(MyStruct) /* inicialización del puntero a instancias del tipo MyStruct con el operador new */

func myFunc() {
  myLocalVar := 50
  myLocalPtr := &myLocalVar /* deducción del tipo de puntero a *int */
}
```
Declaración de arrays, slices y maps

Arrays: Los arreglos o arrays de declaran con el operador corchetes <code>[]</code>, indicando allí la longitud del arreglo. Los elementos del arreglo se pueden modificar indexandolos con el operador <code>[]</code> pero no se pueden elmininar ni agregar elementos al arreglo. Con el operador <code>len()</code> se puede averiguar la longitud del arreglo.

Nota: Puede omitir indicar la longitud del arreglo utilizando el operador elipsis <code>...</code>, aunque solo es válido cuando inicializa el arreglo directamente en la declaración.

```go
/* distintas formas de declarar arreglos en Go */
var arrayOne [10]uint64
var arrayTwo = [2]string{"hello", "world"}

func myFunc() {
  myLocalArray := [...]int{-2, -1, 0, 2, 4}
  otherArray := [4]string{}
}
```
Slices: Los slices se pueden pensar como arreglos dinámicos, es decir que pueden cambiar de longitud en tiempo de ejecución, o sea que se pueden eliminar o agregar elementos a conveniencia. El tipo slice está compuesto por un arreglo interno, un puntero al arreglo interno y la longitud y la capacidad del arreglo interno. La longitud es la cantidad de elementos en el arreglo interno, mientras que la capacidad es la cantidad de elementos que puede almacenar el slice antes de tener que hacer una nueva localización de memoria (esto generalmente pasa cuando se van agregando nuevos elementos al slice y se supera su capacidad en el momento).

Los slices se declaran de forma similar a los arreglos pero no se debe indicar el tamaño en la declaración, esa característica es la que le indicará al compilador que la variable es del tipo slice y no del tipo array.

Puede declarar un slice como una porción de un arreglo existente, pero debe tener cuidado porque al modificar el slice también estará modificando el arreglo original.

Puede utilizar la función <code>make</code> para declarar slices, debe indicar como mínimo el tipo de los elementos del slice y la longitud. Opcionalmente puede indicar la capacidad inicial del slice.

También puede utilizar la función <code>len()</code> para averiguar la longitud del slice en determinado momento, y la función <code>cap()</code> para averiguar la capacidad.

Los slices cuentan con la función <code>append()</code> que permite agregar nuevos elementos al final del slice, por lo tanto incrementando la longitud del slice y consumiendo la capacidad. Una vez se supera la capacidad disponible se ejecuta una nueva localización de memoria para aumentar la capacidad del slice.

```go
/* distintas formas de declarar slices en Go */
var array = [4]string{"hello", "world", "go", "golang"}
var slice = array[1:3] /* declarando un slice a partir de un arreglo existente */
var otroSlice = []int{} /* declarando un slice vacío */
var otroVacio []int
var otroMas = []string{"word", "palabra"} /* inicializando en la declaración */
var ss = make([]int64, 10) /* slice de elementos de tipo int64 y longitud y capacidad iniciales 10 */
var slicie = make([]string, 0, 15) /* slice de elementos de tipo string, longitud 0 y capacidad inicial 15 */

func myFunc() {
  myLocalSlice := make([]MyStruct, len(datos)) /* suponga que datos es un slice que ya existe */
}
```
Maps:

Los mapas son colecciones de pares clave-valor donde la clave y el valor pueden ser de distintos tipos. Se pueden declarar utilizando la función <code>make</code> indicando el tipo del mapa (el tipo de un map está compuesto por el tipo de la clave y el tipo del valor) y una capacidad opcional si de antemano se sabe cuántos elementos va a almacenar el mapa.

Un map está constituido internamente por una función hash, una instancia de una estructura de datos hash y un puntero a esa instancia, por lo tanto siempre que vaya agregar elementos a un map por primera vez, debe asegurarse de que esté ha sido inicilizado, de lo contrario ocurrirá un error en tiempo de compilación conocido como <code>panic</code> que indica que está tratando de agregar elementos a un map nulo.

Los mapas cuentan con las funciones <code>len()</code> para saber cuántos elementos hay en el map, y <code>delete()</code> para eliminar un par clave-valor del mapa.

```go
/* distintas formas de declarar maps en Go */
var m1 map[string]int /* mapa de claves de tipo string y valores de tipo int, no inicializado */
var m2 = map[int]bool{} /* mapa de claves de tipo int y valores de tipo bool, inicializado */
var m3 = make(map[string][]string) /* mapa de claves de tipo string y valores de tipo []string (slice de strings), inicializado */

func myFunc() {
  myLocalMap := make(map[string]int64) /* mapa local inicializado */
}
```
Definición e instanciación de structs

Las estructuras o structs son los tipos de datos personalizados por el usuario en Go, son agrupaciones de variables de distinto tipo que se conocen como campos de la estructura. Además, como veremos más adelante, puede asociarse métodos a las estructuras. Luego, solo vamos a poder acceder a los campos y métodos de una estructura a través de una variable o instancia del tipo definido por la estructura.

```go
type Estructura struct {
  CampoUno string
  CampoDos int
  CampoTres []uint64
}

var obj Estructura /* Los campos de obj se inicializan con los zero-values de los tipos correspondientes */
var myObj = Estructura{"hello", 20, []uint64{1, 2, 3}} /* Deducción del tipo, e inicialización de los campos ignorando los nombres */

func myFunc() {
  myLocalPtr := &Estructura{CampoUno: "name"} /* deducción del tipo puntero a Estructura e inicializción selectiva de campos (obliga a especificar nombres) */
}
```
Definición de interfaces

Una interfaz es un contrato para implementar un tipo, el tipo de la interfaz. La interfaz es un conjunto de métodos que deben ser implementados por el los tipos que desean acceder a compartir el tipo de la interfaz (polimorfismo). Recuerde definir las interfaces en la sección del archivo de código fuente destinada a la declaración de variables, definición de tipos e interfaces.

```go
type MiInterfaz interface {
  MetodoUno()
  MetodoDos(param string)int64
}

var myVar MiInterfaz = A{} /* Esto resultará en un error de compilación si el tipo A no implementa los métodos de MiInterfaz */
```
Definición de funciones

Una función es un bloque contenido de código que implementa una funcionalidad o tarea específica. Las funciones pueden tomar unos parámetros de entrada que son procesados en la funcipon y pueden generar una salida o valor de retorno. Las funciones en Go pueden tomar cero o múltiples parámetros de entrada de distintos tipos y pueden retornar cero o varios valores de distintos tipos también. Para declarar funciones se utiliza la palabara clave <code>func</code> seguida del nombre de la función, la lista de parámetros de entrada entre paréntesis y luego los tipos de los valores de retorno (entre paréntesis si son más de uno).

```go
/* función sin parámetros de entrada ni valores de retorno */
func voidFunc() {
  /* lógica de función */
}

/* parámetros de entrada y un valor de retorno */
func myFunc(x, y int) int {
  /* lógica de función */
}

/* múltiples valores de retorno */
func newFunc(s string) (int, error) {
  /* lógica de función */
}
```
Por ser las funciones tipos de datos, se pueden asignar a variables (<i>first-class functions</i>)

```go
/* función como variable */
var myFunc = func(s string) (int, error) {
  /* lógica de función */
}

/* Luego, para invocarla */
x, err := myFunc("hello")
```
Al interior de funciones también se pueden declarar funciones anónimas (<i>first-class functions</i>)

```go
/* función anónima dentro de otra función */
func voidFunc() {
  /* lógica de función */
  func(s string){
    fmt.Printf("Printing %s from anonymous func", s)
  }("web apps")
  /* lógica de función */
}
```
Las funciones en Go clasifican como <i>high-order functions</i> ya que pueden recibir funciones como parámetros de entrada y también pueden retornar funciones.

```go
/* una función que recibe parámetro de tipo función */
func highOrderFuncOne(f func(x, y int) error) {
  /* lógica de función */
}

/* una función que retorna valores de tipo función */
func highOrderFuncTwo() func(x, y int) error {
  /* lógica de función */
  return func(x, y int) error {
    /* lógica de función */
  }
}
```
Definición de tipos derivados (aliases)

Definir un alias o un tipo derivado se puede hacer con la palabra clave <code>type</code> el nombre del alias y el tipo del que se deriva. Estos tipos derivados son útiles cuando se necesita asociar métodos a los tipos nativos de Go, sirven como un intermediario entre el método y el tipo nativo.

```go
type MyIntType int
var myVar = MyIntType(50)
```
### Variables, constantes, tipos y funciones exportadas y no-exportadas

Las variables, constantes, tipos y funciones de un paquete se pueden exportar para que puedan ser visibles (accesibles) en los paquetes que lo importen. Los nombres exportados comienzan siempre con letra mayúscula mientras que los no-exportados (solo visibles en el paquete donde se declaran/definen) comienzan siempre con letra minúscula.

```go
package mi_paquete

import "fmt"

var VariableExportada string /* esta variable será visible en los paquetes que importan al paquete mi_paquete */
const noExportada = "non-exported" /* esta constante no será visible en los paquetes que importan al paquete mi_paquete */

/* tipo visible en los paquetes que importan al paquete mi_paquete */
type A struct {
}

/* tipo no visible en los paquetes que importan al paquete mi_paquete, solo en los archivos fuente que conforman a mi_paquete */
type client struct {
}
```

### Creando e importando paquetes

Para crear un paquete lo único que se necesita es crear un directorio para agrupar los archivos de código fuente que harán parte del paquete. Se debe procurar que el nombre del directorio se use en la línea package de los archivos.

<img width="1111" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/42c37b12-9a99-491f-90c7-e29eb6b7f65a">

Para importar el paquete lo debe incluir en la sección de import del archivo fuente desde donde está importando.

<img width="787" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/968fbde2-31f8-4b1b-962d-39101e8d3df5">

Puede dar un nombre o alias al import para llamar las funciones, variables, constantes y tipos exportados por el paquete. Antes de invocar una función o hacer referencia a una variable debe utilizar el alias dado al paquete.

<img width="789" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/f88f1235-6c2a-4c01-bc8a-f504f20754b0">

O si está seguro de que no van a haber colisiones entre los nombres de las funciones (variables, constantes y tipos) de los paquetes importados puede utilizar el caracter punto como alias del import para poder utilizar directamente los nombres importados.

<img width="790" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/dd3c00bc-3b04-419f-8b8e-a6fbf72bd2d0">

Por defecto, Go retira, de forma automática, de la lista de imports los paquetes de los que no se está utilizando ningún nombre exportado, si se desea conservar el paquete en la lista a pesar de no estar usandolo se debe utilizar como alias el operador <code>blank identifier</code> cuyo símbolo es el caractér <i>underscore</i> <code>_</code>.

<img width="762" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/69522fff-c7e8-401b-99d2-c9e9bffa4139">

### Operadores disponibles en Go

__Aritméticos y de nivel de bit__

```
+    suma                      integers, floats, complex, strings (concatenación)
-    resta                     integers, floats, complex
*    producto                  integers, floats, complex
/    división                  integers, floats, complex
%    modulo                    integers

&    bitwise AND               integers
|    bitwise OR                integers
^    bitwise XOR               integers
&^   bit clear (AND NOT)       integers

<<   corrimiento a izquierda   integers
>>   corrimiento a derecha     integers
```
__Comparación__

```
==    igual
!=    no igual
<     menor
<=    menor o igual
>     mayor
>=    mayor o igual
```
__Lógicos__

```
&&    AND condicional   
||    OR condicional   
!     negación NOT                
```
__Conversiones de tipo (type castings)__

Recuerde que Go es un lengiuaje fuertemente tipado, por tanto el lenguaje no realiza conversiones implícitas (truncamientos o amplificaciones). Las conversiones que sean necesarias se deben realizar de forma explícita.

```go
i := 42 // int por defecto
f := float64(i) // coversión de int a float64
u := uint32(f) // conversión de float64 a uint32
```

```go
type CustomString string

ss := CustomString("hello")
tt := string(ss)  // esta conversión es posible ya que el tipo CustomString es un tipo derivado de string
```
### Estructuras de control de flujo de la aplicación

__Ciclos (loops)__

Ciclo for clásico

```golang
for [ Initial Statement ] ; [ Condition ] ; [ Post Statement ] {
    [Actions]
}
```
```golang
for i := 0; i < N; i++{
  //lógica del ciclo
}
```
```golang
counter := 0 /*  sacando la variable fuera del alcance (scope) del ciclo
for ; counter < N; counter++{
  //do something
}
```
```golang
ss := []uint{1. 2, 3, 4, 5, 6}
/* ciclo for con dos variables al tiempo */
for i, j := 0, len(ss)-1; i < len(ss)/2; i, j = i+1, j-1 {
  ss[i], ss[j] = ss[j], ss[i]
}
fmt.Printf("%#v", ss)
```
Ciclo while

```golang
cond := N
for cond > 0 {
  // lógica del ciclo
  cond--
}
```
Ciclo infinito

```golang
for {
  if algunaCondicion {
    break
  }
  if otraCondicion{
    continue
  }
  // lógica del ciclo
}
```
Ciclo iterativo sobre elementos de slices y maps

```golang
ss := []string{"hello", "world"}
m := map[int]string{1:"hello", 2:"world"}

for i, val := range ss {
  fmt.Println("ss en la posición", i, "es", val)
}
for k, v := range m {
  fmt.Println("m en la clave", k, "tiene el valor", v)
}
```
Puede silenciar variables de ciclos con el operador <i>blank identifier</i> cuando no las necesite
```golang
ss := []string{"hello", "world"}
m := map[int]string{1:"hello", 2:"world"}

for _, val := range ss {
  fmt.Println(val)
}
for _, v := range m {
  fmt.Println(v)
}
```
__Condicionales__

if...else if...else

```golang
if [ condition ] {
    [Actions]
}else if [ alternative condition ] {
    [Actions]
}else{
    [Actions]
} 
```
Asignación de una variable y evaluación de la condición en el mismo lugar

```golang
if value := rand.Intn(100); value > 60 {
  //do something
}
```
```golang
if resp, err := someFunc(); err != nil {
  // error!
}
```
```golang
m1 := map[string]int
if value, ok := m1["key"]; !ok {
  // no existe esa clave en el mapa
}
```
Switch, case... default

Switch clásico

```golang
now := time.Now().Unix()
mins := now % 2
switch mins {
  case 0:
    fmt.Println("even")
  case 1:
    fmt.Println("odd")
}
```
Múltiples evaluaciones para el mismo caso

```golang
score := 7
switch score {
  case 0, 1, 3:
    fmt.Println("Terrible")
  case 4, 5:
    fmt.Println("Mediocre")
  case 6, 7:
    fmt.Println("Not bad")
  case 8, 9:
    fmt.Println("Almost perfect")
  case 10:
    fmt.Println("hmm did you cheat?")
  default:
    fmt.Println(score, " off the chart")
}
```
La palabra <i>fallthrough</i> sirve para pasar de un caso a otro de forma directa, aún cuando no se cumple la condición del siguiente caso

```golang
flavors := []string{"chocolate", "vanilla", "strawberry", "banana"}
for _, flavor := range flavors {
  switch flavor {
    case "strawberry":
      fmt.Println(flavor, "is my favorite!")
      fallthrough
    case "vanilla", "chocolate":
      fmt.Println(flavor, "is great!")
    default:
      fmt.Println("I've never tried", flavor, "before")
  }
}
```
Opción evaluada directamente en los cases
```golang
t := time.Now().Hour()
switch {
case t < 12:
  fmt.Println("It's before noon")
default:
  fmt.Println("It's after noon")
}
```

### Asociación de métodos a structs e implementación de interfaces

Para asociar métodos a las estructuras (struct) se debe utilizar un parámetro conocido como receptor o <i>function receiver</i> que se agrega a la definición de una función para indicar que no es cualquier función sino que por el contrario es un método asociado al tipo del struct indicado en el receptor. Por lo tanto, la invocación de esa función solo puede hacerse a través de una instancia o variable del tipo definido por el struct.

[Ejemplo](https://go.dev/play/p/TKv_EIr8ZlL)

```go
type Usuario struct {
   Nombre string
   Edad   uint
   Status string
   Activo bool
}

func(u *Usuario) SetNombre(name string) {
   u.Nombre = name
}

func(u *Usuario) GetNombre() string {
   return u.Nombre
}

func(u *Usuario) EstaActivo() string {
   if u.Activo {
      return fmt.Sprintf("El usuario %s, está activo con status %s", u.Nombre, u.Status)
   }
   return fmt.Sprintf("El usuario %s está inactivo", u.Nombre)
}
```
El receptor (<i>function receiver</i>) del método se agrega después de la palabra <code>func</code>. Y entre paréntesis se agrega un parámetro que simplemente es un nombre para el receptor, en el ejemplo anterior era "u" y el tipo, en el ejemplo "*Usuario". El receptor indica que el método está asociado a ese tipo definido en el receptor mismo. Lo anterior determina que el método solo se puede llamar a través de una instancia del tipo inidicado en el receptor.

Los receptores pueden ser de tipo puntero, como en este ejemplo en particular, esto quiere decir que el método está capacitado para modificar a través del método los campos de la instancia (¿objeto?).

Para implementar una interfaz, simplemente se deben implementar los métodos definidos en la interfaz en el tipo que se quiere que implemente la interfaz.

[Ejemplo](https://go.dev/play/p/honiepXyqW2)

```go
type Reader interface {
  Read() string
}

type Writer interface {
  Write(w string)
}

/* una interfaz construída a partir de otras interfaces (composición) */
type ReadWriter interface {
  Reader
  Writer
}

/* tipo a implementar las interfaces */
type Reporte struct {
   Propiedad string  
}

func(r *Reporte) Read() string{
   return r.Propiedad
}

func(r *Reporte) Write(w string) {
   r.Propiedad = w
}

/* solo un tipo que implemente la interfaz ReadWrite se puede pasar a esta función */
/* polimorfismo en Go */
func ModificarPropiedad(rw ReadWriter, prop string) {
   rw.Write(prop)
   fmt.Println("Nuevo valor de la propiedad:", rw.Read())
}
```
En Go no está desarrollado el concepto de <i>Herencia</i> que conocemos de lenguajes plenamente orientados a objetos (C++, Python, Java, etc...). Pero en cambio si está bastante desarrollado el concepto de <i>Composición</i>. Luego, se puede lograr el mismo efecto de la herencia a través de la composición de structs. La composición se implementa agregando al struct un campo del tipo de la struct de la cuál se quiere recibir los métodos y campos. Este campo puede tener nombre o no, cuándo no tiene nombre se le conoce como <i>campo incrustado</i>.

[Ejemplo](https://go.dev/play/p/Ecrv9x53glm)

```go
type Base struct {
   Propiedad string
   
}

func(b *Base) SetPropiedad(val string) {
   b.Propiedad = val
}

func(b Base) GetPropiedad() string {
   return b.Propiedad
}

type Composed struct {
  Base  /* composición del tipo Base con campo incrustado */
}

func (c Composed) CheckEsValido() bool {
  return c.EsValido  /* campo obtenido por composición */
}

func main() {
  comp := Composed{Base: Base{EsValido: true}}
  comp.SetPropiedad("property")  /* método obtenido por composición */
  fmt.Println(comp.GetPropiedad()) /* método obtenido por composición */
  if comp.CheckEsValido() {
    fmt.Println("es válido")
  }
}
```
### Manejo y propagación de errores en Go

El manejo y propagación de errores en Go se hace de forma manual ya que en Go no existe el concepto de excepciones, y en cambio existe una interfaz con un método que permite encapsular los mensajes de error para propagarlos en las diferentes capas de la aplicación.

La interfaz <code>error</code> pertenece a la librería estándar de Go y está definida de la siguiente manera:

```go
type error interface {
  Error() string
}
```
Luego, cualquier tipo que implemente el método <code>Error()</code>, implementa correctamente a la interfaz <code>error</code>.

El paquete <code>errors</code> de la librería estándar de Go tiene la función <code>New</code> que permite encapsular un mensaje en una variable del tipo de la interfaz <code>error</code>. También el paquete <code>fmt</code> cuenta con la función <code>Errorf</code> que permite formatear una cadena de caracteres que satisface la interfaz <code>error</code>.

Se asume ausencia de error cuando al evaluar el valor de una variable o un retorno del tipo <code>error</code> da como resultado <code>nil</code>.

[Ejemplo](https://go.dev/play/p/PEXmyQk2dzP)

```go
type Custom struct {
  CustomId int64
}

func (c *Custom) SetCustomId(id string) error {
  /* almacenando el error de la función strconv.Atoi en la variable err */
  numId, err := strconv.Atoi(id)
  /* chequeando si hubo error */
  if err != nil {
    /* propagando el error a donde se llame el método SetCustomId */
    return fmt.Errorf("falló al setear el campo CustomId, con error: %w", err)
  }
  c.CustomId = int64(numId)
  return nil
}

func main() {
  cust := &Custom{}
  /* capturando el posible error */
  err := cust.SetCustomId("once")
  /* chequeando si hubo error */
  if err != nil {
    /* lógica en caso de error */
    fmt.Printf("Error: %s", err.Error())
  }
}
```

### Paquetes de la librería estándar de Go y módulos de terceros útiles para el desarrollo Web (backend) con Go

__Librería estándar__
<ul>
  <li><b>fmt:</b> Cuenta con funciones para el formato de entrada y salida por consola.</li>
  <li><b>time:</b> Este paquete proporciona las funcionalidades básicas para el despliegue y cálculo de medidas de tiempo (fechas, hora, etc.)</li>
  <li><b>math:</b> Proporciona constantes y funciones matemáticas de propósito general.</li>
  <li><b>strings:</b> Implementaciones útiles para el manejo, manipulación y formato de cadenas de caractéres codificadas en UTF-8.</li>
  <li><b>strconv:</b> Cuenta con funciones que permiten hacer conversiones entre cadenas de caratéres y otros tipos de dato disponibles en el lenguaje.</li>
  <li><b>encoding:</b> Proporciona utilidades para la conversión de datos desde el nivel de byte a representaciones textuales. El paquete encoding/json cuenta con las funciones para conversion desde/hacia JSON.</li>
  <li><b>net/url:</b> Contiene funcionalidades para la manipulación de URLs y la extracción de paramétros contenidos en estas.</li>
  <li><b>net/http</b> Proporciona interfaces y funciones para imlpementar aplicaciones tipo cliente/servidor sobre los protocolos TCP y HTTP.</li>
</ul>

__Módulos de terceros__
<ul>
  <li><b>carbon:</b> Es un paquete que facilita la manipulación y despliegue de mechas y medidas de tiempo. Con funciones más intuitivas por su semántica que las del paquete <i>time</i>. Repo: https://github.com/golang-module/carbon</li>
  <li><b>gorilla/mux:</b> Incluye un router y un dispatcher para una correcta atención de peticiones HTTP, facilitando la asociación de las peticiones con el respectiva función para atenderlas. Repo: https://github.com/gorilla/mux</li>
  <li><b>http-router:</b> Este paquete proporciona un sencillo pero eficaz router o multiplexador para solicitudes HTTP, muy útil cuando se necesita coincidir patrones específicos en las rutas a los endpoints expuestos por tu servidor. Repo: https://github.com/julienschmidt/httprouter</li>
  <li><b>resty:</b> Resty es un cliente para hacer solitudes HTTP. Ofrece funciones intuitivas y fáciles de utilizar contra un servicio que se requiere consultar. Repo: https://github.com/go-resty/resty</li>
  <li><b>echo:</b> Es un framework para desarrollo web con Go, especialmente útil para la construcción de REST APIs. Repo: https://github.com/labstack/echo</li>
  <li><b>fiber:</b> Otro framework web, pero este está inspirado en el framework para Node.js <i>Express.js</i> Repo: https://github.com/gofiber/fiber</li>
  <li><b>gin:</b> Es uno de los frameworks para desarrollo web más populares de Go. Se diferencia de los demás por su alto desempeño. Repo: https://github.com/gin-gonic/gin</li>
  <li><b>sqlx:</b> Es un paquete con interfaces y funciones para la interacción con bases de datos SQL. Está construido a partir del paquete <i>database/sql</i> de la librería estándar. Repo: https://github.com/jmoiron/sqlx</li>  
  <li><b>gorm:</b> Este paquete es un ORM (Object-Relational Mapping) para SQL. Un ORM es una técnica que permite hacer solicitudes y manipular datos de una base de datos utilizando el paradigma de orientación a objetos. Repo: https://github.com/go-gorm/gorm</li>
  <li><b>mongo-go-driver:</b> Proporciona interfaces y funciones para interactuar con la base de datos NoSQL basada en documentos Mongo-DB. Repo: https://github.com/mongodb/mongo-go-driver</li>
</ul>

__Cómo incluir modulos de terceros en tu proyecto__

El manejo de dependencias con <i>go-modules</i> es bastante sencillo e intutivo. Con un único comando se descargan los paquetes directamente desde su repositorio y son agregados como dependencias al archivo <i>go.mod</i>.

Para agregar un paquete de terceros se utiliza el comando <code>go get -u &lt;ruta-repositorio-paquete&gt;</code>. Este comando iniciará la descarga del paquete y lo incluirá como dependencia de tu proyecto en el archivo <i>go.mod</i> para que a partir de ese momento este disponible para ser importado desde cualquier archivo fuente del proyecto.

Si tiene dudas acerca de la ruta correcta para el comando <i>go get</i>, por lo general los repositorios cuentan con información acerca de la instalación del paquete. Ejemplo:

<img width="953" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/649d6316-62b1-4a73-9f7f-30a2ecf5bcb9">

<img width="942" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/bdc674c6-30b9-49d7-8335-32063698b816">

Veamos por ejemplo como incluir el paquete <i>gorilla/mux</i> en nuestro proyecto:

1. Dirijase al repositorio https://github.com/gorilla/mux

<img width="1479" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/74a34c43-4971-4244-849e-f438e042030d">

2. Busque allí la documentación para la instalación (comando go get)

<img width="948" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/9976764d-87da-43b6-940a-e50072f67eb7">

3. Copie el comando y ejecutelo en la terminal de su proyecto

<img width="1226" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/dc7f20c4-b67b-4a4d-96d5-a0c959556d66">

4. Opcionalmente puede verificar la instalación en el archivo go.mod

<img width="874" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/82a1de57-0739-4d73-afe6-6dcbb386f5d3">

5. Importe el paquete en el archivo fuente donde lo va a utilizar. Tenga en cuenta que la línea en el import corresponde con la ruta al repositorio desde donde lo descargó.

<img width="1016" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/8996f368-72c2-4d98-b5ae-4e7af7e63582">

6. Ya está listo para utilizar las funciones y demás utilidades del paquete.

<img width="898" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/332d623a-74b5-4727-a0b6-e67bec111823">

### Comandos útiles para la construcción y ejecución del proyecto

<code>go run main.go</code>







