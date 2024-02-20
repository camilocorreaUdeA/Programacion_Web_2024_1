# HTML

## Introducción

<p>El <i>front-end</i> o capa de presentación de una aplicación web hace referencia al software que se desarrolla usando el conjunto de tecnologías <b>HTML</b>, <b>CSS</b> y <b>Javascript</b>. Con <b>HTML</b> se define la estructura de la presentación del contenido de la página web, <b>CSS</b> permite dar estilo y personalizar la parte visual de la página, mientras que <b>Javascript</b> permite agregar lógica y funcionalidad dinámica.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/ba588505-7830-4823-b38d-ff8a7ba9c4be)

<p>A partir de esta presentación y las siguientes vamos a enfocarnos en aprender HTML, <em>HyperText Markup Language</em> (Lenguaje de Marcado de Hipertexto), la anatomía general de una página web, el conjunto de etiquetas que componen el lenguaje HTML y cómo utilizarlas para estructurar el contenido de un sitio web.</p>

<p>HTML no es un lenguaje de programación ya que no permite codificar lógica de programación ni describir algoritmos o funcionalidades, por el contrario HTML solo se encarga de definir la forma en que se presentan los contenidos de un sitio web, y para ese fin utiliza un conjunto de etiquetas que permiten marcar o señalizar el contenido del sitio web para que el navegador sepa cómo mostrarlo en pantalla.</p>

<p>Cuando nos referimos a la forma de estructurar el contenido estamos hablando de muchas cosas: dividir el texto en párrafos, agregar estilo al texto (cursiva, negrilla, resaltado, etc.), definir hipervínculos, agregar imágenes, videos y animaciones, definir listas, tabular información, definir el esquema de la página web, etc.</p>

<p>Para poder definir esa estructura HTML permite la marcación del texto con etiquetas que el navegador interpreta para poder visualizar dicha estructura en pantalla. Dichas etiquetas son la unidad fundamental de HTML, permiten encerrar o delimitar parte del contenido para crear elementos HTML (que vamos a definir más adelante) que son los que dan la estructura y estilo al contenido.</p>

<p>De acuerdo con el tipo de elemento que permite crear se puede agrupar las etiquetas de HTML en las siguientes categorías:</p>

<ul>
  <li><b>Raíz principal</b>: Sólo hay una etiqueta en esta categoría y define el elemento de mayor jerarquía en el documento HTML.</li>
  <li><b>Documentación y metadatos</b>: Los metadatos proporcionan información de la página web, por ejemplo los estilos y scripts necesarios para renderizar la página en el navegador.</li>
  <li><b>Cabeza y cuerpo de la página</b>: Las etiquetas de esta categoría engloban en sus interior los elementos que se despliegan o no en el navegador como por ejemplo el texto y los metadatos respectivamente.</li>
  <li><b>Organización del contenido en secciones</b>: Permiten organizar el contenido en partes lógicas permitiendo crear un esquema general de la página web proporcionando elementos que permiten identificar secciones particulares del contenido.</li>
  <li><b>Contenido textual</b>: Permiten organizar en bloques o secciones el contenido que se ubica en el cuerpo del documento HTML.</li>
  <li><b>Semánticas de texto</b>: Permiten definir el significado, estructura o estilo de cualquier pieza de texto del contenido.</li>
  <li><b>Imágenes y multimedia</b>: Brindan soporte a distintos recursos multimedia como imágenes, video y audio.</li>
  <li><b>Contenido incrustado</b>: Permiten incrustar contenido externo en un punto específico del documento HTML, contenido como otras páginas web, contenido interactivo basado en plugins, etc.</li>
  <li><b>Scripting</b>: Para crear contenido dinámico e interfaces a otras aplicaciones web.</li>
  <li><b>Tabular contenido</b>: Permiten crear y manejar contenido en forma de tablas.</li>
  <li><b>Formas</b>: Son los elementos que permiten la creación de formularios que un usuario del sitio puede llenar para enviarlos a la aplicación web.</li>
  <li><b>Demarcar ediciones</b>: Permiten indicar partes donde el texto ha sufrido alteraciones.</li>
  <li><b>SVG y MathML</b>: Permiten incrustar contenido SVG y MathML directamente en el documento HTML.</li>
  <li><b>Elementos interactivos</b>: Ayudan en la creación de objetos interactivos de interfaz de usuario.</li>
  <li><b>Componentes web</b>: Permiten agregar contenido relacionado con la tecnología HTML de Web Components.</li>
</ul>

## Anatomía de un documento HTML

### Elementos HTML

