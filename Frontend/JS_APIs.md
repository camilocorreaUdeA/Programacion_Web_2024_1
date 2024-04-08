# Consumo de APIs REST con Javascript

## Conceptos útiles antes de comenzar

<img width="1081" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2024_1/assets/42076547/8a9c5ca3-1578-495a-b622-89c1300b4a64">

### Protocolo HTTP: Solicitudes, respuestas, métodos y códigos de estado

<ul>
  <li><b>Mensajes HTTP: </b>Solicitudes (requests) y respuestas (responses). Línea de comienzo, encabezados (headers), cuerpo de la solicitud/respuesta</li>
  <li><b>Métodos: </b>GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS, CONNECT, TRACE</li>
  <li><b>Códigos de estado (status codes): </b>200 (OK), 400 (errores del cliente), 500 (errores del servidor)</li>
</ul>

### API endpoints

<ul>
  <li><b>Uniform Resource Locator (URL): </b>Es una referencia o dirección a un recurso disponible en una red de computadores (Internet) y además indica la forma de conectarse al recurso (protocolo HTTP)</li>
  <li><b>Parámetros de la consulta (query parameters): </b>Elementos que se agregan a la URL y que sirven para filtrar los resultados disponibles en la respuesta a una solicitud HTTP</li>
</ul>

```code
// Ejemplo de un endpoint de una API muy popular
https://pokeapi.co/api/v2/pokemon/ditto
```

### JSON

Es un formato para presentar un recurso web como un objeto en notación del lenguaje Javascript.

```json
{
  "nombre": "Pepe Perez",
  "ocupacion": "Desarrollador Web",
  "edad": 28
}
```

## Consumo de APIs con <code>fetch</code>

[fetch()](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch)

<code>fetch</code> es una función de la API <code>Fetch</code> de Javascript que proporciona una interfaz programática para consumir de manera asincrónica los recursos expuestos por APIs REST. <code>fetch</code> retorna una promesa, objeto <code>Promise</code>, de que eventualmente la solicitud HTTP se va a completar y la respuesta va a estar disponible para procesarla.

### <code>fetch</code> con los handlers de las promesas <code>Promise</code>

```js
const url = 'https://pokeapi.co/api/v2/pokemon/ditto';
const options = {
    method: 'GET'
};

fetch(url, options)
    .then((response) => response.json())
    .then((data) => console.log(data))
    .catch((error) => console.log(error))
```

### <code>fetch</code> con <code>async</code> y <code>await</code>

```js
const url = 'https://pokeapi.co/api/v2/pokemon/ditto';

(async() => {
    try{
        const response = await fetch(url);
        const data = await response.json();
        console.log(data);
    }
    catch(error){
        console.log(error);
    }
})();
```

## Consumo de APIs con <code>axios</code>

[Axios](https://axios-http.com/docs/intro)

### <code>axios</code> con los handlers de las promesas <code>Promise</code>

```js
const axios = require('axios');

const url = 'https://pokeapi.co/api/v2/pokemon/ditto';

axios.get(url)
    .then((response) => console.log(response.data))
    .catch((error) => console.log(error))
```

### <code>axios</code> con <code>async</code> y <code>await</code>

```js
const axios = require('axios');

const url = 'https://pokeapi.co/api/v2/pokemon/ditto';

(async() => {
    try{
        const response = await axios.get(url);
        console.log(response.data);
    }
    catch(error){
        console.log(error);
    }
})();
```
## Incluyendo Javascript en la aplicación web

### El código de nuestros scripts

Anteriormente habíamos visto como incluir estilos CSS para dar estilo al documento HTML. Pudimos conocer 3 alternativas para incluir las reglas CSS: a través de un archivo independiente <code>styles.css</code>, utilizando el elemento <code>style</code> dentro del elemento <code>head</code> o bien mediante el atributo <code>style</code> de cada elemento por separado.

Con Javascript se puede implementar lógica que permite desarrollar una aplicación web dinámica que está en capacidad de interactuar con el usuario y de hacer solicitudes a APIs de REST para consultar recursos alojados en la web.

Para incluir código de Javascript en una aplicación web existen dos alternativas basadas en el elemento <code>script</code> que proporciona HTML:

1. Código en un archivo separado de extensión .js: Desarrollar toda la lógica con Javascript en uno o varios archivos independientes con extensión .js y luego incluir los archivos en el documento HTML con el elemento <code>script</code> utilizando el atributo <code>src</code> para dar la ruta a los archivos y los atributos booleanos <code>async</code> o <code>defer</code> para que el código solo ejecute una vez todo el documento HTML ha sido cargado por completo. El elemento <code>script</code> se puede poner como contenido de los elementos <code>head</code> o <code>body</code> de manera indistinta. Por convención se acostumbra en ponerlo en el elemento <code>head</code> luego de la inclusión de los archivos de estilo CSS.

```html
<!doctype html>
<html>
  <head>
    <script src="scripts/script.js" async></script>
  </head>
  <body>
  </body>
</html>
```

2. Código directamente en el documento HTML: En este caso se incluye el código Javascript como contenido del elemento <code>script</code>.

```html
<!doctype html>
<html>
  <head>
  </head>
  <body>
    <script>
       const code = "<h1>Web Apps Development</h1>";
       document.body.innerHTML = code;
    </script>
  </body>
</html>
```

### Paquetes de terceros

Los paquetes de terceros (que no vienen incluidos de forma nativa en Javascript), como en el caso de <code>axios</code> se deben incluir en el documento HTML de forma explícita utilizando la primera alternativa de las dos que vimos anteriormente. En la documentación oficial de cada paquete externo se suele indicar la forma de incluirlo. Por ejemplo en el caso de axios, la documentación proporciona estas alternativas que utilizan CDN (Content Delivery/Distribution Network):

CDN: jsDelivr

```js
<script src="https://cdn.jsdelivr.net/npm/axios/dist/axios.min.js"></script>
```

CDN: unpkg

```js
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
```
Y luego se debe importar el paquete externo en el código Javascript donde se van a utilizar sus funciones:

Importando usando el viejo estilo de Javascript (con <code>require</code>)

```js
const axios = require('axios');
```

Importando con el nuevo estilo (con <code>import</code>)

```js
import axios from 'axios';
```
   


