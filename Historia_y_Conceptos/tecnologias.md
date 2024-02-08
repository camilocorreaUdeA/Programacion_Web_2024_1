### ¿Cómo funciona la Web?

## ¡Y así funciona Internet!

<img width="599" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/84f7aeb7-d36b-4643-b7d7-6f7bf7748213">

### Capas del modelo de comunicación

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/89b0e275-529c-429a-bd63-1a09cf47c2d6)

<b>Aplicación</b>

Esta capa la utilizan las aplicaciones de red, que así se denomina a las aplicaciones que dependen de estar conectadas a Internet para su correcto funcionamiento (navegadores web, mensajería instantánea, clientes de correo electrónico, música y entretenimiento, etc.)

Para que los mensajes (datos) producidos por las aplicaciones de red puedan transmitirse en la red se necesita de un conjunto de protocolos que dependiendo de la naturaleza de la aplicación proporcionen una interfaz entre la capa de aplicación y las capas subyacentes del modelo de comunicación. Los protocolos más conocidos son HTTP, HTTPS, FTP, SFTP, POP3, SMTP y estos en conjunto con otros forman la base de servicios de red como transferencia de archivos, navegación web, correo electrónico, mensajería instantánea, etc.

<b>Transporte</b>

La capa de transporte garantiza la integridad y confiabilidad de la comunicación mediante la segmentación de los mensajes de datos, el control de flujo y errores de la comunicación.

En el proceso de segmentación se divide los mensajes provenientes de la capa de aplicación en unidades de datos más pequeños llamados segmentos que están caracterizados por puertos de origen y destino para saber desde y hacia qué aplicación de red van dirigidos los datos y un número de secuencia para reconstruir los mensajes segmentados en el destino.

En el proceso de control de flujo la capa de transporte controla la cantidad de datos que se transmiten de forma que se aproveche en su totalidad la máxima tasa de datos entre origen y destino sin incurrir en errores o pérdidas de información.

En el proceso de control de errores la capa de transporte utiliza un esquema conocido como Automatic Repeat Request para solicitar por segmentos faltantes que se hayan pérdido en tránsito y mediante un esquema de checksum verifica que los segmentos no hayan sido modificados o corrompidos.

Los protocolos de la capa de transporte son Transmission Control Protocol (TCP) y User Datagram Protocol (UDP). TCP está orientado a la conexión mientras que UDP no.

<b>Red o Internet</b>

La capa de red es la encargada de transmitir los segmentos de la capa de transporte a un computador destino ubicado en una red distinta a la del origen de los datos.

Esta capa se encarga del direccionamiento lógico (utilizando direcciones IP), la determinación de la mejor ruta y posterior enrutamiento de los datos. Hacen parte de la capa de red los protocolos de enrutamiento OSPF, RIP, IS-IS, BGP entre otros. Y también protocolos como ICMP y ARP.

<b>Enlace (Acceso al medio)</b>

En la capa de enlace se hace la transmisión de tramas de datos entre dos dispositivos ubicados dentro de la misma red. Las direcciones físicas de las tarjetas de red (NIC), también conocidas como direcciones MAC o de acceso al medio.

Existen estrategias de acceso al medio como CSMA/CD y TDMA que la capa de enlace aprovecha para controlar cómo las tramas de datos se ubican y reciben en el medio de transmisión que puede ser confinado como cables de cobre y fibra óptica o no confinado como el aire libre o el vacío.

<b>Física</b>

Los bits que componen una trama de datos son convertidos en señales eléctricas, de luz o radiadas por la capa física dependiendo del medio por el que se vayan a transmitir los datos. La capa física se encarga de la modulación y codificación de las señales electromagnéticas que representan las secuencias binarias de las tramas de datos.

<img width="608" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/bb531161-2e1e-45ef-b8a9-168ae12cc5a5">

### TCP/IP

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/3ddffc4d-a1c9-4c54-8fe4-aa3dfa446f94)

