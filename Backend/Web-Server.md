# Implementando un servidor web con Go

## Paquete net/http de la librería estándar de Go

El paquete <code>net</code> proporciona una interfaz portable para procesos de red (networking) que incluye implementaciones para el establecimiento y mantenimiento de conexiones a través de los protocolos <b>TCP/IP</b> y <b>UDP</b>, resolución de nombres de dominio con el porotocolo <b>DNS</b> y conexiones a dominios de Unix a través de web sockets.

El paquete <code>net</code> proporciona las funciones básicas y de bajo nivel para la conexión a redes de transmisión de datos, donde las interfaces más importantes y de uso más extendido son las que proporcionan la función <code>Dial</code> (que retorna un objeto de conexión a un sistema remoto con alguno de los protocolos <b>TCP/IP</b> o <b>UDP</b>) y la función <code>Listen</code> (crea un objeto que permite escuchar activamente en un puerto <b>TCP/IP</b> o <b>UDP</b>).

<b>Dial: Proporiciona una interfaz de bajo nivel para abrir una conexión a un servidor</b>

```go
conn, err := net.Dial("tcp", "golang.org:80")
if err != nil {
	// atienda el error
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')
// ...
```

<b>Listen: Proporciona una interfaz de bajo nivel que permite que un servidor escuche (espere) en un puerto por conexiones de clientes</b>

```go
ln, err := net.Listen("tcp", ":8080")
if err != nil {
	// atienda el error
}
for {
	conn, err := ln.Accept()
	if err != nil {
		// atienda el error
	}
	go handleConnection(conn)
}
```

A partir de las interfaces proporcionadas por el paquete <code>net</code> se contruye el paquete <code>net/http</code> que cuenta con implementaciones de gran utilidad para el despliegue de clientes y sevidores HTTP con Go, entre las que se distinguen:

<b>Constantes para los métodos HTTP</b>

```go
const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
  // ...
)
```
<b>Constantes para los códigos de estatus de solicitudes y respuestas HTTP</b>

```go
const (
	StatusContinue           = 100 // RFC 9110, 15.2.1
	StatusSwitchingProtocols = 101 // RFC 9110, 15.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusOK                   = 200 // RFC 9110, 15.3.1
	StatusCreated              = 201 // RFC 9110, 15.3.2
	StatusAccepted             = 202 // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 9110, 15.3.4
  // ...
)
```

<b>Funciones para instanciar servidores HTTP y atender solicitudes HTTP</b>

```go
ServeHTTP(w http.ResponseWriter, r *http.Request)
http.Handle("/endpoint", handler)
http.HandleFunc("/endpoint", handler)
http.ListenAndServe("direccion:puerto", nil)
```

<b>Tipos</b>

```go
type Client struct{}
type Handler interface{}
type HandlerFunc func(ResponseWriter, *Request)
type Request struct{}
type Response struct{}
type ReponseWriter interface{}
type Header map[string][]string
```

## Implementación de un servidor web básico con http.Handle

```go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type CustomHandler struct {
	numConn int
	msg     string
}

func (ch *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ch.numConn++
	ch.msg = fmt.Sprintf("Se ha conectado al servidor. Conexion numero: %d\n", ch.numConn)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, ch.msg)
}

func main() {
	handler := &CustomHandler{}
	http.Handle("/hello", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```

### http.ResponseWriter y http.Request

<code>http.Request</code> es un tipo (struct) del paquete <code>net/http</code> para representar una solicitud recibida en un servidor HTTP y enviada desde un cliente HTTP. Algunos campos relevantes del tipo Request son: 

<ol>
  <li><b>Method:</b> es un string que indica el método HTTP de la solicitud. Un string vacío equivale a una solicitud GET.</li>
  <li><b>Header:</b> este campo del tipo http.Header (map[string][]string) contiene los encabezados incluídos en la solicitud.</li>
  <li><b>Body:</b> Este campo contiene el cuerpo de la solicitud. Relevante en solicitudes POST, PUT o PATCH.</li>
</ol>

<code>http.ResponseWriter</code> es una interfaz que define los métodos necesarios para construir la respuesta a una solicitud HTTP, y son los que listan a continuación:

<ol>
  <li><b>Header() Header:</b> Retorna los encabezados que se van a enviar en la respuesta HTTP. Se pueden agregar encabezados con el método Set del tipo http.Header</li>
  <li><b>Write([]byte)(int, error):</b> Escribe los datos de la espuesta HTTP y los envía al cliente a través de la conexión establecida previamente entre cliente y servidor. Para hacer el código del ejemplo más expresivo se reemplazó el llamado directo a esta función por <code>io.WriteString(w, msg)</code> que escribe el string msg directamente en el ResponseWriter w.</li>
  <li><b>WriteHeader(statusCode int):</b> Este método permite indicar el código de estatus de la respuesta. Por defecto se envía un código http.StatusOK</li>
</ol>

### http.Handle

Handle es una función que permite registrar un <code>handler</code> (una función que atiende una solicitud HTTP) contra un patrón que determina un posible conjunto de endpoints válidos expuestos por el servidor. El tipo de solicitud y el endpoint al que va dirigida permiten que Go internamente encuentre una coincidencia para el patrón asociado al handler.

Para construir un handler que sea válido se debe implementar la interfaz <code>http.Handler</code> que solo tiene un método:

```go
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}
```

El método <code>ServeHTTP</code> se implementa de forma que se lean los detalles de la solicitud que vienen en el objeto Request y se escriba la respuesta a la solicitud en un objeto ResponseWriter.

### http.ListenAndServe

```go
func ListenAndServe(addr string, handler Handler) error
```

Esta función ejecuta dos tareas en un solo llamado. En primer lugar se encarga de escuchar (esperar solicitudes HTTP) en la dirección y puerto que se especifican en el parámetro <code>addr</code> y luego invoca atiende la solicitud a través de la implementación disponible en el parámetro <code>handler</code>. Cuando se pasa un handler <code>nil</code> (nulo) entonces se invocan los handlers asociados al multiplexor por defecto del paquete <code>net/http</code> (DefaultServeMux).

## Implementando el mismo servidor web básico con http.HandlerFunc

```go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type CustomHandler struct {
	numConn int
	msg     string
}

func (ch *CustomHandler) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ch.numConn++
		ch.msg = fmt.Sprintf("Se ha conectado al servidor. Conexion numero: %d\n", ch.numConn)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, ch.msg)
	}
}

func main() {
	handler := &CustomHandler{}
	http.HandleFunc("/hello", handler.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
### http.HandlerFunc

```go
type HandlerFunc func(ResponseWriter, *Request)
```
El tipo HandlerFunc es un adaptador que permite implementar handlers HTTP a través de funciones ordinarias y no de tipos como en el caso de http.Handle

### http.HandleFunc

```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

Esta función permite registrar un handler (implementado a través de una función) contra un patrón que determina un posible conjunto de endpoints válidos expuestos por el servidor. El tipo de solicitud y el endpoint al que va dirigida permiten que Go internamente encuentre una coincidencia para el patrón asociado al handler.










