# DOM y Javascript

## Document Object Model (DOM)

### ¿Qué es el DOM?

El <i>Document Object Model</i> o modelado del documento HTML como un objeto es una representación en forma de árbol de todos los elementos que conforman el contenido de un documento HTML. El árbol es un esquemático de las relaciones que existen entre los elementos HTML que componen el documento. Explicado de una manera más sencilla, el árbol es el gráfico de un conjunto de nodos conectados de forma jerárquica, es decir nodos padres conectados a sus nodos hijos, donde los nodos son los elementos HTML del documento y las conexiones indican relación de descendencia/ascendencia entre dos nodos interconectados. Luego, cada rama del árbol indica los niveles de anidamiento que existen en un elemento HTML determinado.

```html
<!DOCTYPE html>
<html>
 <head>
  <title>Mi web</title>
 </head>
 <body>
  <header>
   <h1>Bienvenido a mi sitio web!</h1>
  </header>
  <main>
   <section>
    <article>
     <h2>Acerca de mí...</h2>
     <img src="img/yo.jpg" alt="Yo programando xD" \>
     <p>Soy un desarrollador web con 5 años de experiencia, mis tecnologías favoritas son...</p>
    </article>
   </section>
  </main>
 </body>
</html>
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/3880aabd-561a-425a-b553-17587e42ca6a)

El árbol es solo una representación gráfica de lo que es el DOM, lo realmente importante es que el DOM modela el documento HTML como un objeto (y también a todos los nodos pertenecientes al árbol del DOM), ¿se les hace común ese término?, si pensaste en el concepto de objeto de la programación orientada a objetos estás en la ruta indicada. Al modelar el documento HTML como un objeto el DOM nos proporciona propiedades y métodos que podemos aprovechar para acceder a los elementos HTML del documento y de esa forma podamos manipular, insertar, actualizar y remover elementos HTML del documento con ayuda de un lenguaje de programación, de preferencia un lenguaje interpretado. La opción más óptima en este caso es el lenguaje de programación Javascript ya que el navegador web está en capacidad de interpretar y ejecutar código escrito en dicho lenguaje.

![tomado de https://www.w3schools.com/js/js_htmldom.asp](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/f25880b0-15ec-42c7-9726-246e45fbfb97)


### Propiedades y métodos del DOM

Para acceder al nodo principal de un documento HTML utilizamos la propiedad <code>document.documentElement</code>, este nodo contiene a todos los elementos del documento HTML, es decir, esta propiedad nos devuelve un objeto del elemento <code>html</code> del documento. Recuerde que todos los elementos del documento descienden de este elemento y por encima de el no hay ningún otro elemento, por eso mismo es el nodo principal del documento. Abra una consola de Javascript en su navegador y ejecute el siguiente comando para que verifique esta propiedad.

```js
console.log(document.documentElement)
```
Todos los nodos del árbol del DOM tienen un conjunto de propiedades y métodos (recuerda que estos nodos son modelados como objetos) que nos van a permitir manipular los elementos de la página web con la ayuda de Javascript.

A continuación vamos a proporcionar un listado de las propiedades del DOM de uso más común, para ver un listado más completo puede visitar este [enlace](https://linuxhint.com/html-dom-methods-and-properties/)

Algunas propiedades de los nodos del DOM:
<ul>
  <li><b>innerHTML: </b>Esta propiedad permite acceder y manipular el contenido de un elemento HTML.</li>
  <li><b>nodeName: </b>Devuelve el nombre del nodo.</li>
  <li><b>nodeValue: </b>Devuelve o permite acceder al valor del nodo.</li>
  <li><b>nodeType: </b>Devuelve el tipo de nodo.</li>
  <li><b>parentNode: </b>Devuelve el nodo padre del elemento.</li>
  <li><b>parentElement: </b>Funciona igual que el anterior pero devuelve nulo cuando el nodo padre no es un elemento HTML.</li>
  <li><b>style: </b>Retorna un objeto que permite acceder a las reglas de estilo del elemento HTML (declaradas en línea con el atributo style)</li>
  <li><b>tagName: </b>Devuelve el nombre de la etiqueta del elemento HTML.</li>
  <li><b>textContent: </b>Devuelve y premite cambiar el texto del contenido de un elemento HTML. Puede simplificarse a <i>text</i></li>
  <li><b>innerText: </b>Funciona igual al anterior.</li>
  <li><b>firstChild: </b>Devuelve un objeto que permite acceder a las propiedades del primer nodo descendiente del elemento.</li>
  <li><b>lastChild: </b>Igual que el anterior pero en el caso del último descendiente</li>
  <li><b>className: </b>Permite conocer el valor del atributo <i>class</i> del elemento</li>
  <li><b>childNodes: </b>Devuelve una lista con todos los nodos descendientes del elemento</li>
  <li><b>attributes: </b>Devuelve una lista con todos los atributos del elemento.</li>
  <li><b>childElementCount: </b>Devuelve el número total de elementos que son descendientes directos del elemento.</li>
  <li><b>id: </b>Devuelve el valor del atributo <i>id</i> del elemento.</li>
</ul>

¿Cómo seleccionamos los elementos (nodos) del documento HTML?

¿Recuerdas los selectores que utilizabamos para aplicar las reglas de estilo CSS a un elemento en particular del documento? Pues, esos mismos selectores nos van a ser de utilidad para seleccionar un elemento y procesarlo con Javascript. Para ese propósito vamos a utilizar el método <code>querySelector</code> del nodo <i>document</i>

[Ejemplo:](https://codepen.io/camilocorreaUdeA/pen/KKbrbge?editors=1111)
```html
<html>
 <head></head>
 <body>
  <ul class="lista">
   <li>Frijoles</li>
   <li>Arepa</li>
   <li>Chicharron</li>
  </ul>
 </body>
