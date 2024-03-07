# CSS

## Maquetación de los elementos en el documento HTML (CSS Layout)

### Flujo normal del documento HTML

Por defecto y a menos de que hayas introducido manualmente algún cambio a la maquetación del documento HTML usando CSS los elementos de un sitio web se maquetan o posicionan utilizando el flujo normal o flujo por defecto del navegador. Este flujo normal asegura que el contenido de la página web sea <i>legible</i> para cualquier navegador web o dispositivo de navegación asistida, como un lector de pantalla.

La pregunta ahora es ¿Cómo posiciona los elementos en pantalla el flujo normal? Y la respuesta está en un concepto que ya vimos con anterioridad que es el modelo de cajas. Es decir, los elementos individuales de un documento HTML se posicionan de tal forma que se tenga en estricta consideración las propiedades del modelo de cajas (padding, border, margin) inherentes al elemento HTML en cuestión.

Ahora, primero que todo debemos hacer una distinción entre los elementos HTML separandolos en dos categorías a saber:

<ul>
  <li>Elementos de bloque (<i>Block-level elements</i>)</li>
  <li>Elementos de línea (<i>Inline-level elements</i>)</li>
</ul>

Los elementos de bloque se ubican en pantalla en una línea exclusiva ocupando por completo el ancho de asignado para el contenido del elemento padre (ejemplo: si el elemento padre es el <code>body</code>, entonces será todo el ancho de pantalla), y se apilan uno por debajo del otro de manera sucesiva y de acuerdo al orden de aparición en el documento HTML.

```html
<body>
  <h1>Aplicaciones Web</h1>
  <p>She was infatuated with color. She didn't have a favorite color per se, but she did have a fondness for teals and sea greens. 
  You could see it in the clothes she wore that color was an important part of her overall style. She took great pride that color 
  flowed from her and that color was always all around her. That is why, she explained to her date sitting across the table, that 
  she could never have a serious relationship with him due to the fact that he was colorblind.</p>
</body>
```
```css
h1{
  font-size:70px;
  background-color:lightgreen;
  border:2px solid black;
}
p{
  font-size:25px;
  background-color:lightblue;
  border:2px solid black;
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/66cd900c-2cf4-43e3-b97e-789bf5ae0433)

En el ejemplo vemos 2 elementos de bloque: <code>h1</code> y <code>p</code>, podemos observar lo que se había comentado anteriormente, que ocupan todo el ancho del elemento contenedor (en este caso un elemento <code>body</code>), están apilados uno sobre otro de acuerdo con el orden de aparición en el documento HTML y su dimensión vertical se ajusta al tamaño del contenido (observe que el párrafo cuenta con una mayor dimensión vertical que el encabezado). Además se evidencia una separación entre los dos elementos (para esto es que hemos resaltado el border), lo cual se esperaba ya que por defecto los elemento de bloque tienen algún valor en la propiedad margin, usualmente en margin-top y margin-bottom.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/647c2f8e-9ed9-4456-9f5d-fe0c3d7a27f6)
<p>Se puede observar en la imagen que el margin por defecto asignado al elemento <code>h1</code> es de 46.9px y solo en las direcciones arriba y abajo.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/318f88d3-58b3-489d-bfa1-ada25707ec65)
<p>Mientras que el margin por defecto del elemento <code>p</code> es de 25px y también solo en las direcciones arriba y abajo.</p>

Otro detalle interesante que se destaca es como el margin del elemento <code>h1</code> llega hasta el límite del borde del elemento <code>p</code>, es decir la separación real entre ambos elementos son 46.9px y no una suma de los margin de los dos elementos como se podría llegar a pensar. A ese comportamiento en la separación de dos elementos de bloque contiguos se le conoce como <i>margin collapse</i> o colapso de márgenes, y es simplemente que el margin con mayor dimensión absorbe al de menor dimensión para convertirse en la separación efectiva entre los dos elementos de bloque.

En cuanto a los elementos de línea, estos no son apilados uno tras otro en una nueva línea de la pantalla para cada uno, y en cambio son ubicados uno tras otro de manera consecutiva de acuerdo con el orden de aparición dentro del documento HTML en la misma línea hasta que agotan todo el ancho del elemento contenedor y eventualmente pasan a ocupar la línea siguiente una vez sobrepasan dicho ancho. Por defecto las dimensiones de un elemento de línea coinciden con las dimensiones del contenido del mismo, aún así es posible modificarlas a través de CSS (siempre y cuando hacer eso tenga sentido).

```html
<body>
  <div>
  <a href="#">Hola</a>
  <a href="#">Mundo</a>
  <a href="#">Aplicaciones</a>
  <a href="#">Web</a>
  <a href="#">Telecomunicaciones</a>
  <strong>she explained to her date sitting across the table</strong>
  </div>
