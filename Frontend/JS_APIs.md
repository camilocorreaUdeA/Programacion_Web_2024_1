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