<b>El protocolo de Internet (IP)</b>: Este protocolo es el encargado de enrutar o entregar los paquetes de datos entre una aplicación o dispositivo que es el origen de los datos y un destinatario (otra aplicación o dispositivo). El protocolo se asegura de que los paquetes lleguen al destino correcto, para este fin define un conjunto de reglas y formatos a los cuales deben ajustarse las aplicaciones y dispositivos para poder comunicarse e intercambiar datos en una red o a través de una interconexión de redes. 

<b>Protocolo de control de la transmisión (TCP)</b>: TCP es un protocolo orientado a la conexión que asegura comunicaciones íntegras y confiables entre los dispositivos (aplicaciones) interconectados en una red. Este protocolo proporciona todas las funciones necesarias para el transporte confiable de los datos en una red, siendo la más relevante el mecanismo llamado Positive Acknowledgment with Re-transmission (PAR) que sirve para hacer la re-transmisión automática de datos que se corrompen o se pierden mientras que están en tránsito a través de la red.

<b>3-way Handshake (Apretón de manos de 3 vías)</b>

Este es el proceso para establecer una conexión entre dos actores, un dispositivo (aplicación) abierto-activo conocido como el cliente y otros en modo abierto-pasivo conocido como el servidor.

En el primer estado de este proceso el cliente envía al servidor un segmento TCP denominado SYN (Synchronize Sequence Number) mediante el cual le informa que desea iniciar una comunicación con el número de secuencia indicado en el segmento.

En el segundo paso el servidor responde al cliente con un segmento SYN-ACK, donde responde a la secuencia enviada por el cliente y a su vez propone una secuencia para comenzar el intercambio de datos.

En el tercer y último paso el cliente responde al servidor con un segmento ACK que continúa la secuencia numérica enviada por el servidor, es en este momento en que se considera que la comunicación ha sido establecida.

El 3-way handshake no es un proceso exclusivo de la fase de establecimiento de la conexión. Es en realidad el pilar fundamental para que el intercambio de datos en TCP sea realmente confiable ya que cualquier dispositivo va a reenviar los datos hasta no recibir un acuso de recibo o ACK de parte del otro dispositivo involucrado en la comunicación.

En la capa de transporte se verifica la integridad de los datos mediante un algoritmo de verificación de errores ejecutado por el receptor de los datos. Si el receptor considera que los datos recibidos han sido alterados entonces procede a descartar ese segmento de modo que el dispositivo origen de los datos debe enviar de nuevo hasta que reciba un acuso de recibo positivo.

<b>Proceso completo de transferencia de datos en TCP</b>

El proceso de intercambio de datos en TCP está compuesto en general por 3 fases:

Fase 1: En esta ocurre el 3-way handshake que permite que los dispositivos se conozcan y establezcan una relación de confianza.

Fase 2: Los dispositivos intercambian información usando las secuencias numéricas acordadas en la primera fase y los acuso de recibo correspondientes.

Fase 3: Una vez se termina el intercambio de datos el servidor procede a cerrar la conexión utilizando un segmento FIN al cual el cliente debe dar acuso de recibo y a su vez el servidor debe dar acuso de recibo de la contestación del cliente para dar por terminada la conexión.

### DNS (Domain Name System)

DNS funciona semejante a como funcionaban los antiguos directorios telefónicos, solo que en este caso un nombre de dominio está asociado a una dirección IP que lo identifica en Internet.

Para los seres humanos es más fácil recordar nombres de dominio (también llamados direcciones web) que direcciones IP, luego, para visitar una página web simplemente ingresamos la dirección del sitio web en la barra de búsqueda del navegador.

<ol>
<li>El navegador revisa en su caché interna si ya tiene una dirección IP asociada al nombre de dominio.</li>
<li>El sistema operativo revisa si hay un match en la memoria caché del sistema.</li>
<li>El sistema operativo hace una consulta al “Resolver” que es un servicio que provee el proveedor de servicios de Internet o ISP. El Resolver también revisa su caché para encontrar un match.</li>
<li>El Resolver hace una petición el Root Name Server que responde con la dirección IP del servidor de Top-Level Domain (TLD) que tiene información acerca de los dominios de segundo nivel asociados a los dominios .COM, .EDU, .GOV, .ORG, etc.</li>
<li>Si el nombre de dominio tiene un subdominio entonces el servidor de dominios de segundo nivel proporciona la información acerca de los servidores de subdominio (authoritative name servers) asociados a ese dominio de segundo nivel. Ejemplo: blog.logrocket.com</li>
<li>El servidor de subdominio (authoritative name server) responde al Resolver con la dirección IP asociada a la dirección web.</li>
<li>El Resolver actualiza su caché y devuelve la dirección IP al cliente que la solicitó.</li>
<li>El cliente actualiza su caché, pasa la dirección al navegador y este actualiza su propia caché y hace la petición web a la dirección IP.</li>
</ol>