<p>Un elemento dentro de un documento HTML está compuesto de la siguiente manera:</p>
<ol>
  <li><b>La etiqueta de apertura</b>: Consiste en el nombre del elemento encerrado entre 
paréntesis angulares (<>). Esta etiqueta marca el lugar a partir del cual el elemento empieza a hacer efecto.
</li>
  <li><b>El contenido</b>:Es el contenido como tal del elemento, por ejemplo el texto en un párrafo. </li>
  <li><b>La etiqueta de cierre</b>: Es exactamente igual a la etiqueta de apertura con la única diferencia de que tiene un forward slash (/) antes del nombre del elemento. Marca el punto hasta donde es válido el efecto del elemento.</li>
</ol>
<p>Así las cosas, un elemento HTML es una etiqueta de apertura seguida del contenido y este a vez seguido de una etiqueta de cierre.</p>
<img width="601" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/cfc5c650-7067-477e-bd89-3f2de64562fc">

### Anidación de elementos HTML
<p>HTML permite ubicar elementos al interior de otros elementos, a esto se le conoce como anidamiento. Es importante tener en cuenta que en el momento de anidar se debe conservar el orden en la ubicación de las etiquetas de cierre para que el navegador no vaya a tener problemas al tratar de presentar el contenido.</p>

```html
<p>Este es un <em>texto</em> de nuestro sitio web</p>
```
### Elementos HTML vacíos
<p>Son elementos que no siguen el patrón del que venimos hablando etiqueta de apertura, contenido y etiqueta de cierre y es porque no tienen contenido, entonces no tiene sentido tener etiquetas para limitarlo, de ahí el nombre de elementos vacíos.</p>

```html
<img src="imagenes/web/icon.png" alt="icono UI"/>
<input type="text"/>
```
<p>No es estrictamente necesario agregar el forward slash (/) al final, antes del paréntesis angular de cierre, pero por consistencia y compatibilidad con XML se suele agregar siempre que sea posible.</p>

### Atributos de un elemento HTML
<p>Los elementos HTML pueden tener atributos que proporcionan información adicional acerca del elemento, estos atributos no influyen en el efecto directo del elemento sobre el contenido.</p>

<p>Para definir un atributo en un elemento HTML se deja un espacio luego del nombre del elemento y se agrega el nombre del atributo seguido de un signo de igual (=) y luego entre comillas se agrega el valor del atributo.</p>

<img width="602" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/a14d26f7-0dc4-469a-8443-2c4ee4f71d4c">

<p>Existen algunos atributos conocidos como atributos <b><i>booleanos</i></b> que son aquellos que solo aceptan un valor que coincide en ser el mismo nombre del atributo, de modo que es aceptable expresar el atributo sin otorgarle valor alguno.</p>

```html
<input type="text" disabled="disabled"/>
<!-- estos dos elementos input son equivalentes -->
<input type="text" disabled/>
```
<p>Un elemento puede tener múltiples atributos, lo importante es separar los atributos por un espacio para poder diferenciarlos de los otros. No olvidar tampoco las comillas en los valores de los atributos para evitar problemas con los atributos a continuación.</p>

### Estructura básica de un documento HTML
<p>La combinación de elementos HTML individuales permite conformar un documento HTML o página web, sin embargo existe un esquema general que permite definir la estructura básica de un documento HTML que podemos ver a continuación:</p>

```html
<!doctype html>
<html lang="eng-US">
  <head>
    <meta charset="utf-8" />
    <title>Mi sitio web</title>
  </head>
  <body>
    <p>Bienvenido a mi sitio web</p>
  </body>
</html>
```
<p>El <i>doctype</i></p>
<p>Todos los documentos HTML comienzan con la declaración <code>!doctype html</code>”. El propósito de esta etiqueta es indicar al navegador la versión del lenguaje HTML que debería usarse para desplegar el documento. Nosotros no vamos a complicarnos y la vamos a utilizar tal y como la vemos actualmente ya que por defecto indica que se use la versión más reciente del leguaje, es decir, HTML5.</p>

<p>Elemento <i>html</i></p>

<p>El elemento HTML proporciona la raíz principal del documento HTML, es decir, que cualquier otro elemento que haga parte del documento debe ir anidado al interior de este. A este elemento es común verlo acompañado del atributo <b>lang</b> que sirve para mejorar la accesibilidad del documento ya que le indica a los lectores de pantalla utilizados por personas con problemas o limitaciones visuales el idioma en el que está la página web.</p>

<p>Elemento <i>head</i></p>