</body>
```
```css
strong{
  font-size:25px;
  border:2px solid black;
  background-color:grey;
}

a{
  font-size:25px;
  border:2px solid black;
  background-color:lightpink;
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/abed4d6c-0a19-4c23-a3ed-1161bee594ae)

En la imagen a continuación se observa como los elementos de línea van llenando paulatinamente el ancho del elemento contenedor.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/31a0f06f-f0b1-4c00-90b3-d08efb75e251)

En las siguientes imagenes se puede observar que no tienen margenes asignadas por defecto (ni padding ni margin) y que el tamaño del elemento se ajusta al del contenido.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/144b83ca-62c8-4147-9bb3-15fe2fb6527b)

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/120407cd-4a31-4604-bd0e-e0a03650e2c3)

En este [enlace](https://www.w3schools.com/html/html_blocks.asp) puede ver listas separadas de los elementos de bloque y de línea de HTML.

### Contenedores <code>div</code> y <code>span</code>

<code>div</code>

<p>El elemento <code>div</code> es un elemento de bloque que actúa como contenedor genérico para otros elementos. A diferencia de los elementos semánticos de HTML <code>div</code> no expresa el significado de su contenido y en parte se debe a que su función es simplemente la de ser un elemento auxiliar que permita dividir la página web en distintos bloques que pueden ser estilizados con CSS individualmente.</p>

```html
<body>
  <div class="bloque-1"></div>
  <div class="bloque-2"></div>
  <div class="bloque-3"></div>
</body>
```
```css
div{
  height:200px;
  border:2px solid black;
}

.bloque-1{
  background-color:#4061c4;  
}

.bloque-2{
  background-color:rgba(251, 203, 80, 0.54);
}

.bloque-3{
  background-color:rgba(180, 45, 4, 0.52);
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/c3765efe-0d88-423a-943a-071785578a6f)

<code>span</code>

<p>El elemento <code>span</code> también es un contenedor genérico para agrupar otros elementos y estilizarlos. Pero <code>span</code></p> es un elemento de línea, por tanto es más útil para agrupar textos u otros elementos de línea como los hipervínculos (elemento <code>a</code>).</p>

```html
<body>
  <p>She was infatuated with color. She didn't have a favorite color per se, but she did have a fondness for teals and sea greens. You could see it in the clothes she wore that color was an important part of her overall style. She took great pride that <span id="text">color flowed from her and that color was always all around her.</span> That is why, she explained to her date sitting across the table, that she could never have a serious relationship with him due to the fact that he was colorblind.</p>
</body>
```
```css
#text:hover{
  background-color:lightpink;
  border:2px solid black;
  font-size:35px;
  font-style:italic;
}
```

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/d100349b-cb99-4afa-a645-0a3cf2a69c9a)

### Propiedad <code>display</code>

La propiedad <code>display</code> sirve para determinar si un elemento es de bloque o de línea. Modificando esta propiedad se puede lograr que un elemento que por naturaleza es de bloque se convierta en un elemento de línea y vicerversa, y además nos va a ser también de utilidad para construir elementos de cuadrícula (<i>grid</i>) y elementos de caja flexible (<i>flex</i>).

Esta propiedad también puede tener el valor <i>inline-block</i> que permite que los elementos se comporten como elementos de línea a los que se les puede asignar dimensiones de ancho y alto (<i>width</i> y <i>height</i>). Recuerde que los elementos de línea ignoran esas propiedades porque ajustan sus dimensiones a las del contenido.

Posibles valores para la propiedad <code>display</code>:
<ul>
  <li><code>none</code>: Oculta el elemento de modo que no es visible en el navegador web.</li>
  <li><code>block</code>: Elementos de bloque que ocupan todo el ancho del elemento contenedor y se apilan uno tras otro separados por el margen externo</li>
  <li><code>inline</code>: Elementos de línea que se ubican de manera contigua en una misma línea y cuyas dimensiones son las mismas de su propio contenido.</li>
  <li><code>inline-block</code>: Elementos que se comportan como si fueran de línea pero con la posibilidad de manipular sus dimensiones.</li>
  <li><code>grid</code>: Elementos de bloque que permiten maquetar los elementos en su contenido de acuerdo al modelo grid de CSS</li>
  <li><code>flex</code>: Elementos de bloque que permiten maquetar los elementos en su contenido de acuerdo al modelo flexbox de CSS</li>
</ul>

### Maquetación con CSS Grid

CSS grid es un sistema de maquetación en dos dimensiones que permite organizar el contenido de un elemento contenedor en una cuadrícula formada por filas y columnas. Este sistema facilita la creación de diseños complejos ya que es posible realizar anidamientos de contenedores grid al interior de otros contenedores grid.

<ul>
  <li>Para definir un elemento <code>grid</code> se debe utilizar la propiedad <code>display</code> con el valor <code>grid</code>.</li>
  <li>Para especificar el número de columnas y el ancho de las mismas se utiliza la propiedad <code>grid-template-columns</code></li>
  <li>Para especificar el número de filas de la cuadrícula y la altura de las mismas se utiliza la propiedad <code>grid-template-rows</code></li>
  <li>Para determinar el espacio entre filas se utiliza la propiedad <code>grid-row-gap</code>. Para las columnas la propiedad <code>grid-column-gap</code>, o bien el mismo valor para ambos con la propiedad <code>grid-gap</code></li>
</ul>

Por ejemplo una cuadrícula de 6x6 (de 200px por 200px cada celda) con espacio entre columnas y filas de 15px:

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/dd491d85-33b3-4f75-a8e8-a67a0131c32e)

```html
<body>
  <div class="contenedor">
    <div>1</div>
    <div>2</div>
    <div>3</div>
    <div>4</div>
    <div>5</div>
    <div>6</div>
    <div>7</div>
    <div>8</div>
    <div>9</div>
  </div>