<img width="603" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/9351239d-bc1e-4775-8d99-7ffd641a6aec">

### ¿Y qué es el hipertexto?

<img width="387" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/f1c96d75-610b-45e4-9041-9bfa83195c03">

El Hipertexto es simplemente texto que contiene un enlace (hiperenlace) que conecta con otro texto. El hipertexto es la piedra angular y motor de la web. El hipertexto está intrínsecamente ligado a la web y sin este, está última no podría funcionar tal y como la conocemos hoy en día.

El hipertexto permite enlazar la información de una forma no lineal o secuencial pero asociativa al mismo tiempo. Ese tipo de enlazamiento permite que quien lea un hipertexto tenga la posibilidad de hacerlo en el orden que mejor le parezca. Además, permite que los enlaces en los textos (o documentos) puedan enlazar a cualquier otro tipo de recurso (aplicaciones, audios, videos, animaciones, etc.) u otros textos de diferentes autores (enlaces externos) distribuidos en distintos lugares.

De acuerdo con lo anterior el hipertexto se constituye en una forma única de comunicación humana que posibilita nuevas maneras para crear significado.

<b>¿Cómo funciona el hipertexto en una página web?</b>

Para crear hipertexto se necesita primero de un hiperenlace, este último se puede incrustar en una página web utilizando el lenguaje <strong>HTML</strong> (<i>Hypertext Markup Language</i>) el cual tiene una etiqueta llamada “ancla” (anchor) que permite asociar una URL a un texto. De este modo se incrusta en el texto una referencia digital a otra página web y al dar clic al texto se accede a la página referenciada.

Al hacer clic en el hipertexto se da inicio a una petición al servidor que aloja el sitio web, utilizando el protocolo <strong>HTTP</strong>, generalmente, una petición <strong>GET</strong> que permite descargar desde el servidor al navegador el contenido de la página web (el código HTML).

<img width="597" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/9e121ca5-ee23-41b2-a57a-e32cd87a28c3"><br>
a: Etiqueta ancla (anchor) de HTML<br>
href: referencia a un contenido o recurso en la web

### URL: Uniform Resource Locator (Localizador Uniforme de Recursos)

<img width="569" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/27125788-876e-47b4-8ce6-7c82749be4af">

Una URL es la dirección única de un recurso alojado en la web. Cuando se hace una petición a un servicio web se debe especificar una URL para que el servidor web pueda identificar las características del servicio que es solicitado.

<b>Partes de una URL</b>

Esquema: Se refiere al protocolo que se utiliza para hacer la petición. En el caso de hipertextos se utiliza el protocolo HTTP (Hypertext Transfer Protocol).

Nombre de dominio: Es la dirección web que identifica al recurso en la web y que está asociada a una única dirección IP.

Ruta al recurso: Identifica el lugar exacto dentro del servidor donde se encuentra el recurso en específico que se está solicitando. Dicho de otra manera es la ruta a través de los directorios donde está almacenado el recurso (la página web).

Parámetros de la solicitud: Son opcionales, se señalizan con el símbolo de interrogación (?) y se separan con el símbolo ampersand (&). Son parejas de clave y valor que se utilizan para solicitar al servidor que realice acciones específicas sobre los resultados que debería retornar la petición.

Los usos más comunes de los parámetros de solicitud son:

