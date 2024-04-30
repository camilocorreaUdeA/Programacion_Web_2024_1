# Proyecto final del curso de Desarrollo de Aplicaciones Web 2024-1

## Acerca del proyecto

En el presente documento se presentan las condiciones y productos esperados para la entrega del proyecto final del curso cuyo porcentaje en la evaluación total del mismo corresponde al 25%.

La idea con este proyecto es poner en práctica todos los conceptos aprendidos a largo del semestre:

<ul>
  <li>Definición de la estructura y contenido de un sitio web (cliente de la aplicación web) con el lenguaje HTML</li>
  <li>Aplicación de estilos al sitio web con las reglas de hojas de estilo en cascáda CSS</li>
  <li>Implementación de clientes de API REST con las librerías de Javascript</li>
  <li>Contenido dinámico y lógica del cliente web con Javascript (DOM)</li>
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

### Endpoints sugeridos para ser expuestos por el servidor

Son sugeridos porque últimamente pueden ser los que ustedes como desarrolladores decidan tener en su aplicación web

<b>registro:</b> Esta es una ruta para permitir dar de alta a nuevos usuarios en la aplicación.

<b>usuarios:</b> Esta ruta base debería permitir validar las credenciales de un usuario que ya se ha registrado en la aplicación. También debería estar en capacidad de retornar los detalles asociados a las reservas de un usuario en específico (cantidad de reservas a su nombre en el momento, etc.)

<b>autos:</b> Esta ruta estaría encargada de todo el proceso de búsqueda de los autos para reservar de acuerdo a múltiples criterios como por ejemplo: tipo (sedan o SUV), tipo de transmisión (automática o manual), tipo de combustible (gasolina, gas, híbrido o eléctrico), modelo, marca, etc.

<b>reservas:</b> Esta ruta está asociada a toda la lógica del proceso de reserva. Retorna la lista de automóviles a nombre de un usuario, crea la reserva en la base de datos, etc.

<b>reporte:</b> Esta ruta debería retornar la información consolidada de las reservas de un usuario y el total a pagar para confirmar dichas reservas.

### Funcionamiento básico de la aplicación

En el <code>home</code> o página principal se debe presentar un elemento para que un usuario: 1. se registre en la aplicación, si es la primera vez que la usa, o bien 2. se autentique si es que ya se había registrado previamente.

Cuando un usuario va a registrarse se debe presentar una página en la que se recolecten datos básicos de identidad del usuario. Pero lo más importante es registrar un nombre de usuario y una contraseña que van a ser las credenciales de autenticación de ese usuario en la aplicación.

Cuando un usuario se autentica exitosamente entonces se carga una página de reservas en la que se muestra una sección con las reservas asociadas actualmente al usuario (lista con los automóviles reservados por ese usuario), una sección para hacer una nueva reserva que tiene un buscador que sirve para consultar los automóviles disponibles de acuerdo con algún criterio de búsqueda tipo (sedan o SUV), tipo de transmisión (automática o manual), tipo de combustible (gasolina, gas, híbrido o eléctrico), modelo, marca, etc. La búsqueda debe retornar la lista de automóviles para reservar que se ajustan al/los criterios de búsqueda.

La lista de autómoviles se debe mostrar como una galería de tarjetas en donde en cada tarjeta se puede observar una imagen del automóvil y algunos detalles adicionales (marca, modelo, combustible, etc).

Al seleccionar una tarjeta se debe cargar un menú de opciones en el que se muestren adicionales de la reserva, tales como seguros, asistencia en carretera, silla para bebés, equipo de lujo, etc. Además de un botón Reservar que debe crear la reserva y un botón Cancelar para descartar la reserva de ese autómovil.

Se debe poder contar con una opción para que el usuario elimine una reserva que ya no quiere tener.

Esta página también debe contar con un botón Reporte que al ser presionado carga en una sección un informe con las reservas del usuario y el detalle de los costos de la reserva junto con el precio total que debería ser cancelado para confirmar las reservas.