</body>
```
```css
.contenedor{
  display:grid;
  grid-template-columns: 200px 200px 200px;
  grid-template-rows: 200px 200px 200px;
  grid-gap:15px;
}

.contenedor > div{
 border:2px solid black;
 text-align:center;
 background-color:rgba(80, 146, 251, 0.54);
 padding:85px;
 font-size:25px;
 font-weight:bold;
}
```
Hay ocasiones en las que se prefiere indicar las dimensiones de las cuadrículas en fracciones del espacio disponible dentro del elemento contenedor, es en esos casos que es de utilidad la unidad <i>fr</i>, ya que permite expresar la dimensión de un elemento como tantas partes de fracción del espacio disponible en el contenedor.

Entonces el navegador calcula cuántas unidades <i>fr</i> y distibuye el espacio disponible en esa suma de unidades. Y luego a cada columna o fila se le asigna tantas unidades como estén expresadas en las propiedades <code>grid-template-columns</code> y <code>grid-template-rows</code> respectivamente.

En el ejemplo siguiente, los pixeles restantes luego de separar los 200px para la primera columna se reparten en un total de 3 unidades <i>fr</i>, correspondiendole 1 unidad a la segunda columna y 2 a la tercera.

```html
<body>
  <div class="contenedor">
    <div>200px</div>
    <div>1fr</div>
    <div>2fr</div>    
  </div>
