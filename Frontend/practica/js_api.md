# Una app para consultar tus recetas favoritas...

## The Meal DB API

The Meal DB es una API de libre acceso que contiene una nutrida base de datos de platos de la gastronomía a nivel mundial. Acceda a la documentación de la API para que se familiarice con ella: [The Meal DB](https://www.themealdb.com/api.php)

En pocas palabras la API tiene disponibles los siguientes endpoints que pueden ser consultados con peticiones GET desde un cliente de REST API:

<ul>
  <li><b>Buscar platos por nombre: </b>www.themealdb.com/api/json/v1/1/search.php?s=<code>nombre_plato</code></li>
  <li><b>Buscar platos por ingrediente principal: </b>www.themealdb.com/api/json/v1/1/filter.php?i=<code>ingrediente_principal</code></li>
  <li><b>Buscar platos por categoría: </b>www.themealdb.com/api/json/v1/1/filter.php?c=<code>categoria</code></li>
  <li><b>Buscar platos por área (nacionalidad): </b>www.themealdb.com/api/json/v1/1/filter.php?a=<code>area_nacionalidad</code></li>
</ul>

Como se puede ver en los endpoints la API permite filtrar los resultados por 4 criterios distintos, y para los cuales devuelve respuestas distintas también.

Criterios de filtrado:

<ul>
  <li><b>Nombre del plato: </b>Permite buscar por el nombre exacto del plato. Solo nombres en inglés desafortunadamente. EN la respuesta los datos más relevantes vienen en los campos <code>strInstructions</code>, <code>strMealThumb</code>, <code>strTags</code> y <code>strIngredient#</code> cuya información respectivamente es las instrucciones de la preparación del plato, la imagen del plato, etiquetas asociadas al plato y los ingredientes (máximo 20 por plato)</li>
  <li><b>Ingrediente principal: </b>Filtra los resultados por ingrediente, para buscar todos los platos en cuya preparación está presente ese ingrediente (nombre del ingrediente en inglés :-)). El campo más importante de la respuesta es <code>strMeal</code> que trae el nombre del plato para poder hacer consultas por nombre.</li>
  <li><b>Categoría: </b>Este filtro devuelve los platos pertenecientes a alguna de las 14 categorías disponibles en la API: Beef, Chicken, Dessert, Lamb, Miscellaneous, Pork, Seafood, Side, Starter, Vegan, Vegetarian, Breakfast, Goat. El campo más importante de la respuesta es <code>strMeal</code> que trae el nombre del plato para poder hacer consultas por nombre.</li>
  <li><b>Área (nacionalidad): </b>Filtra los platos de acuerdo al lugar de origen. No todas las cocinas ni países del mundo están disponibles. El campo más importante de la respuesta es <code>strMeal</code> que trae el nombre del plato para poder hacer consultas por nombre. </li>
</ul>

Otros endpoints útiles:
<ul>
<li>Lista de todas las categorías disponibles en la API: www.themealdb.com/api/json/v1/1/list.php?c=list</li>
<li>Lista de todos los ingredientes disponibles en la API: www.themealdb.com/api/json/v1/1/list.php?i=list</li>
<li>Lista de todas las áreas disponibles en la API: www.themealdb.com/api/json/v1/1/list.php?a=list</li>
</ul>

## ¿Qué hay que hacer pues?

### Implemente un cliente (frontend) que consulte a The Meals DB API

<b>1. Página principal de la app</b>

La pagína principal de la app solo debe tener como contenidos: un elemento input de tipo texto para recibir la entrada del susuario, un botón que debe iniciar la búsqueda al dar click y un menú desplegabale con las distintas opciones de búsqueda (nombre de plato, categoría, ingrediente principal o área).

El botón no debe tener texto sino una imagen de una lupa y debe estar ubicado al lado de la caja de texto del elemento input. Mientras que el menú desplegable de opciones de búsqueda debe estar ubicado debajo de los dos elementos anteriores.

Debe validar la información que ingresa el usuario (nombre de plato, categoría, ingrediente o área existan en The Meals API) y mostrar un mensaje para que el usuario repita la búsqueda con información válida.

<b>2. Búsquedas</b>

Si la búsqueda es por nombre del plato se debe cargar la información del plato (ver la implementación de información del plato).

Pero si la búsqueda es por alguno de los otros criterios entonces debe capturar el nombre del plato de la respuesta y luego hacer una nueva consulta por nombre de plato para poder cargar la información del plato con la respuesta de está última solicitud.

<b>3. Búsquedas por ingrediente, categoría y área.</b>

Cuando se hagan búsquedas por los criterios ingrediente, categoría y área se debe cargar una galería de tarjetas con los resultados. Cada tarjeta esta compuesta de la imagen del plato acompañada del nombre. Límite la cantidad de resultados en la galería a máximo 18 (en el orden que los recibe en la respuesta o escogiendolos de manera aleatoria).

Se debe poder dar click en cada tarjeta para consultar la información del plato (cargar la información del plato).

<b>4. Información del plato</b>
   
Información:

<ul>
  <li>Nombre del plato</li>
  <li>Imagen de referencia</li>
  <li>Ingredientes</li>
  <li>Receta (pasos preparación)</li>
</ul>

La información del plato de se debe presentar completa (los puntos en la lista anterior) y de una forma agradable para el usuario. Es libre el diseño que desee utilizar para presentar la información en la pantalla del navegador. Dele rienda suelta a su creatividad.

<b>5. Recapitulando el diseño</b>

<ol>
  <li>Página principal para realizar la búsqueda: Allí se hace la selección del criterio de filtrado y se ingresan los datos para la búsqueda.</li>
  <li>Galería de resultados: Solo se debe cargar cuando el criterio de búsqueda es por categría, ingrediente o área. Cada tarjeta debe mostrar la imagen y el nombre del plato.</li>
  <li>Información del plato: Se carga cuando se hace búsqueda directa por nombre del plato o cuando se da click en una tarjeta de la galería de resultado cargada al hacer la búsqueda por alguno de los otros criterios.</li>
</ol>

Nota: Siempre debe estar disponible un enlace a la página principal que permita retornar a hacer una nueva búsqueda.

<b>6. Requisitos y restricciones</b>

<ul>
  <li>El diseño de la interfaz de usuario es libre. Acuda a su creatividad para hacer una aplicación atractiva y que geenre una gran experiencia de usuario. Pero no olvide ninguno de los componentes indicados en el diseño.</li>
  <li>Para la implementación del diseño de la interfaz de usuario además de lo que hemos aprendido de HTML y CSS puede utilizar librerías de CSS externas y con las que se encuentre familiarizado (bootstrap, tailwind, etc.)</li>
  <li>NO está permitido el uso de frameworks web para Javascript (React, Next, Angular, etc.). Solo puede utilizar las funcionalidades que hemos visto para manipular el DOM con Javascript.</li>
  <li>Las solicitudes a los endpoints de The Meals DB api SOLO se deben hacer ya sea con axios o con fetch.</li>
  <li>Implementar la app en Github pages. Se penalizará fuertemente fallar en este requisito.</li>
  <li>Recuerde que a mayor detalle, mejor la nota de la actividad. Exprima su creatividad al máximo en la implementación de la apariencia gráfica de la app. Recuerde que la apariencia influye en gran medida en una experiencia de usuario de mayor calidad.</li>
</ul>

## Entrega: Responder al correo (y por favor solo a ese correo :-D) de entrega (que se va a enviar en el transcurso de la semana). Donde debe incluir los enlaces al sitio web desplegado en Github pages y al repositorio.