<ol>
<li>Rastreo (tracking): Son parámetros pasivos, es decir que son agregados automáticamente en la solicitud y no directamente desde el navegador. Se utilizan para rastrear detalles del cliente que hace la solicitud, como por ejemplo si pulsó algún enlace para llegar a la página, la geolocalización, etc.</li>
<li>Ordenamiento (sorting): Son parámetros activos que permiten indicarle al servidor que ordene los resultados en un orden determinado.</li>
<li>Búsqueda (searching): Para indicarle al servidor que utilice su motor interno de búsqueda.</li>
<li>Identificación (identifying): Para retornar resultados que cumplan con una característica determinada que los diferencia de otros.</li>
<li>Paginación (pagination): Sirven para dividir largas listas de resultados en bloques o series de menor longitud.</li>
<li>Traducción (translation): Para indicar el idioma en el que se quiere recibir los resultados o el recurso.</li>
<li>Filtrado (filtering): Para retornar solo resultados que cumplan con un criterio específico.</li>
</ol>

<img width="589" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/0e3494c8-fd4e-4cd4-884a-10f1915902fb">

Ancla: Esta parte de la URL también es opcional y sirve para ubicarse en un punto en específico del contenido de la página web. Vale la pena mencionar que dicho punto debe estar señalizado por una etiqueta HTML llamada ancla nombrada (named anchor) para poder ser alcanzable. Se utiliza el símbolo numeral (#) para indicar el nombre del ancla en la URL.

### HTTP: HyperText Transfer Protocol (Protocolo para la transferencia de hipertexto)

HTTP es un protocolo (sistema de reglas, formatos y estándares que definen cómo se deben compartir/intercambiar datos en redes de computadores) de Internet que se utiliza para interactuar con recursos que están alojados en servidores que están interconectados en la red.

HTTP es un protocolo para conexiones bajo arquitectura cliente-servidor, donde el cliente, usualmente un navegador web, realiza una solicitud a un servidor web que almacena un recurso que le interesa consultar que puede ser desde un documento HTML, una base de datos, una imagen, un video, etc.

En HTTP los clientes y servidores se comunican mediante el intercambio de mensajes individuales. A los mensajes enviados por el cliente se les conoce usualmente como peticiones (requests) mientras que a los mensajes enviados por el servidor, en respuesta a las peticiones del cliente, se les conoce simplemente como respuestas (responses).

Para mostrar la página web en pantalla el navegador web (cliente) envía una petición solicitando un documento HTML al servidor. Cualquier otro script presente (o referenciado) en el documento HTML ejecutado por el navegador y que necesite hacer solicitudes adicionales al servidor para actualizar el contenido de la página web lo hará a través de nuevas peticiones HTTP al servidor.

<img width="414" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/2da47129-b3fe-4353-aefb-8cf933b1d41d">

Breve descripción del flujo de una petición HTTP:

<ol>
<li>Se ejecuta la resolución del nombre de dominio: Cuando se da enter al ingresar la URL en la barra de búsqueda del navegador o al dar clic en un hiperenlace de una página web (cuando se enlaza con una web externa) lo primero que el cliente hace es encontrar la IP asociada al nombre de dominio. Esto se logra con ayuda del protocolo DNS, cuyo funcionamiento se revisó a detalle con anterioridad.</li>
<li>Apertura de una conexión TCP: Una vez se ha resuelto el nombre de dominio con el protocolo DNS se procede a establecer una conexión TCP con el servidor (three-way handshake).</li>
<li>El cliente envía una petición HTTP al servidor: El cliente utiliza un mensaje HTTP (usualmente usando la versión 1.1 del protocolo HTTP) para solicitar al servidor la página web que se desea mostrar en el navegador. En este caso la petición HTTP es del tipo GET, más adelante se profundizará en los distintos tipos de peticiones HTTP que el cliente puede hacer al servidor.</li>
<li>El servidor envía la respuesta: El servidor responde a la petición del cliente con una mensaje HTTP en el cual viene el código HTML de la página web (hipertexto).</li>
<li>El navegador despliega la página web: El navegador web procesa el documento en HTML y despliega la página web en la pantalla.</li>
</ol>

### Y los métodos HTTP?

<img width="600" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/d9a7c840-197c-4ec6-837f-d382cbc96964">

Un método HTTP es un verbo o un nombre que define una operación que el cliente quiere ejecutar en el servidor. Usualmente un cliente busca realizar una de estas dos operaciones: 1) obtener un recurso (una página web) alojado en el servidor; o 2) enviar datos al servidor que fueron recolectados en un formulario de la página web. Pero además de esas dos operaciones existen otras más que cuentan con sus propios casos de uso.

