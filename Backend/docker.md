# Dockeriza tu backend en Go

El objetivo de esta presentación es mostrar el proceso para "dockerizar" una aplicación de Go. Y en especial una aplicación de backend que está conectada a una base de datos.

Inicialmente exploraremos como crear un contenedor en Docker de una aplicación de Go a través de un <b><i>Dockerfile</i></b> y luego vamos a crear una infraestrucutra de contenedores con <b><i>docker compose</i></b>, este proceso nos va a permitir conectar el contenedor de nuestra aplicación de backend a un contenedor de base de datos.

## Docker y la tecnología de contenedores

### Contenedores

Definición tomada de la página oficial de Docker: Un contenedor es una unidad estándar de software que funciona como un paquete donde se empaca el código y las dependencias de una aplicación, de modo que la aplicación se puede ejecutar de forma rápida y confiable desde cualquier ambiente de computo. Un contenedor virtualiza el sistema operativo subyacente hciendo que la aplicación en el contenedor perciba que tiene dentro del contnedor todos los recursos de computo de la máquina anfitriona: CPU, memoria, almacenamiento de archivos, conexiones de red, etc. Esta abstracción permite que un contenedor pueda ser desplegado y ejecutado en cualquier parte, siempre y cuando la imagen base sea consistente.

La utilidad primordial del uso de contenedores es que permite aislar el código y sus dependencias a un ambiente en específico de modo que para el código pasa desapercibido en que sistema operativo se está ejecutando o que archivos y aplicaciones componen el sistema anfitrion. De allí se desprende que una aplicación en contenedores se puede mover fácilmente de un sistema a otro y a distintos ambientes dentro del ciclo de desarrollo de software (desarrollo, pruebas, producción, etc) sin que repercuta en la ejecución y funcionalidad de la aplicación.

Suele confundirse a los contenedores con máquinas virtuales, pero son dos conceptos muy distintos, y aunque ambos son utilizados para virtualización la principal diferencia radica en que las máquinas virtuales vienen con su propio sistema operativo y kernel y solo comparten los demás recursos de la máquina anfitriona con las demás máquinas virtuales. Por el contrario, los contenedores utilizan el mismo sistema operativo y kernel de la máquina anfitriona, y todos los contenedores que hayan sido desplegados en la máquina comparten los mismos recursos de cómputo. Es por esta razón que no se pueden desplegar contenedores con sistema operativo Windows en una máquina anfitriona Linux y viceversa. Pero en cambio si se pueden tener máquinas de cualquier sistema operativo indiferente del sistema operativo de la máquina anfitriona.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2024_1/assets/42076547/2680e983-796a-4979-bbef-375dbd5c4fa0)

### Docker

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2024_1/assets/42076547/a2ae551c-dfa8-493a-90f0-381bb6a4cfd4)

Docker es una plataforma de software de código abierto (open source) que le permite a los desarrolladores construir, probar y desplegar aplicaciones de manera ágil utilizando contenedores. Docker brinda las herramientas necesarias para empaquetar una aplicación con sus dependencias, configuraciones, librerías dentro de un contenedor como una única entidad de software. Esta característica permite automatizar y agilizar el despliegue de una aplicación para que sea ejecutada en cualquier sistema operativo anfitrion.

Las aplicaciones construidas con Docker se pueden desplegar fácilmente sin tener en consideración el sistema operativo. De ahí se desprende que sea más fácil para los desarrolladores crear aplicaciones sin incurrir en problemas o inconvenientes ya que la aplicación va a estar aislada del resto de procesos que están corriendo en el sistema operativo anfitrion.

<b>Imágenes y contenedores de Docker</b>

