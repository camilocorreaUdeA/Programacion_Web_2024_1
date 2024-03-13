# APIs y otras cosas

## API

### ¿Qué es una API?

![Tomado de https://blog.back4app.com/](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/aa6dd98f-35f4-4602-b597-2ff4c48b707d)

Una Interfaz de Programación de Aplicación (Application Programming Interface) es en general un conjunto de funciones, protocolos o funcionalidades encapsuladas en subrutinas que permiten la comunicación e integración entre dos artefactos o productos de software distintos (aplicaciones, servicios web, sistemas operativos, drivers, etc.) que no tienen una interfaz nativa de interacción.

Habitualmente se le conoce a la API como el contrato establecido entre el proveedor y el usuario de un objeto de datos o información. Allí se establece el formato y proceso necesario para que el usuario solicite la información que puede recibir de parte del proveedor. Y también el formato en el que la información es entregada por el proveedor al usuario, es decir la respuesta a la solicitud.

Desde esa perspectiva la API actúa como un mediador que simplifica la comunicación entre el usuario y el proveedor removiendo la necesidad de que alguno tenga que tener  conocimiento de la implementación interna del otro. Por ejemplo, cuando usted instala una aplicación en su computador o dispositivo móvil, y esta hace uso de varios periféricos del sistema, en ese caso su aplicación no necesita saber como acceder y manipular los recursos disponibles en el sistema sino que es este último el que a través del sistema operativo proporciona una serie de funciones que le permiten a la aplicación conectar con los periféricos sin tener que manipularlos directamente, y por el contrario es el sistema operativo quien de acuerdo con las solicitudes de la aplicación va a comunicarse con los drivers del hardware del sistema para acceder a los dispositivos.

![Tomado de https://www.techspot.com/article/2237-what-is-api/](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/8f2b13d2-2f1b-49b4-8c16-47f906544728)

El uso de APIs facilita en gran medida el desarrollo de nuevas aplicaciones y productos de software y además genera ahorros en tiempo y recursos al permitir que la integración con otras aplicaciones y artefactos de software sea transparente, homogenea y no requiera del desarrollo de interfaces a la medida (específicas) para la transferencia de datos y comunicación entre las aplicaciones.

Por ejemplo, suponga que usted desea desarrollar una aplicación para poner filtros (mediante algoritmos de IA) a las fotografías que los usuarios se toman con la cámara del móvil. Y además, necesita que las fotos se suban directamente a la cuenta de Instagram del usuario une vez este lo haya solicitado. Piense por un momento cuánto tiempo tomaría crear una aplicación de estas características si usted estuviera obligado a tener que crear una pieza de software para manipular la cámara del Iphone y otra más para la subida y publicación de las fotos en Instagram. Afortunadamente eso no es necesario ya que lo único en lo que usted debe concentrarse es en el desarrollo de la aplicación porque el sistema operativo del móvil le proporciona las funciones de alto nivel que usted invocar para comunicarse con la cámara del dispositivo. Y por otra parte, la aplicación de Instagram (instalada en el móvil también) proporciona una web API para que con un llamado a una función usted pueda subir y publicar las fotos sin preocuparse del funcionamiento interno de la aplicación de Instagram. 

Clasificación de las APIs

Hasta este punto hemos dado una definición general al concepto de API y se han comentado algunas situaciones ejemplo para ilustrar más claramente el concepto. Ahora daremos un vistazo a la clasificación de las APIs y vamos a empezar por una clasificación general en la cual dividiremos las APIs en dos grandes categorías: Locales y Remotas.

Las APIs locales son precisamente aquellas que hemos venido comentando en los ejemplos anteriores, son interfaces programáticas que actúan como puentes o enlaces que permiten la comunicación a dos artefactos de software que están contenidas o instaladas dentro de un mismo sistema. Estás APis no requieren de un sistema auxiliar de conexión a distancia entre usuario y proveedor sino que al contrario ambos coinciden en compartir un bus de datos físico mediante el cual se pueden comunicar de forma directa. Por ejemplo cuando el mouse envía las señales de pulsado de botones o del sensor de movimiento al sistema operativo para que se vea el cursor moverse y dar clic en pantalla. En este caso puntual el mouse no requiere de ningún sistema auxiliar que le ayude a encontrar la ruta al sistema operativo para enviar su información, y es obviamente porque tanto los drivers del mouse y la pantalla, como el sistema operativo están instalados en el mismo sistema y pueden comunicarse utilizando los medios de comunicación física que este ofrece.

Esta categoría de APIs no es de nuestro interés para los alcances y objetivos del curso.

Las APIs remotas por el contrario si requieren de un agente externo que proporcione los medios y capacidades para que las aplicaciones puedan llevar a cabo el proceso de comunicación, y esto se debe en gran medida a que las aplicaciones residen en sistemas distintos, entonces antes de comunicarse a través de la API se debe establecer una conexión que debe ser mantenida por el tiempo que dure la interacción entre las aplicaciones. El ejemplo más común es cuando usted navega en un sitio web, su navegador web está instalado en su computador mientras que el sitio web está alojado en un servidor que no es más que otro computador localizado en otro lugar. Por tanto, para que su navegador pueda solicitar contenidos al servidor web requiere de establecer una conexión remota al servidor para luego poder hacer la solicitud del contenido del sitio weba través de la API (incluída en el navegador).

Esta categoría de APIs es la que nos interesa para los fines de este curso.

Clasificación de las APIs remotas en términos de acceso:
<ul>
  <li><b>Open/Public API:</b> Son las APIS que están disponibles de manera pública y no existen restricciones para utilizarlas. Públicas no necesariamente significa que su uso sea gratuito.</li>
  <li><b>Partner API:</b> Son APIs que no están disponibles de forma pública y por tanto se necesita contar con permisos o licencias para utilizarlas. Dentro de esta categoría están las APIs que un proveedor proporciona para poder interactuar con los productos de software que vende.</li>
  <li><b>Internal/Private API:</b> Son APIs para usar solo entre componentes internos del sistema y por tanto no se exponen al público. Un buen ejemplo es una API que permite a un servicio de un sistema distribuido hacer una ejecución remota de un procedimiento que está instalado en algún otro servicio perteneciente al mismo sistemaectado mediante una red de comunicación en común.</li>
  <li><b>Composite API:</b> Son APIs que se modelan como una combinación de distintos datos y APIs. Por lo general son secuencias de tareas que corren de manera sincrónica para incrementar la velociad de ejecución de un proceso al poder agrupar y ejecutar varias solicitudes en un solo llamado a la API. La ventaja de estas APIs es que mejora el desempeño del cliente ya que combina múltiples solicitudes a distintas APIs en una única solicitud y de igual manera para las respuestas.</li>
</ul>

Clasificación de las APIs remotas de acuerdo con su arquitectura:
<ul>
  <li><b>API REST:</b> REST es el acrónimo de "Representational State Transfer", este tipo de API permite que un cliente solicite recursos a un servidor utilizando los métodos HTTP (GET, POST, PUT, PATCH y DELETE) Para este propósito el servidor proporciona "endpoints" identificados por URLs y que tienen asociada una operación sobre los recursos del servidor. En REST la API depende o está basada en los recursos, es decir, el cliente solicita recursos y el servidor los proporciona.</li>
  <li><b>API GraphQL:</b> El lenguaje GraphQL permite desarrollar APIs que ofrecen acceso a datos vinculados que se pueden pensar como un grafo de entidades, por lo tanto se pueden crear solicitudes fuertemente tipadas en las que se especifica de manera exacta los datos que se necesitan. Luego, mediante un solo "endpoint" expuesto por el servidor se obtienen datos o recursos de cualquier tipo o naturaleza. Al contrario de las REST APIs en las que por cada recurso y operación se tienen distintos "endpoints".</li>
  <li><b>API SOAP:</b> SOAP es el acrónimo de "Simple Object Access Protocol" que es un estándar de mensajería que se basa en el formato XML para la construcción de los mensajes de solicitud y respuesta. SOAP utiliza el patrón de RPC (Remote Procedure Call) donde la comunicación entre cliente y sevidor se modela como el llamado a funciones o métodos a los que se les pasan parámetros específicos para que se retorne un resultado con los datos esperados.</li>
  <li><b>API RPC:</b> La arquitectura RPC propone un enfoque en el cual cargas de trabajo remotas se ejecutan como si fuera de manera local, esta característca hace que RPC sea muy útil en la construcción de servicios distribuidos ya que una aplicación se puede componer de rutinas o funciones que internamente hacen llamados a funciones que se ejecutan en otros servicios externos. Esto también significa una gran felxibilidad en comparación con arquitecturas como REST poque se puede ejecutar practicamente cualquier función y solicitar cualquier tipo de acción sobre los datos, mientras que REST está restringido a los métodos ofrecidos por el protocolo HTTP.</li>
</ul>

Para los fines y objetivos de este curso nos interesa enfocarnos solo en las APIs REST así que desde este punto y en adelante cada vez que usemos el término API nos estaremos refiriendo específicamente a REST API.

### REST APIs

![Tomado de https://rapidapi.com/blog/api-glossary/api/](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/7ef4cb5e-77c4-410f-90d2-a90f221e5711)

Las REST APIs comunican un cliente y un servidor a mediante el intercambio de mensajes del protocolo HTTP (HyperText Transfer Protocol). Las operaciones asociadas al protocolo HTTP permiten ejecutar acciones como crear, leer, actualizar y eliminar, sobre los recursos almacenados en (o a los que tiene acceso) el servidor. Por ejemplo para leer/obtener un recurso a través de una REST API se tendría que usar el método GET del protocolo HTTP, para crear un nuevo recurso el método POST (en casos PUT), para actualizar están disponibles los métodos PUT y PATCH y para eliminar un recurso está el método DELETE.

Hay que tener en cuenta que cuando hablamos de obtener o leer un recurso no estamos haciendo referencia al recurso como tal sino a una representación del estado del mismo en el momento de la solicitud, de ahí viene el mismo acrónimo de REST, Representational State Transfer, transferencia del estado de representación. Y una representación de un recurso no es el recurso mismo sino su estado (valores de los datos) en un momento en particular. 

La representación del recurso puede ser presentada por la API en distintos formatos, practicamente en cualquiera, entre las más populares se incluyen JSON (Javascript Object Notation), HTML, XML, Python, PHP, texto plano, mapas de bits, entre otras. Siendo la más popular de todas JSON ya que es legible para los seres humanos (human readable), fácil de codificar/decodificar en objetos por los lenguajes de programación (del lado del cliente y del servidor) y además es agnóstico del lenguaje de programación.

En REST APIs que estén bien diseñadas se incluye el uso de encabezados y parámetros en las solicitudes, ya que permiten incluir información importante como metadatos, credenciales de autenticación y autorización, las rutas a los recursos o URIs, datos para caching y cookies entre otros. Y en el caso de las respuestas también se considera importante incluir los códigos de estado de la solicitud HTTP.

### 6 condiciones para que una REST API sea RESTful

<ul>
  <li><b>Interfaces uniformes:</b> Todas las solicitudes al mismo recurso deben ser iguales sin importar desde dónde se haga la solicitud. Se debe asegurar que cualquier dato pertenece a un solo identificador uniforme de recursos URI. Se recomienda también que los recursos no sean muy grandes pero que contengan la información que necesita el cliente.</li>
  <li><b>Desacople entre cliente y servidor:</b> El cliente y el servidor deben ser sistemas independientes. La única información que el cliente debe conocer del servidor son las URIs a los recursos y aparte de los "endpoints" proporcionados por el servidor no debe tener ningún otro canal de interacción con este. Y por el lado del servidor, este no debe poder modificar nada en la aplicación del cliente y su única función de cara al cliente es la de responder con los datos solicitados por este.</li>
  <li><b>Ausencia de estado en las solicitudes:</b> Cada solicitud requiere compartir toda la información necesaria para su procesamiento. No deben existir sesiones de lado del servidor y este no debe almacenar ningún dato del cliente.</li>
  <li><b>Caché-abilidad:</b> Para incrementar el desempeño del lado del cliente y la escalabilidad del lado del servidor se pueden utilizar cachés de los recursos en el lado del servidor. El servidor debe informar al cliente que recursos pueden ser cacheables.</li>
  <li><b>Arquitectura de capas:</b> En el canal de comunicación entre cliente y servidor pueden existir muchas capas intermediarias, por tanto se deben diseñar las APIs para que tanto cliente como servidor no tengan en cuenta si en realidad se están comunicando de manera directa con una apliación intermediaria.</li>
  <li><b>Código bajo demanda:</b> En caso de que una API se utilice para enviar código que pueda ser ejecutado se debe restringir está ejecucuión a que solo sea bajo demanda, solicitud o autorización explícita.</li>
</ul>

Estos principios hacen que REST sea un estilo arquitectónico flexible para la construcción de sistemas orientados a servicios basados en estándars web. REST proporciona beneficios entre los que se incluyen escalabilildad, reusabilidad, y bajo acoplamiento los cuales le permiten ajustarse a las necesidades de las aplicaciones web modernas que atienden a millones de usuarios.

Hemos mencionado en varias ocasiones dos conceptos para los cuales no hemos dado una definición concreta. Veamoslas a continuación:

<b>JSON</b>: Javascript Object Notation es un formato basado en texto que se usa para representar datos estructurados en la sintaxis de objetos de Javascript. Es comúnmente utilizado para transmitir datos en aplicaciones web (como formato de datos para el cuerpo de solicitudes y respuestas HTTP). Al ser un formato de uso extendido casi todos los lenguajes de programación cuentan con funciones auxiliares que permiten convertir un objeto de datos a JSON y desde JSON a la representación en un objeto de datos del lenguaje de programación.

Es importante aclarar que JSON solo se utiliza para representar los datos de un objeto, los métodos del objeto no pueden ser representados con JSON.

Los objetos de JSON pueden ser almacenados en archivos de texto con la extensión <code>.json</code>, y para transmitir JSON en solicitudes/respuestas HTTP se recomienda adjuntar el encabezado <code>Content-Type</code> con el valor <code>application/json</code>.

Estructura de un objeto JSON:

```json
{
  "propiedad":"valor_en_texto",
  "propiedadNumerica":25,
  "propiedadBooleana":true,
  "unArreglo":[
     "header",
     "nav",
     "section",
     "footer"
  ],
  "unObjeto":{
     "propiedad":"valor"
  }
}
```
Aspectos a tomar en cuenta acerca de la estructura de JSON:

<ul>
  <li>Para los nombres de las propiedades y los valores que sean texto es estricto el uso de comilla doble, la comilla sencilla no es válida.</li>
  <li>Las propiedades se separan con coma, la última propiedad del objeto no debe llevarla.</li>
  <li>Los arreglos se representan como una lista de valores separados por coma y encerrados en corchetes.</li>
  <li>Un objeto JSON puede tener propiedades cuyo valor son representaciones de otros objetos, es decir, se permite la anidación de objetos en JSON. Los objetos van encerrados dentro de llaves (curly braces)</li>
  <li>Para los nombres de las propiedades que se compongan de dos o más palabras se recomiendan los estilos <code>camelCase</code> y <code>snake_case</code>. Evite usar espacios o caracteres especiales para separar las palabras. Evite también mezclar ambos estilos de escritura.</li>
  <li>No debe existir más de una propiedad con el mismo nombre</li> 
  <li>Todo el objeto va delimitado por llaves (curly braces)</li>
</ul>

<b>Endpoint</b>: Los endpoints son los extremos en el sistema de comunicación de una API, por lo general se asume que el endpoint incial es cliente, y el cliente puede ser un navegador, una interfaz de usuario o frontend, o bien un servicio que solicita datos a otro. Los endpoints finales están del lado del servidor y son proporcionados por los servicios que están allí alojados y están encargados de realizar una operación determinada sobre un recurso en específico.

En las APIs REST un endpoint se identifica mediante una URL que determina la localización de un recurso alojado en el servidor. Cada URL identifica un endpoint y cada endpoint representa una operación que se pueda ejecutar sobre el recurso señalado por la URL. Es decir, para cada método HTTP u operación autorizada sobre un recurso se destina un endpoint (no necesariamente una URL) individual. Por tanto, dos endpoints pueden tener la misma URL pero depediendo del método HTTP utilizado en la solicitud se ejecuta una función u otra en el servidor.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/db3b1766-fd16-4793-b57c-d1a2b43a8ef5)

Endpoint: <code>resources/videos</code>

### CRUD: Create, Read, Update and Delete

CRUD son las siglas de las palabras <code>CREATE</code>, <code>READ</code>, <code>UPDATE</code> y <code>DELETE</code> que son los términos que describen las cuatro operaciones básicas o esenciales para trabajar con datos almacenados en una persistencia o base de datos.

Antes de seguir detallando a mayor profundidad los conceptos asociados a CRUD, hagamos un paréntesis para comentar acerca de la arquitectura de una aplicación web, o más bien de las aplicaciones web que dentro del alcance de aprendizaje de este curso.

A partir de este momento vamos a relacionar el término aplicación web con una arquitectura de software de varios niveles, del término en inglés <i>multi-tier</i>, donde interactúan por lo menos 3 capas distintas a saber: 

<ul>
  <li>La interfaz de usuario o UI: También conocida como frontend es la interfaz con la que interactúa directamente el usuario a través de un navegador web. Allí conviven el lenguaje de marcación HTML, las hojas de estilo CSS y el lenguaje de programación Javascript. Esta capa es la encaragada de procesar los datos introducidos por el usuario y disparar la ejecución de las operaciones de la aplicación mediante el envío de peticiones/solicitudes al servidor de la aplicación. Luego recibe las respuestas y las procesa de modo que el usuario las pueda visualizar en un formato previamente determinado.</li>
  <li>Servicio web: Es la primera instancia de lo que se conoce como el backend, o lado del servidor, de la aplicación web. Puede ser un solo servicio o múltiples servicios interconectados, que simplemente son rutinas de software que definen las funciones que van a procesar las peticiones/solicitudes recibidas desde la interfaz de usuario. En el servicio web se definen los distintos endpoints que componen la aplicación web y les asigna la lógica necesaria para cumplir con la operación para la que fueron diseñados.</li>
  <li>Persistencia o base de datos: Es otra instancia dentro del backend de la aplicación web, y es simplemente el repositorio donde se almacenan los datos que procesa la aplicación. No tiene canal de interacción directa con la interfaz de usuario y cualquier operación sobre los datos que almacen se da como consecuencia de la ejecución de la lógica en los servicios web. El acrónimo CRUD se refiere precisamente a las operaciones que se pueden ejecutar en esta capa.</li>
</ul>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/edda1967-641b-4e97-a34d-303fba9e10c2)

Las operaciones CRUD en una API REST se implementan sobre la base de las peticiones y respuestas de los métodos HTTP, por tanto vamos a hacer un pequeño recuento de la anatomía de las peticiones y respuestas en el protocolo HTTP.

Las peticiones y respuestas del protocolo HTTP comparten una estructura similar compuesta por las siguientes partes:

<ol>
<li>Comienzo (start line): Es la primera línea de texto que describe la petición que se va a hacer, o bien, indica el estatus de la respuesta para saber si la petición se atendió correctamente o no. Este comienzo es una única línea de texto.</li>

<li>Encabezados (headers): Conjunto de encabezados, cada uno en su propia línea de texto individual, que describen detalles específicos de la petición o la información incluída dentro del cuerpo del mensaje (tanto de una petición como de una respuesta).</li>

<li>Línea en blanco (empty/blank line): Es obligatoria y se agrega para separar la parte informativa (descriptiva) del mensaje del contenido incluído en el cuerpo del mensaje.</li>

<li>Cuerpo del mensaje (body): Es opcional en algunos casos ya que no todos los métodos HTTP requieren de un cuerpo con datos asociados a la petición o la respuesta. La presencia del cuerpo en el mensaje sugiere que además existen encabezados que detallan la longitud y el tipo de datos de la información allí incluída.</li>
</ol>

Peticiones HTTP (Requests)

<ol>
<li>Comienzo (start line)

El comienzo de una petición HTTP debe tener 3 elementos:

<ol>
<li>Un método HTTP, es un verbo, GET, o un nombre, OPTIONS, que describe la operación a realizar.</li>
<li>El destinatario de la petición, es por lo general una URL cuyo formato varía de acuerdo con el método HTTP usado en la petición. Puede ser una ruta absoluta al recurso en el servidor, la URL completa o en algunos casos un asterisco que representa al servidor como un todo.</li>
<li>La versión del protocolo HTTP, indica la versión del protocolo en la que se espera establecer el intercambio de mensajes con el servidor.</li>
</ol>
</li>

<li>Encabezados (headers)

Los encabezados son pares de clave y valor separados por dos puntos (:) y que van en líneas individuales. Las claves y valores no diferencian entre mayúsculas y minúsculas.

Se pueden subdividir en 3 categorías: 

<ol>
<li>Encabezados generales: Aplican al mensaje como un todo.</li>
<li>Encabezados de petición: detallan a mayor profundidad la petición, dando contexto de la misma y en ciertos casos restringiendo su campo de actuación.</li>
<li>Encabezados de representación: Describen el formato original de los datos que van en el cuerpo del mensaje, por tanto sólo tienen sentido cuando la petición tiene un cuerpo.</li>  
</ol>
</li>

<li>Línea en blanco (empty line)</li>

<li>Cuerpo (body)

Solo utilizan esta parte de la petición los métodos POST, PUT y PATCH. Existen 2 categorías:

<ol>
<li>Cuerpo de un único recurso: Un solo archivo definido por dos encabezados que son “Content-Type” y “Content-Length”.</li>
<li>Cuerpo multi-recurso: Son varios archivos que contienen información de distintos recursos, están generalmente asociados a las HTML Forms.</li>
</ol>
</li>
</ol>

Respuestas HTTP (Responses)

<ol>
<li>Comienzo (start line)

Más conocido como línea de estatus, contiene la siguiente información:

<ol>
<li>La versión del protocolo HTTP, usualmente HTTP/1.1</li>
<li>El código de estatus para indicar si la petición se procesó correctamente o no.</li>
<li>El texto informativo del estatus, no es más que una breve descripción textual del estatus de la petición.</li>
</ol>
</li>

<li>Encabezados (headers)

Aplica para la respuesta la misma descripción vista para el caso de las peticiones.
</li>

<li>Línea en blanco (empty line)</li>

<li>Cuerpo (body)

Es opcional porque no todas las respuestas a peticiones HTTP requieren de un cuerpo que tenga datos, un buen ejemplo es la respuesta a una petición POST. Aquí aplican las mismas categorías vistas para el caso de las peticiones y se adiciona la siguiente:

Cuerpo de un único recurso: siendo este un archivo de longitud desconocida y codificado en “chunks”, es decir, que el encabezado con clave “Transfer-Encoding” viene con el valor “chunked”.
</li>
</ol>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/cec82073-8297-48c1-aebf-5be0905784d1)