</hmtl>
```
```js
/* seleccionando el primer elemento li de la lista ul */
const primerElementoLista = document.querySelector(".lista :first-child");
console.log(primerElementoLista.textContent); /* Frijoles */
```
Con el método <code>querySelectorAll</code> puedes seleccionar todos los elementos que compartan el mismo selector CSS:
[Ejemplo:](https://codepen.io/camilocorreaUdeA/pen/oNJQJZj?editors=1111)

```html
<html>
 <head></head>
 <body>
  <ul class="lista">
   <li class="item">Frijoles</li>
   <li class="item">Arepa</li>
   <li class="item">Chicharron</li>
  </ul>
 </body>
</hmtl>
```
```js
/* seleccionando todos los elementos li de la lista ul */
const elementosLista = document.querySelectorAll(".item");
const arepa = elementosLista[1];
console.log(arepa.textContent); /* Arepa */
```
Para seleccionar un elemento por el valor del atributo <i>id</i> se utiliza el método <code>getElementById</code>
[Ejemplo](https://codepen.io/camilocorreaUdeA/pen/RwEqEVo?editors=1111)

```html
<html>
 <head></head>
 <body>
  <p id="par">Texto del parrafo</p>
 </body>
</hmtl>
```
```js
/* seleccionando un elemento por id */
const parrafo = document.getElementById("par");
console.log(parrafo.textContent); /* Texto del parrafo */
```

Para seleccionar elementos por el valor del atributo <i>class</i> se utiliza el método <code>getElementsByClassName</code>
[Ejemplo](https://codepen.io/camilocorreaUdeA/pen/JjwewJo?editors=1111)

```html
<html>
 <head></head>
 <body>
  <p class="par">Texto del parrafo 1</p>
  <p class="par">Texto del parrafo 2</p>
  <p class="par">Texto del parrafo 3</p>
 </body>
</hmtl>
```
```js
/* seleccionando todos los elementos de la misma class */
const parrafos = document.getElementsByClassName("par");
console.log(parrafos[1].textContent); /* Texto del parrafo 2*/
```
O bien se pueden seleccionar elementos por la etiqueta del elemento con el método <code>getElementsByTagName</code>
[Ejemplo](https://codepen.io/camilocorreaUdeA/pen/QWzJzMx?editors=1111)

```html
<html>
 <head></head>
 <body>
  <p class="par">Texto del parrafo 1</p>
  <p class="par">Texto del parrafo 2</p>
  <p class="par">Texto del parrafo 3</p>
 </body>
</hmtl>
```
```js
/* seleccionando todos los elementos con la misma etiqueta */
const parrafos = document.getElementsByTagName("p");
console.log(parrafos[1].textContent); /* Texto del parrafo 2*/
```
Combinando los métodos selectores y las propiedades de los nodos del DOM podemos hacer modificaciones al contenido (y al estilo) de los elementos HTML del documento.
[Ejemplo](https://codepen.io/camilocorreaUdeA/pen/yLGQGPG?editors=1111)

```html
<html>
 <head></head>
 <body>
  <div></div>
 </body>
</hmtl>
```
```js
/* seleccionando todos los elementos con la etiqueta div y modificando su contenido HTML*/
const newContent = '<p><a href="https://learnjavascript.online/">¡Aprende Javascript en línea!</a></p>';
document.getElementsByTagName("div")[0].innerHTML = newContent;
```
[Ejemplo](https://codepen.io/camilocorreaUdeA/pen/bGOQOvM?editors=1111)

```html
<html>
 <head></head>
 <body>
  <div>
   <strong>Cursos recomendados para aprender Javascript en 2023:</strong>
   <ul id="lista-cursos-js">
    <li>Curso 1</li>
    <li>Curso 2</li>
    <li>Curso 3</li>
   </ul>
  </div>
 </body>