</body>
```
```css
.contenedor{
  display:grid;
  grid-template-columns: 200px 1fr 2fr;
  grid-gap:3px;
}

.contenedor > div{
 border:2px solid black;
 text-align:center;
 background-color:rgba(80, 146, 251, 0.54);
 padding:85px;
 font-size:25px;
 font-weight:bold;
}
```

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/f6bc01c4-14c2-4020-b76a-2eecbe8c13b8)

Para definir diseños más dinámicos y complejos las propiedades <code>grid-column</code> y <code>grid-row</code> permiten  definir sobre cuáles columnas o filas se debe desplegar un elemento, de esta forma se puede tener elementos que ocupan varias columnas y filas al mismo tiempo. De esa forma es posible romper con la cuadrícula tradicional perimitiendo diseños más dinámicos donde hay cuadrículas que cuentan con dimensiones mayores en comparación las otras.

```html
<body>
  <main>
    <header>header</header>
    <nav>nav</nav>
    <aside>aside</aside>
    <article>article</article>
    <footer>footer</footer>
  </main> 
</body>
```
```css
main{
  display:grid;
  grid-template-columns: 1fr 2fr;
  grid-template-rows: 2fr 1fr 4fr 1fr;
  width:800px;
  height:600px;
}

header{
 border:2px solid black;
 text-align:center;
 background-color:rgba(80, 146, 251, 0.54);
 font-size:25px;
 font-weight:bold;
 grid-column:1/3;
 grid-row:1;
}

nav{
 border:2px solid black;
 text-align:center;
 background-color:rgba(137, 251, 80, 0.54);
 font-size:25px;
 font-weight:bold;
 grid-column:1/3;
 grid-row:2;
}

aside{
 border:2px solid black;
 text-align:center;
 background-color:rgba(251, 120, 80, 0.48);
 font-size:25px;
 font-weight:bold;
 grid-column:1;
 grid-row:3;
}

article{
 border:2px solid black;
 text-align:center;
 background-color:rgba(251, 80, 80, 0.67);
 font-size:25px;
 font-weight:bold;
 grid-column:2;
 grid-row:3;
}

footer{
 border:2px solid black;
 text-align:center;
 background-color:rgba(251, 249, 80, 0.48);
 font-size:25px;
 font-weight:bold;
 grid-column:1/3;
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/f38263a9-f0ee-4d67-bf64-de59ede61451)

Otra manera de posicionar los elementos con <b><i>grid</i></b> es utilizando la propiedad <code>grid-template-areas</code>. El mecanismo con está propiedad es el de dar nombres a las partes en las que se divide el documento HTMLy a cada una de esas partes asignar los elementos que va a contener. Entonces a <code>grid-template-areas</code> se le asigna como valor un esquema de las distintas áreas que van a componer la página web, se usan los nombres asignados a las áreas en cada columna o fila que vaya a componer dicha área. Es importante mencionar que solo son posibles áreas cuadradas o rectangulares.

Luego en cada elemento, con la propiedad <code>grid-area</code> se indica el área donde se debe posicionar ese elemento. Esto funciona similar a como se vio anteriormente con las propiedades <code>grid-column</code> y <code>grid-row</code>.