Operaciones CRUD

1. <code>CREATE</code>: Esta operación permite crear un nuevo recurso en la aplicación web, es decir, agrega un nuevo registro en la base de datos. Si es una base de datos SQL agrega una nueva fila en la tabla respectiva, o si es una tabla NoSQL agrega una nueva instancia de la estructura de datos utilizada en la base de datos, por ejemplo en MongoDB sería un nuevo documento.

Para implementar esta operación en una API REST se utiliza el método HTTP <code>POST</code> y en el cuerpo de la solicitud se pasa el nuevo objeto que será agregado en la base de datos, ese objeto debe ir codificado en el formato escogido para compartir datos entre los endpoints de la aplicación. El formato que utilizaremos en el curso es JSON.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/1bde9247-1944-4932-a0ba-c078d9ebeb72)

2. <code>READ</code>: Esta operación permite consultar la representación de los recursos de la aplicación web. Retorna registros almacenados en la base de datos.

La implementación de esta operación se logra a través del método HTTP <code>GET</code> y puede ser de dos formas, lectura de una representación individual de un recurso, o leer una lista con las representaciones de los recursos disponibles en la aplicación web. Para el primer caso se debe especificar en la URL un dato que distinga completamente al recurso de otros, por lo general se utiliza el id que se le asignó al almacenarlo en la base de datos. Para el segundo caso no es necesario especificar un distintivo ya que se van a retornar una lista de los recursos existentes, más adelante veremos que podemos limitar la cantidad de recursos en la lista utilizando algunos parámetros en la solicitud (filtros).

Un solo recurso:

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/9ee9a0f3-21c1-4d28-b293-d4945e20d060)

Lista de recursos:

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/3cb17988-cd37-4fd0-85c6-9626e014b8e1)

3. <code>UPDATE</code>: Esta operación que puede ser implementada con los métodos <code>PUT</code> y <code>PATCH</code> permite editar o actualizar un recurso de la aplicacion web. Recuerde que el método <code>PUT</code> crea un nuevo recurso en caso de que el que se haya solicitado modificar no exista. Y además, solo con el método <code>PATCH</code> es posible pasar solamente los campos de datos sujetos a la actualización. También se debe especificar en la URL un dato que distinga completamente al recurso de otros, por lo general se utiliza el id que se le asignó al almacenarlo en la base de datos.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/5a5aa95b-5aff-4d7c-a42d-399ff38068a1)

5. <code>DELETE</code>: Con esta operación se eliminan registros de la base de datos. Para implementarla se utiliza el método HTTP <code>DELETE</code>. Se debe especificar en la URL un dato que distinga completamente al recurso de otros, por lo general se utiliza el id que se le asignó al almacenarlo en la base de datos.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/ab034419-0289-418d-abf4-46c566743493)