</hmtl>
```
```js
/* modificando los elementos de un listado */
const cursos = [
  {
    nombre:"Learn Javascript Online",
    enlace:"https://learnjavascript.online/"
  },
  {
    nombre:"Introduction to Javascript",
    enlace:"https://learnjavascript.online/"
  },
  {
    nombre:"Javascript.Info",
    enlace:"https://javascript.info/"
  }
];
let itemsLista = document.querySelector("#lista-cursos-js").children;
for(i in itemsLista){
  itemsLista[i].textContent = `${cursos[i].nombre}: ${cursos[i].enlace}`;
}
```
¿Cómo modificar los estilos? Podemos modificarlos de varias formas pero aquí vamos a revisar dos de ellas. La primera forma es a través de la propiedad <i>style</i> de los nodos del DOM, al modificar esta propiedad estamos modificando directamente los estilos en línea del elemento. La segunda forma, que utiliza las reglas de estilo CSS definidas es un archivo externo (styles.css), requiere que las modificaciones se hagan en los atributos del elemento con el objetivo de lograr una coincidencia exacta con el selector de la regla CSS que se ha preparado con anticipación.

```js
/* formas de agregar estilos a un elemento */
const nodo = document.querySelector("div");
/* accediendo a las propiedades directamente a traves de la propiedad style */
div.style.backgroundColor = "lightblue";
/* accediendo a las propiedades a traves de la propiedad cssText */
div.style.cssText = "background-color: lightblue; border: 2px solid black;";
/* asignando el atributo style y su valor con el método setAttribute */
div.setAttribute("style", "background-color: lightblue; border: 2px solid black;")
```
[Ejemplo](https://codepen.io/camilocorreaUdeA/pen/mdaQgwR?editors=1111)
```html
<html>
 <head></head>
 <body>
  <div>
   <p>Desarrollo de aplicaciones web</p>
  </div>
 </body>
</hmtl>
```
```js
/* modificando los estilos */
const div = document.querySelector("div");
div.setAttribute('style', 'background-color: rgba(113, 215, 249, 0.45); border: 2px solid black;');
const p = document.querySelector("div p");
p.style.color = "red";
p.style.fontSize = "30px";
```
Podemos lograr el mismo objetivo haciendo coincidir el selector del elemento con un selector predefinido en las reglas CSS. En este ejemplo en particular vamos utilizar la propiedad <i>classList</i> que permite agregar o modificar los valores del atributo <i>class</i> de un elmento, y si ese atributo no existe lo crea. Luego vamos a llamar al método <i>add</i> para agregar al atributo <i>class</i> el valor que determinará la coincidencia con el selector de las reglas CSS. Se puede lograr el mismo efecto con el método <code>setAttribute</code>, con este último podemos asignar un valor al atributo <i>id</i> y lograr la coincidencia con un selector por id. [Ejemplo](https://codepen.io/camilocorreaUdeA/pen/NWeEmMz?editors=1111) 
```html
<html>
 <head></head>
 <body>
  <div>
   <p>Desarrollo de aplicaciones web</p>
  </div>
 </body>
</hmtl>
```
```css
.my-div{
  background-color: rgba(113, 215, 249, 0.45); 
  border: 2px solid black;  
}

.my-div p{
  color: red;
  font-size: 30px;
}
```
```js
/* modificando los estilos */
const div = document.querySelector("div");
div.classList.add("my-div");
```

Para agregar (insertar) un nuevo elemento al documento puede seguir los siguientes pasos:
<ol>
<li>Crear el nuevo elemento con el método <code>document.createElement</code>, este recibe como parámetro la etiqueta del elemento que desea crear.</li>
<li>Agregue el contenido al nuevo elemento. Puede utilizar la propiedad <code>innerHTML</code> o si el contenido va a ser solo texto puede usar el metodo <code>createTextNode pasandole como parámetro el texto del contenido.</code></li>
<li>Agregue el nuevo elemento al documento. Los nuevos elementos se insertan o se agrean a un elmento existente, que será el nodo padre del nuevo elemento, para insertar el nuevo elemento al final  de los otros descendientes del nodo padre puede utilizar el método <code>appendChild</code> del elemento padre pasando en el parámetro el nuevo elemento descendiente. O si quiere ubicar el nuevo elemento en una ubicación predeterminada (tomando como punto de referencia otro descendiente existente) puede utilizar el método <code>insertBefore</code> del nodo padre indicando en los parámetros el nuevo elemento y el elemento antes de cuya posición quiere insertar el nuevo elmento.</li>
</ol>
[Ejemplo](https://codepen.io/camilocorreaUdeA/pen/ZEVVWXb?editors=1111)

```html
<html>
 <head></head>
 <body>
  <div>
   <strong>Lista de tareas</strong>
    <ul id="to-do-list">
    </ul>      
  </div>
 </body>