```html
<body>
  <main>
    <header>header</header>
    <nav>nav</nav>
    <aside>aside</aside>
    <article>article</article>
    <footer>footer</footer>
  </main> 
</body>
```
```css
main{
  display:grid;
  grid-template-areas: 
    "header header"
    "nav nav"
    "aside article"
    "footer footer";
  grid-template-columns: 1fr 2fr; 
  width:800px;
  height:600px;
}

header{
 border:2px solid black;
 text-align:center;
 background-color:rgba(80, 146, 251, 0.54);
 font-size:25px;
 font-weight:bold;
 grid-area:header;
}

nav{
 border:2px solid black;
 text-align:center;
 background-color:rgba(137, 251, 80, 0.54);
 font-size:25px;
 font-weight:bold;
 grid-area:nav;
}

aside{
 border:2px solid black;
 text-align:center;
 background-color:rgba(251, 120, 80, 0.48);
 font-size:25px;
 font-weight:bold;
 grid-area:aside;
}

article{
 border:2px solid black;
 text-align:center;
 background-color:rgba(251, 80, 80, 0.67);
 font-size:25px;
 font-weight:bold;
 grid-area:article;
}

footer{
 border:2px solid black;
 text-align:center;
 background-color:rgba(251, 249, 80, 0.48);
 font-size:25px;
 font-weight:bold;
 grid-area:footer;
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/f3d067cd-c27d-4930-ac1a-ac2698a0b496)

### Maquetación con CSS Flexbox

Flexbox es un mecanismo para posicionar elementos de un documento HTML en filas y columnas. Flexox utiliza dos tipos de "cajas": <i>flex containers</i> y <i>flex items</i>. Los flex containers sirven para agrupar varios flex items y definir la forma de posicionarlos. Con CSS se especifica de forma explicita cuando un elemento HTML es un flex container pero en el caso de un flex item este solo será tal si está contenido en un flex container, de forma que la única tarea de los flex items es indicar al flex container cuántos elementos este debe tener en cuenta para posicionar o maquetar.

Un elemento HTML puede ser un flex container y un flex item al mismo tiempo, siempre y cuando esté contenido dentro de otro flex container y al tiempo ser un flex container para otros elementos (sus propios flex items en ese caso).

Para definir un elemento como flex container simplemente se utiliza la propiedad <code>display</code> dandole el valor <code>flex</code>, luego se debe definirla disposición de los flex items dentro del flex container. Vale la pena observar que con el modelo flexbox la posición de los elementos se determina en el elemento contenedor y no en los mismos elementos, que es como se haría usualmente en otros modelos de maquetación.

Para distribuir los flex items dentro del flex container se puede utilizar la propiedad <code>justify-content</code> que entre otros puede tener los siguiente valores para determinar la distribución de los elementos:
<ul>
  <li>center</li>
  <li>flex-start</li>
  <li>flex-end</li>
  <li>space-around</li>
  <li>space-between</li>
</ul>

Por ejemplo, alineando los flex items en el centro del flex container, para esto usamos el valor <code>center</code>:

```html
<body>
  <div class="flex-container">
    <div class="flex-item">1</div>
    <div class="flex-item">2</div>
    <div class="flex-item">3</div>
  </div>
</body
```
```css
.flex-container{
  display:flex;
  justify-content:center;
  background-color:rgba(251, 80, 80, 0.17);
}

.flex-container > .flex-item{
  padding:50px;
  border:2px solid black;
  margin: 10px;
  background-color:rgba(80, 163, 251, 0.38);
  font-size:20px;
}
```

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/7fbc5c8f-3a3b-4479-821a-7b873c51ec9b)

Aquí es importante anotar que a pesar de que los flex items son elementos de bloque (son elementos div) flexbox los acomoda como si fueran elementos de línea.

Pro ejemplo si quisieramos distribuir los flex items en todo el ancho del flex container de forma que queden igualmente espaciados, usamos <code>justify-content</code> con el valor <code>space-around</code>:

```html
<body>
  <div class="flex-container">
    <div class="flex-item">1</div>
    <div class="flex-item">2</div>
    <div class="flex-item">3</div>
  </div>
</body
```
```css
.flex-container{
  display:flex;
  justify-content:space-around;
  background-color:rgba(251, 80, 80, 0.17);
}