<p>Este elemento actúa como un contenedor para todos aquellos elementos que se van a incluir en el documento HTML pero que no hacen parte del contenido que se despliega en el navegador. Estos elementos incluyen las palabras claves y descripciones para los motores de búsqueda y otras aplicaciones web, las referencias a las hojas de estilo en CSS y scripts en Javascript que va a utilizar el documento, el título para la pestaña del navegador y la información para la sección de favoritos del navegador y el conjunto de caracteres que puede utilizar el texto del documento  entre otros.</p>

<p>Elemento <i>body</i></p>

<p>Es el elemento que contiene todo el contenido que se va a desplegar en la página web y sera visible en el navegador, esto incluye texto, hipervínculos, imágenes, multimedia, etc.</p>

<p>Comentarios en un documento HTML</p>

<p>Para agregar comentarios se utilizan los marcadores <code>&lt;!--</code> y <code>--&gt;</code> para delimitar en su interior el comentario.</p>

```html
<p>Esto no es un comentario</p>
<!-- esto es un comentario en html -->
```
<p>Manejo de espacios y caracteres especiales en HTML</p>
<p>Los espacios no son tenidos en cuenta y para incluir caracteres especiales se deben utilizar referencias conocidas como entidades que empiezan con un ampersand (&) seguido de un código en texto y terminadas con punto y coma (;). En este enlace puede ver la lista de entidades usadas en HTML: https://html.spec.whatwg.org/multipage/named-characters.html#named-character-references</p>

### ¿Y qué se pone en el elemento <i>head</i> de un documento html?
<p>Como ya se había mencionado, el contenido del elemento head no se despliega en la pantalla y por el contrario este contenido sirve para informar acerca del documento a través de metadatos.</p>

<p>Dependiendo del sitio web el contenido del head puede ser muy extenso o no, a continuación se muestran los elementos que se pueden encontrar con mayor regularidad en esta sección del documento HTML.</p>

<p><b>Título</b>: El elemento <code>&lt;title&gt;</code> del head se usa para dar un título a la pestaña del navegador donde se despliega la página web. También se utiliza este título cuando se agrega la página a la sección de favoritos del navegador.</p>

<p><b>Metadatos</b>: Los metadatos son datos que describen o detallan a otros datos, en esta caso a una página web. Para especificar metadatos se utiliza el elemento vacío <code>&lt;meta&gt;</code>, recuerde que un elemento vacío no tiene contenido, por tanto solo tiene etiqueta de apertura y no tiene contenido ni etiqueta de cierre.</p>

<p>Para especificar la información disponible en los metadatos se utilizan atributos en el elemento <code>&lt;meta&gt;</code>.</p>

<p>Por ejemplo para especificar el conjunto de caracteres para codificar el texto en la página se utiliza el atributo <i>charset</i> y el valor en este caso es el conjunto de caracteres que la página va a utilizar. Por lo general el valor que se da es <i>utf-8</i> que es un conjunto de caracteres universal que tiene casi todos los caracteres de todos los idiomas de la humanidad.</p>

```html
<head>
  <meta charset="utf-8" />
</head>
```
<p>El elemento <code>&lt;meta&gt;</code> también utiliza los atributos <i>name</i> y <i>content</i> para definir otros metadatos, por ejemplo si quiere informar acerca del autor del documento HTML:</p>

```html
<head>
  <meta name="author" content="Universidad de Antioquia" />
</head>
```
<p>O si quiere agregar una descripción del contenido de la página web para que salga en los motores de búsqueda en Internet:</p>

```html
<head>
  <meta
     name="description"
     content="Sitio web de la Universidad de Antioquia. Alma Mater de los antioqueños" />
</head>
```
<p>Iconos del sitio web: En el head también se puede agregar una referencia a un icono que se puede usar al igual que el título en la pestaña del navegador y en los favoritos o marcadores del navegador. Para este propósito se utiliza el elemento vacío <code>&lt;link /&gt;</code> con los atributos <i>rel</i>, <i>href</i> y <i>type</i>.</p>

```html
<head>
  <link rel="icon" href="favicon.ico" type="image/x-icon" />
</head>
```
<p>Hojas de estilo CSS y scripts de Javascript: Las hojas de estilo permiten mejorar la apariencia y presentación del sitio web, mientras que los scripts proporcionan interactividad y funcionalidad dinámica. Para referenciar una hoja de estilo se utiliza el elemento <code>&lt;link /&gt;</code> y sus atributos de forma similar a como se acaba de ver para el icono del sitio web:</p>

```html
<head>
  <link rel="stylesheet" href="mis-estilos-css.css" />
</head>
```
<p>También se pueden incluir hojas de estilo con el elemento <code>&lt;style&gt;</code> directamente en el <i>head</i></p>

