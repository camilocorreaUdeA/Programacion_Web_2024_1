# Proyecto final del curso de Desarrollo de Aplicaciones Web 2024-1

## Acerca del proyecto

En el presente documento se presentan las condiciones y productos esperados para la entrega del proyecto final del curso cuyo porcentaje en la evaluación total del mismo corresponde al 25%.

La idea con este proyecto es poner en práctica todos los conceptos aprendidos a largo del semestre:

<ul>
  <li>Definición de la estructura y contenido de un sitio web (cliente de la aplicación web) con el lenguaje HTML</li>
  <li>Aplicación de estilos al sitio web con el lenguaje de hojas de estilo en cascáda CSS</li>
  <li>Implementación de clientes de API REST con Javascript</li>
  <li>Contenido dinámico y lógica del cliente web con Javascript</li>
  <li>Implementación de los servicios o capa de servidor de una aplicación web con el lenguaje de programación Go</li>
  <li>Implementación de una capa de repositorio en el servidor mediante el uso de un motor de base de datos</li>
  <li>Despliegue de la aplicación web utilizando contenedores (Docker)</li>
</ul>

Van a crear una REST API con sus dos componentes principales: el cliente (frontend) y el servidor (backend).

Se va a implementar una aplicación web para la reserva de automóviles. La descripción detallada de las condiciones y requerimientos necesarios para la implementación se puede ver a continuación.

### Descripción del proyecto

El proyecto consiste en una aplicación web sencilla con un frontend desarrollado con las tecnologías HTML, CSS y Javascript. Completada con su capa de backend desarrollada con el lenguaje de programación Go y una base de datos PostgreSQL.

La aplicación debe permitir a un usuario registrarse en la misma para luego poder tener acceso al servicio de reserva de los automóviles que se encuentran disponibles para ser reservados.

Al usuario se le deben proporcionar en la apliación los servicios necesarios para que la experiencia de reserva sea intuitiva y satisfactoria. Eso incluye una interaz gráfica agradable, amigable con el usuario y fácil de usar. Y unos servicios ágiles y que retornen la información que el usuario requiere para completar su proceso de reserva.

### Funcionamiento de la aplicación
<ul>
  <li>La primera vez que el usuario accede a la aplicación debe resgistrarse para poder empezar a hacer reservas</li>
  <li>Los usuarios registrados pueden acceder al módulo de reserva para buscar un autómovil para reservar</li>
  <li>Los automóviles se pueden buscar por distintos criterios: tipo, color, modelo, marca, precio</li>
  <li>Un usuario puede reservar la cantidad de automóviles que desee</li>
  <li>El usuario puede marcar opciones adicionales para la reserva (tipo de transmisión, equipamiento de lujo, seguros, asistencia en carretera, etc.)
  <li>La aplicación debe marcar los automóviles reservados para que no puedan ser reservados por otro usuario</li>
  <li>La aplicación debe presentar al usuario un reporte de la reserva donde se incluyan los detalles de la reserva y el precio total</li>
</ul>

### Requerimientos generales y entregables

<b>Requerimientos</b>

<ul>
  <li>La aplicación debe contar con un frontend o interfaz de usuario para que los usuarios de la aplicación puedan interactuar con la misma</li>
  <li>La aplicación debe contar con un backend o servidor donde está implementada la lógica que conforma los servicios expuestos por la aplicación web</li>
  <li>Los datos que procesa la aplicación se deben almacenar en un motor de base de datos SQL conectado al servidor</li>
  <li>El backend se debe implementar bajo un modelo de capas: handlers, servicios, modelos de datos y repositorio</li>
  <li>La comunicación entre fronted y backend se debe implementar como una REST API (solicitudes y respuestas HTTP) y utilizando el formato de datos JSON</li>
  <li>El proyecto se debe desplegar utilizando tecnología de contenedores (Docker)</li>
</ul>

<b>Entregables</b>

<ul>
  <li>La aplicación se debe implementar en un repositorio de Git alojado en github.com</li>
  <li>En la raíz del repositorio deben existir dos directorios, uno para el frontend y otro para el backend</li>
  <li>Se debe compartir el enlace al repositorio</li>
  <li>Se debe compartir un enlace a un video en el que se documente el funcionamiento de la aplicación (live demo)</li>  
</ul>

## Aplicación Web para reserva de automóviles

### Endpoints expuestos por el servidor

<b>admin</b>
<b>registro</b>
<b>login</b>
<b>reservar</b>
<b>reporte</b>