.flex-container > .flex-item{
  padding:50px;
  border:2px solid black;
  margin: 10px;
  background-color:rgba(80, 163, 251, 0.38);
  font-size:20px;
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/d9c71816-35a5-456a-b9fa-7d51475bbc39)

Anteriormente habíamos comentado que un elemento puede tener los roles de flex container y flex item al mismo tiempo, esta característica es de bastante utilidad ya que un flex container solo puede detallar la distribución de sus descendientes directos, así que las distribuciones no aplican a descendientes de mayor nivel de anidamiento. Por lo tanto los flex items que tengan descendientes deben a su vez ser flex containers para que sea posible posicionar esos descendientes utlizando el modelo flexbox.

Sigamos con el ejemplo que venimos trabajando pero en esta ocasión agrupemos dos de los flex items en un nuevo contenedor que servirá como flex container para ellos y como flex item para el flex container orginal.

```html
<body>
  <div class="flex-container">
    <div class="flex-item">1</div>
    <div class="new-container">
      <div class="flex-item">2</div>
      <div class="flex-item">3</div>
    </div>
  </div>
</body
```
```css
.flex-container{
  display:flex;
  justify-content:space-between;
  background-color:rgba(251, 80, 80, 0.17);
}

.new-container{
  display:flex;
  justify-content:space-between;
}

.flex-item{
  padding:50px;
  border:2px solid black;
  margin: 10px;
  background-color:rgba(80, 163, 251, 0.38);
  font-size:20px;
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/d595f911-e479-4045-86da-bde6e3c09012)

De momento hemos visto que con la propiedad <code>justify-content</code> podemos controlar la distribución de los flex items a lo ancho del flex container, esto es generalmente útil cuando la altura (height) del flex container no es significativamente diferente a la de los flex items, tal como vimos en los ejemplos anteriores. Pero en los casos donde el flex container tiene una altura definida y que es comparable a la de los flex items va a ser necesario una propiedad que nos permita alinear también de manera vertical los flex items en el espacio del flex container. Esto último lo podemos lograr con la propiedad <code>align-items</code> que nos va a permitir mover los flex items desde arriba hasta abajo del flex container (y también estirarlos) según nuestra necesidad.

Los posibles valores para la propiedad son los siguientes:
<ul>
  <li>center</li>
  <li>flex-start</li>
  <li>flex-end</li>
  <li>stretch</li>
  <li>baseline</li>
</ul>

Veamos un ejemplo centrando los flex items en medio del flex container:

```html
<body>
  <div class="flex-container">
    <div class="flex-item">1</div>
    <div class="flex-item">2</div>
    <div class="flex-item">3</div>    
  </div>
</body>
```
```css
.flex-container{
  display:flex;
  justify-content:space-between;
  background-color:rgba(251, 80, 80, 0.17);
  height: 500px;
  align-items:center;
}

.flex-item{
  padding:50px;
  border:2px solid black;
  margin: 10px;
  background-color:rgba(80, 163, 251, 0.38);
  font-size:20px;  
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/34ef582f-e688-4817-83c1-14231d790696)

Poniendo los flex items en la parte baja del flex container:

```css
.flex-container{
  display:flex;
  justify-content:space-between;
  background-color:rgba(251, 80, 80, 0.17);
  height: 500px;
  align-items:flex-end;
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/47a2bc5e-cbfd-4b5e-9413-78c0c53a536c)

O estirandolos en toda la altura del flex container:

```css
.flex-container{
  display:flex;
  justify-content:space-between;
  background-color:rgba(251, 80, 80, 0.17);
  height: 500px;
  align-items:stretch;
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/13c546b3-f034-48e9-be78-13b8b3d5d5ae)

Hasta ahora nos hemos enfocado en flex items se se despliegan a lo ancho del flex container como formando filas donde cada flex items se ubica justo al lado del otro hasta ocupar todo el ancho del flex container. Es posible cambiar la dirección de disposición de los flex items para que también puedan desplegarse de manera vertical, esto se logra con una simple propiedad: <code>flex-direction</code> que permite indicar si los flex items se distribuyen en forma de fila horizontal o de columna vertical.