```html
<head>
  <style>
    div {
      color: white;
      background-color: black;
    }

    p {
      color: red;
    }
  </style>
</head>
```
<p>Para los scripts se utiliza el elemento <code>&lt;script&gt;</code>, que no es un elemento vacío, pero que pareciera serlo porque no se utiliza contenido para referenciar al script de javascript sino un atributo del elemento. Se recomienda incluir el atributo booleano defer para que el script sea lo último que se carga en la página luego del HTML y las hojas de estilo.</p>

```html
<head>
  <script src="mi-script.js" defer></script>
</head>
```
### ¿Y qué va en el <i>body</i> de un documento HTML?
<p>En el elemento <code>&lt;body&gt;</code> de un documento HTML va todo el contenido que se va a desplegar en el navegador. Este elemento funciona como un nodo del cual se desprenden muchos otros elementos que están anidados en su contenido y hacen parte del mismo.</p>

<p>En el <i>body</i> se ubican el texto, hipervínculos, imágenes y otros recursos multimedia que conforman una página web. A partir de esta sección vamos a estudiar los elementos que hacen parte del contenido del elemento body de una página web, y aprovechando que el texto es el recurso de mayor presencia en un sitio web empezaremos por conocer las etiquetas para manipular texto y para crear hipervínculos.</p>

<p>Encabezados y párrafos</p>
<p>En HTML los párrafos son el contenido del elemento <p>. El texto encerrado entre las etiquetas de apertura y cierre de este elemento hace parte de un párrafo individual del texto.</p>