<img width="605" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/01feef70-6a6d-4d03-8b33-d198fec54be0">

<b>GET</b>: Este método HTTP se utiliza para solicitar una representación de un recurso en específico. En respuesta a la petición, el servidor devuelve la representación del recurso en un formato concreto: HTML, XML, JSON, imagen, video, etc. Una petición GET debería ser utilizada para solicitar datos solamente y no debería incluir ningún dato en el cuerpo de la solicitud. Es una petición idempotente, es decir, el resultado de realizar varias peticiones de este tipo es siempre el mismo, como cuando se navega a una URL, el resultado es siempre cargar la página en el navegador. 

<b>POST</b>: Este método permite enviar datos al servidor en el cuerpo de la petición. Tanto el tipo como la cantidad de datos que se envían se pueden especificar en encabezados de la petición. Generalmente, el resultado de una petición POST es un cambio en los datos que almacena el servidor. El uso más común de este tipo de método es cuando se recolectan datos para una suscripción mediante un formulario en un página web.

El servidor responde a una petición POST indicando que un nuevo recurso ha sido creado (que los datos pasados en la petición fueron correctamente almacenados).

<b>PUT</b>: Este método crea un nuevo recurso o reemplaza la representación de un recurso objetivo con los datos que vengan en el cuerpo de la petición. La diferencia con el método POST es que PUT es idempotente, es decir, cuando crea un nuevo recurso solo lo hace una vez y repetir la misma petición no creará recursos adicionales, mientras que POST si crea un nuevo recurso cada vez que se hace la misma petición.

El servidor responde a una petición PUT indicando que un nuevo recurso ha sido creado (que los datos pasados en la petición fueron correctamente almacenados). O en caso de hacer una modificación a un recurso, el servidor indica si la modificación se hizo correctamente o no.

<b>DELETE</b>: Este método permite eliminar un recurso en específico. El servidor responde indicando si la operación se realizó correctamente o si no había ningún recurso que cumpliera con la petición.

<b>PATCH</b>: Este método sirve para hacer modificaciones parciales a un recurso en específico. La diferencia con el método PUT es que este último modifica o sobreescribe completamente la representación del recurso, mientras que PATCH solo modifica los datos que se especifican en el cuerpo de la petición.

<b>HEAD</b>: Este método funciona igual que el método GET con la diferencia de que en este caso el servidor no devuelve la representación del recurso en la respuesta. Esta petición le indica al servidor que el cliente solo está interesado en saber si el recurso es existe y algunas de sus características expresadas en los encabezados (headers) HTTP, más no en el recurso en sí mismo.

<b>TRACE</b>: Este método sirve para diagnosticar y analizar el estado de la conexión entre cliente y servidor. El servidor responde a esta petición reflejando la misma información de la misma.

<b>OPTIONS</b>: Este método permite al cliente identificar los métodos HTTP permitidos para un recurso o por todo el servidor. El servidor agrega a la respuesta un encabezado Allow en donde especifica los métodos permitidos. Este método es comúnmente utilizado en las peticiones preflight para averiguar de antemano si el servidor aceptará la petición que se pretende hacer o no.

<b>CONNECT</b>: No es un método de uso común pero permite iniciar una comunicación de doble vía entre el cliente y el servidor, como un túnel TCP o un web socket.

<b>Códigos de estatus de respuesta a una petición HTTP</b>

Los códigos de estatus de una respuesta a una petición HTTP son códigos numéricos que expresan el resultado obtenido luego de atender la petición. De acuerdo con dicho resultado el servidor agrega un código a la respuesta para que el cliente pueda identificar si la petición fue atendida con éxito o si por el contrario hubo alguna falla o evento externo que no permitió que el servidor la atendiera correctamente.