Una imagen de Docker es una plantilla o <i>template</i> que define a un contenedor. La imagen contiene el código que se va a ejecutar incluyendo cualquier definición de cualquier librería o dependencia que el código necesita. Docker tiene un producto llamado <b><i>DockerHub</i></b> (https://www.docker.com/products/docker-hub/) que funciona como un registro de imágenes de Docker (<i>container registry</i>), allí se pueden encontrar tanto imágenes oficiales como otras creadas por otros usuarios de Docker.

Un contenedor de Docker es una instancia de una imagen de Docker que se está ejecutando.

<b>Instalación de Docker</b>

La forma más sencilla de instalar Docker es instalando la aplicación <b><i>Docker Desktop</i></b> (https://docs.docker.com/engine/install/) que está disponible para los sistemas operativos Windows, MacOS y Linux. El único problema de <b><i>Docker Desktop</i></b> es que solo es de uso libre para implementar proyectos que no sean comerciales o tengan apariencia de comerciales. Para esa situación existen alternativas como <b><i>Rancher Desktop</i></b> (https://rancherdesktop.io/) y también <b><i>Colima</i></b> (https://github.com/abiosoft/colima) para los amantes de hacer todo en la línea de comandos sin necesidad de utilizar aplicaciones de escritorio.

También puede instalar Docker Engine sin necesidad de instalar la aplicación de escritorio Docker Desktop o Rancher Desktop. 

Linux Ubuntu: https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository<br>
MacOS: https://docs.docker.com/desktop/install/mac-install/<br>
Windows: https://docs.docker.com/desktop/install/windows-install/<br>

<b>Alternativas para trabajar con contenedores</b>

Si tu sistema no soporta Docker o simplemente no te gusta, aquí puedes ver una lista de alternativas a Docker que te permiten trabajar con contenedores: https://spacelift.io/blog/docker-alternatives

## Desplegando una aplicación Go en un contenedor Docker

### Desplegando con Dockerfile

Docker ouede construir imágenes automáticamente al leer un conjunto de instrucciones en un Dockerfile. Un Dockerfile es un documento de texto que contiene todos los comandos que un usuario ejecutaría en la línea de comandos para ensamblar una imagen.

Pasos:

1. Cree un nuevo proyecto de Go e inicialice el modulo

```bash
go mod init github.com/{nombre-repositorio}/go-containers
```
2. Agregue el paquete externo <code>gorilla/mux</code> a su proyecto

```bash
go get -u github.com/gorilla/mux
```  
3. Agregue un archivo <code>main.go</code> con el siguiente código

```go
package main

import (
  "io"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

func main() {
  mux := mux.NewRouter()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello World!")
  }).Methods(http.MethodGet)

  log.Fatal(http.ListenAndServe(":8080", mux))
}
```
4. Agruegue un archivo Dockerfile al proyecto para crear la descripción de la imagen de Docker. El archivo debe llamarse Dockerfile, sin extensión. Agregue la siguiente lista de comandos al archivo (conservando el orden).

```docker
FROM golang:1.22.2

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /godocker

EXPOSE 8080

CMD [ "/godocker" ]
```

Descripción de los comandos en el Dockerfile:<br>

<ul>
  <li><code>FROM golang:1.22.2</code>: Este comando indica que el contenedor se va a construir sobre la base de una imagen existente en el registro de contenedores de Docker, en este caso la imagen correspondiente a la versión 1.22.2 del lenguaje Go.</li>
  <li><code>WORKDIR /app</code>: Este comando crea un directorio dentro de la imagen de Docker e indica que el resto de comandos del Dockerfile se van a ejecutar con referencia a ese directorio.</li>
  <li><code>COPY . .</code>: Copia todos los archivos en el proyecto de Go al directorio especificado para la imagen de Docker (/app)</li>
  <li><code>RUN go mod download</code>: Corre el comando de Go que instala todas las dependencias listadas en el archivo go.mod</li>
  <li><code>RUN go build -o /godocker</code>: Corre el comando de Go que sirve para construir el binario de la aplicación con el nombre indicado (godocker)</li>
  <li><code>EXPOSE 8080</code>: Expone el puerto por el que va a estar escuchando la aplicación.</li>
  <li><code>CMD [ "/godocker" ]</code>: Este comando ejecuta el binario dentro del contenedor de Docker. Observe que es el mismo nombre del comando go build</li>
</ul>

5. Cree la imagen con Docker a partir del Dockerfile

En la línea de comandos ejecute la siguiente instrucción, y reemplace <code>{nombre-imagen}</code> con el nombre o etiqueta que le quiera asignar a la imagen. El punto al final indica que el Dockerfile y los archivos del proyecto se encuentran en el directorio actual, por eso es importante correr el comando desde el directorio del proyecto.

```bash
docker build -t {nombre-imagen} .
```
6. Verifique que se haya creado la imagen correctamente

Ejecute el siguiente comando y verifique que el nombre o etiqueta que le dio a la imagen aparece en la lista que ha sido retornada.

```bash
docker image ls
```

7. Ejecute el contenedor

Ejecute el comando a continuación para ejecutar un contenedor basado en la imagen que recien ha creado. Reemplace <code>{nombre-contenedor}</code> con un nombre para el contenedor y reemplace <code>{nombre-imagen}</code> con el nombre o etiqueta. La bandera d indica que se corra el contenedor en modo "detach" que no bloquee la terminal, la bandera p sirve para hacer la correspondencia entre el puerto expuesto por la aplicación y el puerto de la máquina afitriona.

```bash
docker run -d -p 8080:8080 --name {nombre-contenedor} {nombre-imagen}
```

8. Verifique que el contenedor esté corriendo

Busque el nombre del contenedor en la lista que retorna el comando a continuación.

```bash
docker container ls
```
9. Pruebe la aplicación

Visite la URL <code>localhost:8080</code> en el navegador o haga una petición GET a esa misma URL y verifique la respuesta esperada de la aplicación.

10. Detenga el contenedor

```bash
docker container stop {nombre-contenedor}
```
11. Elimine el contenedor

```bash
docker rm --force {nombre-contenedor}
```
Si quiere profundizar más en el uso de Docker puede consultar las guías oficiales de comandos de Dockerfile (https://docs.docker.com/reference/dockerfile/) y Docker (https://docs.docker.com/reference/cli/docker)

### Desplegando con docker-compose

Hemos visto como desplegar una aplicación sencilla de Go en un contenedor de Docker, esto es especialmente útil cuando vamos a desplegar una aplicación web usando una arquitectura de microservicios donde cada aplicación de servicio web construída con Go se puede desplegar en su propio contenedor e independiente de los demás servicios.

AHora el próposito será conectar nuestro contenedor con otras aplicaciones también desplegadas en sus propios contenedores de Docker, por ejemplo, si queremos conectar una aplicación de backend con una instancia de un motor de base de datos que también corra sobre un contenedor. ¿Cómo se puede lograr el despliegue de varios contenedores simultáneamente y que puedan existir canales de comunicación entre ellos?

Para ese propósito la tecnología Docker ofrece la herramienta <b><i>docker compose</i></b> que permite desplegar aplicaciones construídas con múltiples contenedores de manera sencilla a través de un simple archivo de configuración en formato YAML (https://www.cloudbees.com/blog/yaml-tutorial-everything-you-need-get-started). 

Entonces, en un archivo llamado <code>docker-compose.yml</code> se definen todos los contenedores que van a hacer parte de la aplicación y una vez está listo el archivo se ejecuta un comando para activar todos los servicios (contenedores). Además, de manera automática se crea también una red de datos interna que permite la comunicación entre los contenedores que componen la aplicación.