```html
<body>
  <p>El elemento <p> sirve para definir parrafos de texto en una página web</p>
</body>
```
<p>El elemento <code>&lt;p&gt;</code> soporta los atributos globales y de eventos de las etiquetas HTML (https://www.w3schools.com/tags/ref_standardattributes.asp, https://www.w3schools.com/tags/ref_eventattributes.asp), siendo los más utilizados con este elemento los siguientes: class, id, lang, title, style y hidden.</p>

<p>Los encabezados se construyen utilizando el elemento <code>&lt;h#&gt;</code> donde # es un número entre 1 y 6. Cada número representa un nivel distinto de jerarquía del encabezado en el texto, siendo <code>&lt;h1&gt;</code> el encabezado principal del texto, <code>&lt;h2&gt;</code> un sub-encabezado y así sucesivamente hasta <code>&lt;h6&gt;</code> que es el de menor orden jerárquico.</p>
  
<p>Se recomienda que en todo el documento HTML solo se utilice un elemento <code>&lt;h1&gt;</code> y que se conserve correctamente el orden jerárquico de los sub-encabezados, de modo que uno de menor rango no preceda a uno de mayor jerarquía, por ejemplo un <code>&lt;h3&gt;</code> precediendo a un <code>&lt;h2&gt;</code>.</p>

<p>Al igual que en el caso del elemento <p> los encabezados <code>&lt;h1&gt;...&lt;h6&gt;</code> soportan los atributos globales y de eventos de las etiquetas HTML</p>

```html
<body>
  <h1>Encabezado 1</h1>
  <h2>Sub-encabezado 2</h2>
  <h3>Sub-sub-encabezado 3</h3>
  <h4>Sub-sub-sub-encabezado 4</h4>
  <h5>Sub-sub-sub-sub-encabezado 5</h5>
  <h6>Sub-sub-sub-sub-sub-encabezado 6</h6>
</body>
```
<p>Otros elementos para resaltar la semántica de los textos dentro de un documento HTML son:</p>
<p>strong: Resalta el texto del contenido en negrilla para dar la sensación de que esa información resaltada es importante y no debe pasarse por alto.</p>

```html
<body>
  <p>Este <strong>palabra</strong> es importante</p>
</body>
```
<p>b: Tiene el mismo efecto de strong. La diferencia radica cuando se usan lectores de pantalla, strong da importancia al texto mientras que b no.</p>

```html
<body>
  <p>Este <b>palabra</b> esta resaltada en negrilla</p>
</body>
```
<p>em: Este elemento convierte el estilo del texto en <i>cursiva</i> para marcar una semántica de énfasis en el texto.</p>

```html
<body>
  <p>Este <em>palabra</em> es importante</p>
</body>
```
<p>i: Tiene el mismo efecto de em. Se sugiere utilizar para enfatizar términos técnicos, extranjerismos y palabras en otro idioma, nombres propios, entre otros.</p>

```html
<body>
  <p>Estoy aprendiendo <i>machine learning</i> en la universidad</p>
</body>
```
<p>mark: Resalta, si con color como un resaltador de tinta, una porción de texto.</p>

```html
<body>
  <p>Este <mark>palabra</mark> esta resaltada con color</p>
</body>
```
<p>cite: Tiene el mismo efecto de em al poner el texto en cursiva. Se sugiere utilizar para enfatizar el título de un trabajo artístico o creativo como un libro, un poema, un cuadro, una escultura, una película, una canción, etc.</p>

```html
<body>
  <p><cite>The Lord of the Rings</cite> por J.R.R. Tolkien</p>
</body>
```
<p>dfn: Se utiliza para indicar que el contenido es una definición de un término incluido en el contenido del texto. El efecto en el texto es el de cambiar el estilo en cursiva.</p>

```html
<body>
  <p><dfn>HTML</dfn> es el lenguaje estandar para crear páginas web.</p>
</body>
```
<p>q: Este elemento permite definir una citación corta, el navegador agrega doble comillas para delimitar el texto citado.</p>

```html
<body>
  <p>Martin Luther King Jr. dijo: <q>Yo tuve un sueño</q></p>
</body>
```
<p>blockquote: Se usa para incluir citas textuales de texto obtenido de una fuente externa. Se utiliza el atributo <i>cite</i> para referenciar la fuente original del texto. El efecto en el texto es que la cita se ubica en un nuevo párrafo con un indentado que la diferencia del texto regular.</p>

```html
<body>
  <blockquote cite="http://www.worldwildlife.org/who/index.html">
    For 50 years, WWF has been protecting the future of nature. The world's leading conservation organization, WWF works in 100 countries and is supported by 1.2 million 
    members in the United States and close to 5 million globally.
  </blockquote>
</body>
```
<p>s: Con esta etiqueta se indica que un texto ya no es correcto o relevante y está propenso a ser editado para su corrección. El efecto en el texto es el de cruzar una línea horizontal en el texto.</p>

```html
<body>
  <p>Esta <s>palbra</s> debe ser corregida</p>
</body>
```
<p>u: Se utiliza para subrayar un texto.</p>

```html
<body>
  <p>Esta <u>palabra</u> está subrayada</p>
</body>
```
<p>br: Es una etiqueta vacía que permite insertar un salto de línea en el texto.</p>

```html
<body>
  <p>Un salto de línea <br> otro <br> y otro más <br></p>
</body>
```
<p>hr: Permite separar secciones independientes o no relacionadas de un texto. Por ejemplo, un cambio de tema entre dos párrafos consecutivos. El efecto visual es el de una línea horizontal que separa los párrafos.</p>

```html
<body>
  <p>Aquí se está hablando de un tema</p>
  <hr>
  <p>Pero aquí de otro completamente distinto</p>
</body>
```
<p>sub: Este elemento permite insertar subíndices en el texto.</p>

```html
<body>
  <p>La formula química del agua es H<sub>2</sub>O</p>
</body>
```
<p>sup: Y este permite insertar superíndices.</p>

```html
<body>
  <p>2<sup>10</sup>=1024</p>
</body>
```
<p>code: Este elemento indica que su contenido hace parte de un código escrito con un lenguaje de programación. El efecto visual es el de un cambio en la fuente para distinguir el código del texto regular.</p>

```html
<body>
  <p>La etiqueta HTML <code>button</code> define un botón en el que puede hacerse clic.</p>
</body>
```
<p>Listas</p>
<p>HTML tiene elementos para construir listas en las que no importa el orden y también listas numeradas en las que el orden es relevante.</p>

<p>El elemento <code>&lt;ul&gt;</code> (unordered list) en combinación con el elemento <code>&lt;li&gt;</code> (list item) sirve para definir listas sin orden cuyos elementos son señalizados por viñetas o puntos por el navegador.</p>

```html
<body>
  <ul>
    <li>Elemento de la lista no ordenada</li>
    <li>Elemento de la lista no ordenada</li>
    <li>Elemento de la lista no ordenada</li>
  </ul>
</body>
```  
<p>Por su parte las listas en las que el orden de los elementos de la lista es importante y por tanto un número que indica el orden se construyen con una combinación de los elementos <code>&lt;ol&gt;</code> (ordered list) y <code>&lt;li&gt;</code> (list item).</p>

```html
<body>
  <ol>
    <li>Elemento 1 de la lista ordenada</li>
    <li>Elemento 2 de la lista ordenada</li>
    <li>Elemento 3 de la lista ordenada</li>
  </ol>
</body>
```  



