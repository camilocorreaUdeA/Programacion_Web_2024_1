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

Ahora el próposito será conectar nuestro contenedor con otras aplicaciones también desplegadas en sus propios contenedores de Docker, por ejemplo, si queremos conectar una aplicación de backend con una instancia de un motor de base de datos que también corra sobre un contenedor. ¿Cómo se puede lograr el despliegue de varios contenedores simultáneamente y que existan por defecto canales de comunicación entre ellos?

Para ese propósito la tecnología Docker ofrece la herramienta <b><i>docker compose</i></b> que permite desplegar aplicaciones construídas con múltiples contenedores de manera sencilla a través de un simple archivo de configuración en formato YAML (https://www.cloudbees.com/blog/yaml-tutorial-everything-you-need-get-started). 

Entonces, en un archivo llamado <code>docker-compose.yml</code> se definen todos los contenedores que van a hacer parte de la aplicación y una vez está listo el archivo se ejecuta un comando para activar todos los servicios (contenedores). Además, de manera automática se crea también una red de datos interna que permite la comunicación entre los contenedores que componen la aplicación.

Por favor siga atentamente los pasos descritos a continuación para desplegar una aplicación de backend (un CRUD sencillo) que estará conectado a una base de datos PostgreSQL en contenedores de Docker utilizando la herramienta <b><i>docker compose</i></b>.

1. Cree un nuevo módulo de Go y agregue el código fuente que se encuentra aquí.
2. Ejecute en la terminal los comandos <code>go mod download</code> y <code>go mod tidy</code> para actualizar el árbol de dependencias del proyecto (esto actualizará los archivos go.mod y go.sum)

```bash
go mod download
go mod tidy
```
3. Agregue un archivo <code>Dockerfile</code> al proyecto, en este archivo se definirá el contenedor para la aplicación de backend. Debe quedar como se muestra a continuación:

```docker
FROM golang:1.22.2

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /godocker

EXPOSE 8080
```
En resumen se está construyendo una imagen de Docker para nuestra aplicación de backend a partir de la imagen <code>golang:1.22.2</code>, y luego se construye el binario de la aplicación con el nombre godocker y se expone el puerto 8080 en el contenedor para la API REST. Observe que a diferencia del Dockerfile que habíamos hecho en el ejemplo anterior, en este no hay un comando <code>CMD</code> y es porque la ejecución del contenedor la vamos a controlar desde otro archivo llamado <code>docker-compose.yml</code>.

4. Agregue un archivo <code>.env</code> al proyecto, en este archivo se van a definir unas variables de ambiente que van a permitir establecer la conexión de la aplicación de backend al motor de base de datos PostgreSQL. El archivo debe verse así:

```env
DB_USER={usuario_base_datos}
DB_PASSWORD={password_base_datos}
DB_NAME={nombre_base_datos}
DB_PORT=5432
```
Los valores entre {} los puede escoger a su elección pero debe tenerlos muy presentes a la hora de definir el contenedor que desplegará la base de datos.<br>
Observe en el archivo <code>main.go</code> en la función <code>main</code> la invocación a la función que establece la conexión a la base de datos, <code>ConectarDB</code>, los valores que se mapean allí para componer la URL de conexión a la base de datos son tomados precisamente desde el archivo <code>.env</code>

```go
db, err := ConectarDB(fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")), "postgres")
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}
```
5. Agregue un archivo <code>docker-compose.yml</code> al proyecto, en este archivo se va a definir la infraestructura de la aplicación, en este ejemplo en particular consta de dos componentes: contenedor aplicación de backend en Go y contenedor de base de datos PostgreSQL.

Los componentes de la infraestructura se instancian en el archivo como servicios y volúmenes. Un servicio es un contenedor que despliega una funcionalidad o servicio de la aplicación, por ejemplo, la aplicación de backend o una base de datos. Mientras que los volúmenes son persistencias o almacenamientos de datos que utilizan los servicios (pueden ser reutilizados por múltiples servicios).

En el archivo <code>docker-compose.yml</code> defina una sección para los servicios de la aplicación (contenedores) y otra para los vólumenes o persistencias de datos compartidas entre los servicios:

```docker
services:

volumes:
```


5.1 Agregando el contenedor de base datos PostgreSQL:

Agregue un servicio de base de datos a la sección de servicios del archivo <code>docker-compose.yml</code>. Este contenedor se construye a partir de la imagen <code>postgres:alpine</code>, configura el usuario, password y nombre de la base de datos a partir de los valores asignados a las variables en el archivo <code>.env</code>, también expone y mapea el puerto por el que se establece la conexión a la base de datos (5432, puerto por defecto del servicio de base de datos SQL) y finalmente se instala o monta el volúmen de datos en la ruta <code>/var/lib/postgresql/data</code> con el nombre <code>postgres-db</code> y una vez inicializada la base de datos se corre un comando que a su vez ejecuta las sentencias SQL que se definan en un archivo llamado <code>queries.sql</code> que se encuentra en la raíz del proyecto. Este último archivo es muy útil para inicializar la base de datos creando las tablas y populando las filas y columnas con datos. El archivo <code>docker-compose.yml</code> debe verse así por el momento:

```docker
services:
  db:
    image: postgres:alpine
    environment:
      - POSTGRES_USER=${DB_USER}
      - POSTGRES_PASSWORD=${DB_PASSWORD}
      - POSTGRES_DB=${DB_NAME}
    ports:
      - 5432:5432
    volumes:
      - postgres-db:/var/lib/postgresql/data
      - ./queries.sql:/docker-entrypoint-initdb.d/create_tables.sql

volumes:
  postgres-db:
```
5.2 Agregue el archivo <code>queries.sql</code> para inicializar la base de datos:

Las sentencias de SQL que se encuentran en este archivo serán ejecutadas para inicializar la base de datos, incluye la creación de la tabla y unos datos iniciales en la misma. Este archivo debe agregarse a la raíz del proyecto de backend ya que en el archivo <code>docker-compose.yml</code> en la sección de volúmenes del servicio de base de datos se indicó que se buscara allí ese archivo para ser ejecutado una vez el contenedor se encuentre corriendo.

```sql
CREATE TABLE IF NOT EXISTS comentarios(
  id SERIAL PRIMARY KEY,
  time TIMESTAMP,
  comment TEXT,
  reactions INT
);

INSERT INTO comentarios (time, comment, reactions) 
VALUES
('2023-10-14 07:33:12', 'Estaban pasando bueno, no? xD', 10)
('2024-04-29 23:59:59', 'Felicitaciones, se ven muy bien juntos', 5)
('2022-12-30 12:11:45', 'Happy New Year!', 20)
('2024-04-29 23:59:59', 'You look great!', 2) RETURNING id;
```
5.3 Agregando el contenedor de la aplicación de backend definido en el archivo Dockerfile:

```docker
services:
  db:
    image: postgres:alpine
    environment:
      - POSTGRES_USER=${DB_USER}
      - POSTGRES_PASSWORD=${DB_PASSWORD}
      - POSTGRES_DB=${DB_NAME}
    ports:
      - 5432:5432
    volumes:
      - postgres-db:/var/lib/postgresql/data
      - ./queries.sql:/docker-entrypoint-initdb.d/create_tables.sql
  web:
    build: .
    env_file:
      - .env
    ports:
      - 8080:8080
    volumes:
      - .:/app
    command: go run main.go -b 0.0.0.0

volumes:
  postgres-db:
```











