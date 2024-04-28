# Handlers y Multiplexers

## Formato del cuerpo (Body) de las solicitudes y respuestas de la API REST

En la revisión del tema de APIs REST se estableció que el formato elegido para construir el cuerpo de una solicitud HTTP (métodos POST, PUT y PATCH) o de una respuesta HTTP (GET, POST, PUT, PATCH, DELETE) iba a ser Javascript Object Notation más conocido como JSON. Este formato permite organizar los datos en estructuras similares a los objetos del lenguaje Javascript donde a una clave de tipo string se le asocia un valor de cualquier tipo (numérico, string, arreglo, otros objetos, etc).

El formato JSON permite serializar los datos de fora que se facilite el proceso para convertirlos a una representación binaria que es la que finalmente viaja en el protocolo a través de la conexión establecida entre cliente y servidor (conexión TCP/IP).

Ejemplo de datos serializados utilizando JSON:

```json
{
   "nombre": "Chispita",
   "especie": "Felino",
   "edad": 5,
   "sexo": "F",
   "detalles": {
     "adoptado": true,
     "congenitas": [],
     "estado": "optimo"
   }
}
```
El lenguaje de programación Go cuenta con funciones en el paquete <code>encoding/json</code> para serializar (codificar) una estructura (<code>struct</code>) a su representación en JSON y también para el proceso inverso, es decir, para de-serializar un objeto en formato JSON a una estructura de Go.

### json.Marshal

<code>json.Marshal</code> es una función de la librería estándar de Go que permite serializar los campos de una estructura en su respectiva representación en el formato JSON. La función retorn dos valores, un slice de elementos de tipo <code>byte</code> que es propiamente el objeto serializado en su representación binaria y un error para registrar cualquier inconveniente que se haya presentado en el proceso de serialización y impidiendo que culmine exitosamente. La función recibe comp parámetro de entrada el objeto o instancia de la estructura que se pretende serializar.

Los siguientes son requerimientos básicos del proceso de serialización a JSON en Go:

<ul>
  <li>Solo es posible serializar estructuras cuyos campos son exportados</li>
  <li>Se pueden utilizar <code>tags</code> o etiquetas en las estructuras para indicar un nombre específico a un campo en el objeto JSON resultante. Si no se usan las etiquetas entonces un procedimiento interno de reflejo de campos de estructuras en el proceso de serialización asignará a los campos de el objeto JSON los nombres correspondientes en la estructura de Go tal cual están escritos</li>
</ul>

```go
import "encoding/json"

type Mascota struct{
  Nombre   string `json:"nombre"`
  Especie  string `json:"especie"`
  Edad     uint8 `json:"edad"`
  Sexo     string `json:"sexo"`
  Detalles *Detalles `json:"detalles"`
}

type Detalles struct{
  Adoptado               bool `json:"adoptado"`
  EnfermedadesCongenitas []string `json:"congenitas"`
  EstadoEnLaAdopcion     string `json:"estado"`
}

chispita := Mascota{//inicializacion de los campos del objeto...}

mascota, err := json.Marshal(chispita)
if err != nil{
   //atienda el error
}

//utilice el objeto serializado "mascota" ...  para enviarlo en el body de una respuesta HTTP por ejemplo ...
```

### json.Unmarshal

<code>json.Unmarshal</code> es la función que realiza el proceso inverso a la anterior, es decir, convierte un objeto serializado en JSON a su correspondiente representación en una instancia u objeto de una estructura definida en la aplicación de Go. Esta función retorna un error en caso de que el proceso falle por alguna razón y recibe dos parámetros de entrada, un slice de <code>byte</code> para pasar el objeto JSON y una referencia a un objeto del tipo de estructura donde se va a almacenar el resultado del proceso de de-serialización.

Para que el proceso de de-serialización se complete correctamente se requiere que por lo menos:

<ul>
  <li>La estructura destino de los datos del objeto JSON realmente tenga campos correspondientes con los nombres de las claves en el objeto</li>
  <li>Si la estructura no cuenta con todos los campos definidos en el objeto JSON entonces solo se van a de-serializar aquellos que tengan correspondencia</li>