</hmtl>
```
```js
/* agregando nuevos elementos */
const tareas = [
    "Repasar HTML",
    "Consultar acerca de CSS",
    "Aprender Javascript",
    "Practicar la consulta de APIs",
    "Construir mi primer aplicación web"    
  ];
  for(tarea of tareas) {
    const lista = document.querySelector("#to-do-list");
    const nuevaTarea = document.createElement("li");
    const textoTarea = document.createTextNode(tarea);
    nuevaTarea.appendChild(textoTarea);
    lista.appendChild(nuevaTarea);   
  }
```
[Ejemplo](https://codepen.io/camilocorreaUdeA/pen/mdaaPYa)

```html
<html>
 <head></head>
 <body>
  <div>
   <p>Yo soy el primer párrafo</p>     
  </div>
 </body>
</hmtl>
```
```js
/* agregando nuevos elementos */
const div = document.querySelector("div");
const primerP = document.querySelector("div p");
const newP = document.createElement("p");
newP.textContent = "Eras...yo ya soy el primer párrafo";
div.insertBefore(newP, primerP);
```
Y para remover elementos existe el método <code>removeChild</code>. Así:

```js
nodoPadre.removeChild(nodoDescendiente);
```

### Captura de eventos

Un evento es una acción que ocurre o se ejecuta en una página web, por jemplo cuando se da clic a un botón o a un enlace. Con Javascript pdemos hacer que la página web escuche o este pendiente de esos eventos para luego reaccionar a su ocurrencia. Los eventos más comunes son: el clic del mouse, la presión de la tecla enter, pasar el puntero del mouse por encima de algún elemento, la selección de una opción en un menú desplegable, marcar una caja de selección, etc. Para ver una lista completa de los eventos disponibles puede visitar este [enlace](https://developer.mozilla.org/en-US/docs/Web/Events).

Captura de eventos como atributos de los elementos HTML.

Podemos agregar atributos a los elementos HTML que permiten que puedan escuchar y atender un evento. Por ejemplo, el atributo <code>onclick</code> permite especificar una función que se debe ejecutar una vez dan clic en el elemento. [Ejemplo](https://codepen.io/camilocorreaUdeA/pen/ZEVVOPx)

```html
<html>
 <head></head>
 <body>
  <p>Desarrollo de aplicaciones web!</p>    
  <button onclick="cambiar()">Dale clic!</button>
 </body>
</hmtl>
```
```css
.clase {
  color: green;
  font-size: 40px;
  font-weight: bold;
}
```
```js
/* función que atiende el evento */
function cambiar() {
  const p = document.querySelector("p");
  p.setAttribute("class", "clase");
}
```
Agregando <i>listeners</i> en Javascript

Si prefiere separar la lógica del contenido puede trasladar la escucha de los eventos al script de Javascript. La primera forma de hacerlo sería seleccionando el elemento y asignando la ejecución de una función a la propiedad asociada al evento que se desea atender. En el siguiente ejemplo reproduciremos el anterior pero trasladando la esucha del evento al código de Javascript. [Ejemplo](https://codepen.io/camilocorreaUdeA/pen/XWooKvx)

```html
<html>
 <head></head>
 <body>
  <p>Desarrollo de aplicaciones web!</p>    
  <button>Dale clic!</button>
 </body>
</hmtl>
```
```css
.clase {
  color: green;
  font-size: 40px;
  font-weight: bold;
}
```
```js
const boton = document.querySelector("button");
boton.onclick = () => cambiar(); /* escuchando el evento */

function cambiar() {
  const p = document.querySelector("p");
  p.setAttribute("class", "clase");
}
```

En la segunda forma vamos a utilizar el método <code>addEventListener</code> con el que vamos a indicar el evento a escuchar y la función con la que se debe atender. Vamos a reproducir de nuevo el primer ejemplo pero aplicando esta segunda forma. [Ejemplo](https://codepen.io/camilocorreaUdeA/pen/gOZZwpw)

```html
<html>
 <head></head>
 <body>
  <p>Desarrollo de aplicaciones web!</p>    
  <button>Dale clic!</button>
 </body>
</hmtl>
```
```css
.clase {
  color: green;
  font-size: 40px;
  font-weight: bold;
}
```
```js
/* atendiendo el evento con un listener */
const boton = document.querySelector("button");
boton.addEventListener(
  "click",
  () => {
    const p = document.querySelector("p");
    p.setAttribute("class", "clase");    
  }
);
```
