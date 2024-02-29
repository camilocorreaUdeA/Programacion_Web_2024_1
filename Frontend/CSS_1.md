# CSS

## Introducción

CSS es el acrónimo para <i>Cascading Style Sheets</i>, hojas de estilo en cascada. Es un lenguaje para declarar y expresar el aspecto visual de un documento HTML, es definido y mantenido por el W3C (<i>World Wide Web Consortium</i>) y a la fecha su versión más reciente es la 4.15 (Diciembre de 2020) aunque a partir de la versión 3 a todas las nuevas versiones se les conoce como CSS3, a pesar de ser un error ya que al lenguaje se le debe conocer solo como CSS sin ninguna versión, de hecho, de acuerdo con el W3C, no se espera que aparezca un CSS4, simplemente CSS.

Hasta el momento hemos visto cómo definir la estructura de una página web utilizando el lenguaje de etiquetas HTML pero para ser honestos la apariencia de esas páginas es bastante plana y además aburrida, y tratar de hacer cualquier mejora en ese aspecto es una tarea bastante difícil que requiere de unos conocimientos especializados de HTML.

Afortunadamente existe CSS, que es un lenguaje que permite que nuestros sitios web se vean atractivos y más agradables a la vista del usuario. A continuación un listado breve de las principales características que se pueden lograr con CSS:

<ul>
  <li>Controlar la apariencia de los elementos HTML, de forma que en el navegador se presentan con el diseño que se haya elegido.</li>
  <li>Mejorar la presentación de un documento al permitir la división del mismo en distintas secciones que se pueden ubicar en diferentes zonas de la pantalla.</li>
  <li>Modificar el tamaño o la orientación de los elementos.</li>
  <li>Definir estilo para los textos: color, tamaño y tipo de la letra, alineación, formato, etc.</li>
  <li>Controlar efectos visuales sobre las imágenes como la opacidad, sombras, bordes, etc.</li>
  <li>Crear menús de navegación más dinámicos y con una apariencia más fresca.</li>
  <li>Controlar efectos de transición sobre los elementos de acuerdo con los eventos a los que se vean sometidos como al dar clic, pasar el mouse por encima, seleccionar, etc.</li>
</ul>

## Conceptos básicos de CSS

### Modelo de cajas

Uno de los principales objetivos, sino el principal, en el diseño de un sitio web es el de cumplir con un esquema (estructura) del sitio con los elementos bien posicionados dentro del mismo. Hoy en día es bastante raro encontrar sitios web en los que los elementos van uno tras otro como si estuvieran apilados (sí, es verdad, los ejercicios que hemos hecho hasta el momento lucen exactamente así ;-)).

Dicho lo anterior, deducimos que una de las grandes fortalezas de CSS reside en que nos otorga la habilidad de poder definir el esquema que queramos para nuestro sitio web y por tanto nos va a permitir ubicar nuestras imágenes, gifs, vídeos, mapas, menús y otros contenidos justo en el lugar donde nosotros queramos. Así que si queremos llegar a ese nivel de diseño tenemos que empezar por un concepto básico pero sumamente poderoso de CSS: El modelo de cajas.

Este concepto del modelo de cajas se basa en una premisa bastante sencilla y es que todos los elementos HTML de un sitio web están modelados como si fueran una caja, tanto los elementos normales, es decir los que tienen los 3 elementos constitutivos (apertura, contenido y cierre) como los elementos vacíos (que solo tienen apertura).

Además basados en el principio que vimos anteriormente de que podemos tener elementos anidados en otros entonces podemos tener cajas dentro de otras cajas, luego, llegamos a la conclusión que el elemento body es la caja más grande dentro de nuestro sitio web y que las cajas que representan a otros elementos están todas dentro de esta.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/1405857c-e05a-4ba9-8b46-f92da45b86d1)

Puede ver la abstracción del modelo de cajas con mayor claridad con ayuda del siguiente [ejemplo en CodePen](https://codepen.io/camilocorreaUdeA/pen/BavzOLd)

En la imagen (y en el ejemplo también) podemos observar que las separaciones entre las cajas son distintas, la separación entre el elemento <code>h1</code> y el elemento <code>p</code> es mayor a la de este último y el elemento <code>ul</code>, y las separaciones entre las cajas de los elementos <code>li</code> son apenas perceptibles.

Otro detalle observable es que las cajas ocupan todo el ancho de la pantalla, a excepción de las de los elementos <code>li</code>, que están alineadas a la derecha ya que las viñetas (<i>bullet points</i>) están en su propia caja, y todas estas cajas a su vez están contenidas dentro de la caja del elemento <code>ul</code>.

Estos detalles que acabamos de observar tienen explicación y es que cada caja tiene un conjunto de propiedades que van a ser diferentes dependiendo del elemento al que pertenezcan. Dichas propiedades son el margen interno o <b><i>padding</i></b>, el borde, <b><i>border</i></b>, y el margen externo o simplemente <b><i>margin</i></b>. La descripción de cada una de estas propiedades a continuación:

<ul>
  <li>Padding: Es el espacio entre el contenido del elemento, por ejemplo el texto de un elemento <code>p</code>, y el borde (border) de la caja. También conocido como margen interno.</li>
  <li>Border: Es el borde que limita la dimensión de la caja. Es técnicamente un espacio que separa el padding y el margin.</li>
  <li>Margin: Es la distancia de guarda entre los bordes (border) de la caja del elemento y de las cajas de los elementos contiguos (vecinos) del documento HTML. También conocido como margen externo.</li>
</ul>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/3534cc3c-4469-4cd5-9978-cef2391f4a7f)