</ul>

```go
import "encoding/json"

//"mascota" es un objeto JSON (slice de byte)... que viene en el body de una solicitud HTTP por ejemplo
mascota := []byte(`{
  "nombre": "Chispita",
  "especie": "Gato",
  "edad": 5,
  "sexo": "F",
  "detalles": {
    "adoptado": true,
    "congenitas": [],
    "estado": "optimo"
  }
}`)

chispita := &Mascota{}
err := json.Unmarshal(mascota, chispita)
if err != nil{
   //atienda el error
}

//utilice el objeto de-serializado "chispita"...
//por ejemplo para guardarlo en una base de datos o procesar sus campos de manera independiente ...
```

## Handlers: Manejadores para atender solicitudes HTTP a los endpoints (rutas a servicios) expuestos por un servidor web

### Handler

Un web handler, handler HTTP, o simplemente handler es un componente de software del lado del servidor cuya tarea es la de atender peticiones web (solicitudes utilizando los métodos del protocolo HTTP) que son enviadas desde clientes (navegadores, aplicaciones móviles, dispositivos IoT, etc.) a los servicios expuestos por un servidor web. El handler es el encargado de recibir la petición, procesarla y devolver una respuesta al cliente.

El paquete <code>net/http</code> de Go ofrece dos alternativas para la implementación de handlers, la primera está basada en la construcción del handler a partir de un objeto de una estructura que implementa la interfaz <code>Handler</code> incluída en el paquete. La segunda, se basa en la implementación del handler a través de una función que retorne una función del tipo <code>HandlerFunc</code> que no es más que un adaptador del paquete que implementa la interfaz <code>Handler</code>.

### Implementación de un handler con http.Handler

La interfaz Handler del paquete <code>net/http</code> está definida así:

```go
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}
```
De modo que para implementar esta interfaz a través de un tipo concreto basta con solo con implementar el método <code>ServeHTTP</code>.

Una vez implementada la interfaz se completa el handler al asociar el objeto que implementa la interfaz http.Handler a la ruta o endpoint que va a exponer el servicio de la aplicación web. 

Dicha asociación se hace al invocar la función <code>http.Handle</code> que recibe como argumentos un patrón o cadena de caracteres que identifica la ruta o endpoint mediante el que se expone el servicio y un objeto del tipo que implementa la interfaz http.Handler.

Ejemplo:
```go
package main

import (
  "log"
  "net/http"
)

type CustomHandler struct {}

func (ch *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case http.MethodGet:
      // atienda una peticion GET al endpoint
    case http.MethodPost:
      // atienda una peticion POST al endpoint
    // ... otros metodos HTTP
    default:
      http.Error(w, "metodo no permitido", http.StatusMethodNotAllowed)
   }
}

func main() {
  handler := &CustomHandler{}
  http.Handle("/hello", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
```
### Implementación de un handler con http.HandlerFunc

El adaptador <code>http.HandlerFunc</code> funciona como un alias para las funciones con la firma <code>func (ResponseWriter, *Request)</code>. Este adaptador implementa la interfaz http.Handler de la siguiente manera:

```go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
  f(w, r)
}
```
Para implementar un handler con http.HandlerFunc basta con crear una función que retorne precisamente un valor de ese tipo. Y luego asociarlo a una ruta o endpoint con la función <code>http.HandleFunc</code> que recibe como argumentos un patrón o cadena de caracteres que identifica la ruta o endpoint mediante el que se expone el servicio y una función con firma <code>func (ResponseWriter, *Request)</code>, o sea un http.HandlerFunc.