Algo que se debe considerar cuando se decide distribuir los flex items en forma de columna y es que las propiedades <code>justify-content</code> y <code>align-items</code> intercalan su funcionalidad, es decir la primera actuará tal cual lo hace cuando los flex items están en fila, es decir, <code>justify-content</code> va a controlar la posición de los flex items en dirección vertical, mientras que <code>align-items</code> va a permitir desplazar la columna a lo ancho (en la dirección horizontal) del flex container.

En el siguiente ejemplo con la propiedad <code>align-items</code> ubicamos la columna con los flex items en el centro del flex container mientras que con <code>justify-content</code> hacemos la distribución vertical en la columna.

```html
<body>
  <div class="flex-container">
    <div class="flex-item">1</div>
    <div class="child-container">
      <div class="flex-item">2</div>
      <div class="flex-item">3</div>
    </div>
  </div>
</body>
```
```css
.flex-container{
  display:flex;
  justify-content:space-between;
  background-color:rgba(251, 80, 80, 0.17);
  height: 700px;
  align-items:center;
  justify-content:space-between;
  flex-direction:column;
}

.flex-item{
  padding:50px;
  border:2px solid black;
  margin: 10px;
  background-color:rgba(80, 163, 251, 0.38);
  font-size:20px;
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/23e5f543-7f3f-45fc-8649-8f85bf78e648)

¿Recuerdas que con el modelo de maquetación grid podíamos fraccionar el contenedor en partes iguales y luego asignar fracciones a los contenidos para tener cuadrículas en las que no necesariamente los elementos tenían las mismas dimensiones? Pues con el modelo flexbox también se puede hacer algo semejante, podemos definir individualmente el ancho que ocuparán los flex items dentro del flex container con la ayuda de la propiedad <code>flex</code>. Esta propiedad se utiliza en los flex items para indicar que tanto se debe "estirar" su ancho para ocupar cierta cantidad de espacio dentro del flex container. Por ejemplo, un flex item con un valor de 2 en la propiedad <code>flex</code> ocupará dos veces el ancho que ocuparía un flex item con un valor de 1.

```html
<body>
  <div class="flex-container">
    <div class="flex-item" id="one">1</div>
    <div class="flex-item" id="two">2</div>
    <div class="flex-item" id="three">3</div>    
  </div>
</body>
```
```css
.flex-container{
  display:flex;
  justify-content:space-between;
  background-color:rgba(251, 80, 80, 0.17);
}

#one, #three{
  flex: 1;
}

#two{
  flex:2;
}

.flex-item{
  border:2px solid black;
  margin: 10px;
  background-color:rgba(80, 163, 251, 0.38);
  font-size:20px;
  text-align:center;
}


```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/ae1fcd8a-233c-4e4f-a59a-368170cf2422)

Para dar por concluída esta sesión vamos a mencionar que con el modelo flexbox también es posible ubicar los elementos en cuadrículas como aprendimos a hacerlo con el modelo grid. Con la propiedad <code>flex-wrap</code> vamos a lograr que de forma automática los flex items se ubiquen en una nueva fila del flex container una vez desbordan la capacidad su ancho.

```html
<body>
  <div class="flex-container">
    <div class="flex-item">1</div>
    <div class="flex-item">2</div>
    <div class="flex-item">3</div>
    <div class="flex-item">4</div>
    <div class="flex-item">5</div>
    <div class="flex-item">6</div>    
  </div>
</body>
```
```css
.flex-container{
  display:flex;
  justify-content:center;
  background-color:rgba(251, 80, 80, 0.17);
  flex-wrap: wrap;
}

.flex-item{
  border:2px solid black;
  margin: 10px;
  background-color:rgba(80, 163, 251, 0.38);
  font-size:20px;
  text-align:center;
  padding:170px;
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/2fd22e07-3d72-46b7-8afc-4a1327de1437)