Un modelo de cajas alternativo propone que cada una de las propiedades del elemento tiene asociada una caja, entonces se tienen cajas para el contenido, para el padding, para el borde y para el margin. 

Se asemeja mucho este modelo a las matrioskas o muñecas rusas que cada una contiene a otra que a su vez contiene a otra hasta llegar a la más pequeña. En este caso la caja del margin contiene a la caja de padding y esta a su vez contiene a la caja del contenido.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/83e70c37-0083-400a-87ff-75a58339fd26)

Más adelante veremos que con CSS podemos modificar las propiedades de la caja para cambiar el tamaño o la posición de los elementos de un documento HTML.

### Reglas de una hoja de estilo CSS

Por lo general, una hoja de estilos de CSS está conformada por un conjunto de reglas CSS que tienen una sintaxis bien definida conformada por un selector, que permite identificar sobre qué elementos del documento se debe aplicar la regla, y varias declaraciones en las que a una propiedad del elemento se le da un valor específico. La sintaxis general es como vemos aquí a continuación:

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/d39829e6-9021-425e-b5f7-d12ffb774273)

Por ejemplo, si queremos definir una regla CSS que cambia el color de todos los textos de los elementos <code>p</code> que hayan en el documento, debemos hacerlo de la siguiente manera: el selector sería el nombre del elemento en este caso <code>p</code>, y luego entre llaves (curly braces) declaramos las propiedades que vayamos a modificar, en este caso el color del texto que en CSS se denomina <code>color</code>, y separado por dos puntos (<code>:</code>) declaramos el valor de la propiedad, en este caso el color que vamos a dar al texto (más adelante aprenderemos el tema de los colores en CSS) y finalizamos la declaración con punto y coma (<code>;</code>).

```css
p {
  color:purple;
}
```
Vale la pena aclarar que esta sintaxis es válida para los siguientes casos: cuando la hoja de estilos CSS se desarrolla en un archivo distinto al documento HTML, un archivo de extensión <i>.css</i>; y el otro caso aplicable sería cuando las reglas se desarrollan como contenido del elemento <code>style</code> anidado en el elemento <code>head</code> del documento HTML.

El caso excepcional en el que está sintaxis no es válida es cuando las declaraciones se hacen en el atributo <code>style</code> del elemento en particular (porque se omite el selector, solo se usan las declaraciones).

A continuación vamos a ver la diferencia entre las 3 formas de agregar las reglas de estilos CSS a un documento HTML.

### Asociando hojas de estilo CSS a un documento HTML

CSS Externo: Reglas en un archivo externo .css

Lo primero es confeccionar un archivo de extensión .css, por lo general se le llama styles, en el que se definen las reglas que van a aplicar a los elementos del documento.

Luego en el elemento head del documento HTML donde se va a aplicar la hoja de estilos, utilizamos el elemento <code>link</code> para asociar ambos archivos.

<img width="550px" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/fad38d14-40d4-4145-babe-b1d828647f07">

```html
<head>
  <link rel="stylesheet" type="text/css" href="styles.css" />
</head>
```
CSS interno: Reglas definidas en el head

En esta alternativa se definen las reglas como contenido del elemento <code>style</code> que va en el <code>head</code> del documento HTML.

```html
<head>
  <style type="text/css">
    p {
     color:blue;
     font-size:10px;
     text-align:right;
    }

   ul {
     color:grey;
   }
  </style>
</head>
```
CSS interno: Reglas definidas en línea (inline) en los elementos.

En este caso se utiliza el atributo <code>style</code> del elemento para declarar las reglas de estilos. No se debe usar la sintaxis para las reglas que vimos anteriormente, solamente las declaraciones con el nombre de la propiedad y el valor correspondiente.

```html
<p style="color:blue; font-weight:bold;">
  Este es un párrafo con el texto en color azul y letra gruesa
</p>
```
¿Cuál de las formas de agregar estilos es la más adecuada entonces?