Ejemplo:
```go
package main

import (
  "log"
  "net/http"
)

type CustomHandler struct{}

func Handler(ch *CustomHandler) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
      ch.get()
      //respuesta a la peticion
    }
    if r.Method == http.MethodPost {
      ch.post()
      //respuesta a la peticion
    }
    //...atender peticiones con otros metodos
  })
}

func (ch *CustomHandler) get() {
  //procesamiento de peticion GET
}

func (ch *CustomHandler) post() {
  //procesamiento de peticion POST
}

func main() {
  handler := &CustomHandler{}
  http.HandleFunc("/hello", Handler(handler))
  log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Multiplexers o Routers

Un multiplexer o router HTTP es un componente que permite encontrar una coincidencia entre una solicitud (petición) que llega al servidor y una lista de rutas o endpoints que se han registrado en la aplicación web. Una vez se encuentra una coincidencia se invoca el handler que ha sido asociado a esa ruta en particular.

### ¿Cuál es el próposito o beneficio de utilizar un mux o router?

De acuerdo con el autor Alex Edwards en un artículo publicado en su blog personal (https://www.alexedwards.net/blog/which-go-router-should-i-use) algunos de los beneficios de utilizar un mux o router para implementar un servidor web son:

<ul>
   <li><b>Soporte a enrutamiento basado en el método de la solicitud: </b>Utilizar un mux/router debería facilitar el despacho de las solicitudes a los handlers basandose en el método con el que se realizó la solicitud</li>
   <li><b>Soporte a variables en la ruta de la URL del endpoint: </b>Debería facilitar la declaración de endpoints como <code>/empleados/{id}</code> donde "{id}" es un parámetro dinámico (variable) en la URL del endpoint.</li>
   <li><b>Soporte a patrones de endpoints con expresiones regulares (regex): </b>El mux/router facilita la construcción de endpoints mediante el uso de expresiones regulares y además encuentra el "match" o la coincidencia con el handler que se debe ejecutar.</li>
   <li><b>Soporte a rutas basadas en host o nombre de dominio: </b>El mux/router permite despachar solicitudes a handlers basandose en la URL del host o nombre de dominio en lugar de solamente la URL de la ruta del endpoint.</li>
   <li><b>Soporte a rutas en conflicto: </b>Se refiere a que el mux/router esté en capacidad de registrar patrones de enrutamiento que podrían potencialmente coincidir con la ruta a la que va dirigida la solicitud.</li>
</ul>

### http.ServeMux del paquete <code>net/http</code>

<code>ServeMux</code> es el multiplexer para solicitudes HTTP que está incluído en la librería estándar de Go. Su funcionalidad principal es encontrar una correspondencia entre la URL a la que está dirigida la solicitud y un conjunto de patrones registrados contra unos handlers. La coincidencia más cercana será la que ejecute el handler correspondiente.

<code>ServeMux</code> puede encontrar correspondencias con patroness por método, nombre de dominio o ruta de la solicitud. En general, un patrón consta de las siguientes partes: <code>[MÉTODO ][HOST]/[RUTA]</code>, las 3 partes son opcionales y el caracter / es una ruta válida.

<ul>
   <li><b>/index.html </b>coincide la ruta index/html para cualquier host y cualquier método.</li>
   <li><b>GET /static/ </b>coincide cualquier solicitud GET cuya ruta comienza por /static/</li>
   <li><b>example.com/ </b>coincide cualquier solicitud al host example.com</li>
   <li><b>example.com/{$} </b>coincide la solicitud con el host example.com y la ruta /</li>
   <li><b>/files/{file}/photos/{photo} </b>coincide con las rutas cuyo primer segmento es files y el tercer segmento es photos. Los segmentos file y photo pueden ser valores dinámicos.</li>
</ul>

Ejemplo:

```go
package main

import (
  "log"
  "net/http"
)

type HandlerComentarios struct{}

func (hc *HandlerComentarios) listComentarios() http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    //procesamiento de peticion GET que devuelve lista de comentarios
  })
}

func (hc *HandlerComentarios) getComentario() http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    //procesamiento de peticion GET que devuelve un solo comentario
  })
}

func (hc *HandlerComentarios) crearComentario() http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    //procesamiento de peticion POST
  })
}

func main() {
  mux := http.NewServeMux()
  handler := &HandlerComentarios{}
  mux.HandleFunc("GET /comentarios", handler.listComentarios())
  mux.HandleFunc("GET /comentarios/{id}", handler.getComentario())
  mux.HandleFunc("POST /comentarios", handler.crearComentario())

  log.Fatal(http.ListenAndServe(":8080", mux))
}
```

### Otros mux/routers disponibles para desarrollo web en Go
