Los códigos de estatus de HTTP se dividen en 5 categorías:

<ol>
<li>Informativos (100-199): Son códigos utilizados por el servidor para dar información. Comúnmente solo se utilizan los códigos 100 y 101, el primero para expresar que la petición se está atendiendo de forma correcta y se puede continuar y el último para indicar que el servidor está migrando a una nueva tecnología o protocolo.</li>

<li>Éxito (200-299): Con estos códigos el servidor informa que la petición se procesó correctamente. Los  códigos más comunes son 200 para indicar que la petición se atendió de forma correcta; 201 para indicar que un nuevo recurso fue correctamente creado, en respuesta a peticiones POST o PUT.</li>

<li>Redirección (300-399): Estos códigos indican que el recurso solicitado ha sido trasladado a una nueva URL. El servidor envía 301 en respuesta a peticiones GET y HEAD para indicar que el recurso ha sido movido de forma permanente a nueva URL, mientras que el código 308 indica lo mismo pero en respuesta a peticiones POST.</li>

<li>Error de parte del cliente (400-499): Los códigos dentro de esta categoría sirven para informar la razón por la cual el servidor no pudo procesar la petición, indicando que la falla está del lado del cliente ya sea por una petición mal formada o por que está solicitando un recurso que no existe en el servidor. Los códigos más frecuentes dentro de esta categoría: 400, el servidor no puede procesar la petición que está mal formada; 401, no su puede procesar la petición porque las credenciales de autenticación presentadas por el cliente no son válidas; 403, es similar a 401 pero en este caso puede que las credenciales sean válidas pero aún así el cliente no está autorizado a solicitar ese recurso; 404, es el más popular de esta categoría e indica que el cliente está solicitando un recurso que no existe temporal o permanentemente; 405, el servidor indica que el recurso que el cliente solicita no tiene habilitado el método HTTP con el que se hace la petición.</li>

<li>Error de parte del servidor (500-599): Estos códigos los utiliza el servidor para indicar que debido a una condición inesperada no se pudo procesar correctamente la petición. A esta categoría se le conoce como errores del lado del servidor, ya que a pesar de que la petición enviada por el cliente es correcta la petición no se puede responder satisfactoriamente. Los códigos más frecuentes dentro de esta categoría: 500, es un código comodín que el servidor utiliza para indicar que una condición inesperada no permitió el procesamiento de la petición; 501, este código se utiliza cuando el servidor no reconoce el método HTTP que utiliza la petición y que además es incapaz de procesar para cualquier recurso; 503, el servidor que responde que en ese momento no está listo para procesar la petición.</li>
</ol>

<img width="590" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/f7fe5f12-e919-4b57-a668-0bfa09d877a6">

### Anatomía de los mensajes del protocolo HTTP

Las peticiones y respuestas HTTP comparten una estructura similar compuesta por las siguientes partes:

<ol>
<li>Comienzo (start line): Es la primera línea de texto que describe la petición que se va a hacer, o bien, indica el estatus de la respuesta para saber si la petición se atendió correctamente o no. Este comienzo es una única línea de texto.</li>

<li>Encabezados (headers): Conjunto de encabezados, cada uno en su propia línea de texto individual, que describen detalles específicos de la petición o la información incluida dentro del cuerpo del mensaje.</li>

<li>Línea en blanco (empty/blank line): Es obligatoria y se agrega para separar la parte informativa (descriptiva) del mensaje del contenido incluido en el cuerpo del mensaje.</li>

<li>Cuerpo del mensaje (body): Es opcional ya que no todos los métodos HTTP requieren de un cuerpo con datos asociados a la petición o la respuesta. La presencia del cuerpo en el mensaje sugiere que además existen encabezados que detallan la longitud y el tipo de datos de la información allí incluida.</li>
</ol>

### Peticiones HTTP (Requests)

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

<img width="604" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/0a2fac49-9258-4e70-953a-2415487b02e0">

<img width="597" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/8137c3c1-77d5-4654-b888-e3960e508d2c">

### Respuestas HTTP (Responses)

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

<img width="602" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/d282a69a-60ad-4783-9d5f-61cd8a49bf3d">