En términos de resultados y funcionamiento las 3 alternativas van a funcionar y verás idénticos resultados al usar una o alguna de las otras.

El aspecto fundamental para escoger la que mejor se adapte a tus necesidades va a ser el mismo contexto de la aplicación web. Por contexto me refiero a las características específicas del sitio, como por ejemplo si los documentos HTML son muy extensos, si el sitio está formado por varios hipertextos al mismo tiempo o si está sometido a que se le hagan cambios regularmente, entre muchas otras pero siendo estas que he nombrado las más relevantes.

Si el documento HTML es muy extenso y el sitio web está compuesto por varias páginas entonces es más recomendable tener las hojas de estilo en archivos aparte para que de ese modo se puedan manejar como componentes adicionales que se pueden mantener y modificar por separado sin necesidad de hacer cambios en los documentos HTML.

Para sitios web de una sola página aplican tanto el archivo externo de hojas de estilo como el CSS interno definido en el head. Mientras que los estilos en línea son importantes cuando se requiere hacer modificaciones a elementos específicos del documento, el problema es que a medida que crece el volumen de estilos aplicados de esta manera se hacen más difícil de mantener porque hay que buscar el elemento específico donde se debe modificar el estilo y esto implica cambios directos en el documento HTML.

La recomendación es usar un archivo externo para las hojas de estilo CSS.

### Selectores: tipo, id y class

Tipo

Un selector de tipo, también conocido como selector de etiqueta o de elemento permite seleccionar un elemento o etiqueta específica del documento HTML. Este selector aplica la regla de estilo a todos los elementos del documento HTML que estén definidos con la etiqueta del selector, por ejemplo si el selector de la regla de estilo es <code>p</code>, entonces a todos los párrafos del documento se les aplica las declaraciones de la regla.

```css
/* Esta regla aplica a todos los elementos h1 del documento */
h1 {
  text-align:center;
  color: blue;
}
```
Dentro de esta categoría también se incluye al selector universal que se denota con un asterisco (<code>*</code>). Este selecciona todo los elementos dentro de un documento (o todos los elementos anidados dentro de otro elemento, si se usa un selector de encadenamiento, que veremos más adelante).

```css
/* Esta regla aplica a todos los elementos del documento */
* {
  margin:0;
}
```
Id

Como ya se había visto en las sesiones anteriores de HTML (específicamente cuando vimos las anclas para el elemento <code>a</code>) se puede asignar un identificador único a un elemento utilizando el atributo <code>id</code>, sí, ningún otro elemento debe tener un identificador igual.

CSS nos permite utilizar estos identificadores como selectores en las reglas, por tanto las reglas definidas con ese selector solo aplicarán para el elemento que tiene ese identificador. Se utiliza el símbolo numeral (<code>#</code>), seguido del identificador para indicar a CSS que la regla tiene un selector por <code>id</code>.

```html
<!-- Un elemento con identificador -->
<p id="identificador">Un párrafo cualquiera</p>
```

```css
/* Esta regla aplica al párrafo con ese identificador */
#identificador {
     font-family:Courier;
     font-size:20px;
}
```
Class

Este selector funciona de forma similar al de <code>id</code>, solo que en vez de utilizar ese atributo vamos a utilizar el atributo <code>class</code> y además, hay una gran diferencia, y es que varios elementos del documento pueden compartir el mismo valor de ese atributo. Por tanto la regla que utilice este selector se puede aplicar para múltiples elementos del documento que tengan el mismo valor en el atributo <code>class</code>.

En la regla se utiliza el punto (<code>.</code>), seguido del valor dado al atributo <code>class</code> para indicar a CSS que la regla tiene un selector por <code>class</code>.

```html
<!-- Elementos con atributo class -->
<li class="elemento-impar">Parte de una lista cualquiera</li>
<li>Parte de una lista cualquiera</li>
<li class="elemento-impar">Parte de una lista cualquiera</li>
```

```css
/* Esta regla aplica solo a dos elementos de la lista */
.elemento-impar {
     font-family:Courier;
     font-size:20px;
}
```
Un elemento puede tener varios valores para el atributo class al mismo tiempo, basta con separar los valores con un espacio, luego para especificar una regla que aplica para un elemento que tiene cierto conjunto de valores para el atributo class en particular, va a ser necesario especificar estos valores en el selector para que la regla quede aplicada donde realmente se desea.

```html
<!-- Elementos con atributo class -->
<p class="info">Párrafo de class info</p>
<p class="info advertencia">Párrafo de class info y class advertencia</p>
```

```css
/* Esta regla aplica a ambos párrafos */
.info {
     font-family:Verdana;
}

/* Y esta regla aplica solo al segundo párrafo */
.info.advertencia {
     text-decoration:underline;
}
```





