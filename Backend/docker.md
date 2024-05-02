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

Docker es una plataforma de software de código abierto (open source) que le permite a los desarrolladores construir, probar y desplegar aplicaciones de manera ágil utilizando contenedores. Docker brinda las herramientas necesarias para empaquetar una aplicación con sus dependencias, configuraciones, librerías dentro de un contenedor como una única entidad de sorftware. Esta característica permite automatizar y agilizar el despliegue de una aplicación para que sea ejecutada en cualquier sistema operativo anfitrion.

Las aplicaciones construidas con Docker se pueden desplegar fácilmente sin tener en consideración el sistema operativo. De ahí se desprende que sea más fácil para los desarrolladores crear aplicaciones sin incurrir en problemas o inconvenientes ya que la aplicación va a estar aislada del resto de procesos que están corriendo en el sistema operativo anfitrion.

<b>Imágenes y contenedores de Docker</b>

Una imagen de Docker es una plantilla o <i>template</i> que define a un contenedor. La imagen contiene el código que se va a ejecutar incluyendo cualquier definición de cualquier librería o dependencia que el código necesita. Docker tiene un producto llamado DockerHub (https://www.docker.com/products/docker-hub/) que funciona como un registro de imágenes de Docker (<i>container registry</i>), allí se pueden encontrar tanto imágenes oficiales como otras creadas por otros usuarios de Docker.

Un contenedor de Docker es una instancia de una imagen de Docker que se está ejecutando.

<b>Instalación de Docker</b>

La forma más sencilla de instalar Docker es instalando la aplicación <b><i>Docker Desktop</i></b> (https://docs.docker.com/engine/install/) que está disponible para los sistemas operativos Windows, MacOS y Linux. El único problema de <b><i>Docker Desktop</i></b> es que solo es de uso libre para implementar proyectos que no sean comerciales o tengan apariencia de comerciales. Para esa situación existen alternativas como <b><i>Rancher Desktop</i></b> (https://rancherdesktop.io/) y <b><i>Colima</i></b> (https://github.com/abiosoft/colima) para los amantes de hacer todo en la línea de comandos sin necesidad de utilizar aplicaciones de escritorio.

También puede instalar Docker Engine sin necesidad de instalar la aplicación de escritorio Docker Desktop o Rancher Desktop. 

Linux Ubuntu: https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository
MacOS: https://docs.docker.com/desktop/install/mac-install/
Windows: https://docs.docker.com/desktop/install/windows-install/

<b>Alternativas para trabajar con contenedores</b>

Si tu sistema no soporta Docker o simplemente no te gusta, aquí puedes ver una lista de alternativas a Docker que te permiten trabajar con contenedores: https://spacelift.io/blog/docker-alternatives

## Desplegando una aplicación Go en un contenedor Docker

### 

### Desplegando con docker-compose










