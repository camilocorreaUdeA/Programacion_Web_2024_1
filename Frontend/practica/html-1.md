# Manos a la obra

A poner en práctica lo que hemos aprendido de HTML hasta este momento:

<ul>
  <li>Elementos HTML: etiqueta de apertura + contenido + etiqueta de cierre</li>
  <li>Estructura básica de un documento HTML: html{{head{}}+{body{}}}</li>
  <li>Presentación y formato de textos en HTML: parráfos, encabezados, listas, elementos de formato, hipervínculos, etc.</li>
  <li>HTML semántico: imágenes, video, contenido incrustado, semánticas de estructura del documento (main, header, nav, aside, etc...)</li>
</ul>

## Ejercicio propuesto:

Una agencia de turismo está evaluando contratar con usted el desarrollo de su sitio web. Le han solicitado como primer paso en la validación del negocio que proporcione una prueba de concepto (PoC) del sitio web.
El PoC del sitio web que la agencia quiere debe cumplir con los siguientes requisitos:

1. El sitio web debe tener una página principal (<code>index.html</code>) que tiene un menú de navegación con enlaces a 6 destinos turísticos alrededor del mundo, tiene un encabezado con el nombre de la agencia: "Travel Partners", una imagen relacionada con la industria del turismo, una sección "Acerca de nosotros" donde en una corta descripción se presenta la agencia al público, una sección donde se muestran los 6 destinos disponibles y se proporciona enlace a la página de cada destino por medio de una imagen o "thumbnail". Por último, esta página debe tener un "footer" donde deben aparecer datos del desarrollador del sitio web (nombre y correo electrónico).
2. Cada destino turístico debe tener su página individual con un encabezado con el nombre del lugar, una imagen principal, una sección con un texto descriptivo del destino (por ejemplo, las pirámides de Egipto), una sección que proporciona una lista de los atractivos para visitar en el lugar, una sección con un corto itinerario de ejemplo de las actividades que se pueden realizar en 5 días y 4 noches en el destino, y por último una sección donde se explica los costos de viaje, traslados, hospedajes, entradas a sitios turísticos, etc.
3. Agregue un título y un ícono para la pestaña del navegador (aquí puede encontrar muchos: https://icons8.com/icons/set/favicon). Incluya todos los metadatos que considere pertinentes para dar información del sitio web.
4. Utilice todos los recursos que tiene a disposición, imágenes, videos, enlaces a sitios externos, contenido incrustado, etc.
5. Items a tener en cuenta para el desarrollo de las páginas:
<ul>
  <li>Se deben construir utilizando HTML para su estructura, es decir, deben contar con <code>header</code>, <code>nav</code>, <code>main</code>, <code>section</code>, <code>article</code>, <code>footer</code> y <code>aside</code> (este último solo donde aplique o sea necesario)</li>
  <li>En el <code>header</code> debe haber un encabezado con que identifique a la cocina y una imagen del país a la que pertenece.</li>
  <li>En el <code>nav</code> se deben incluir un menú de navegación a todas las páginas que componen el sitio web</li>
  <li>En cada <code>main</code> debe tener un encabezado que permita identificar el tema que se trata en la página (si es un platillo o si es la página de ingredientes)</li>
  <li>En el <code>main</code> de la página principal debe tener dos <code>section</code> uno para hacer una introducción a los platillos (cada uno en su propio <code>article</code> y con un enlace a su respectiva página) y en el otro una introducción a los ingredientes con enlace a esa página</li>
  <li>En el <code>main</code> de las páginas de los platillos incluya <code>section</code> para la descripción del platillo, otro para sus ingredientes y modo de preparación, no olvide incluir una imagen del platillo ;-)</li>
  <li>En el <code>main</code> de la página de ingredientes haga una <code>section</code> para cada ingrediente quee incluya <code>article</code> para una corta descripción del ingrediente, otro para una lista de platillos en los que se puede encontrar ese ingrediente. Incluya imagenes del ingrediente</li>
  <li>En el <code>footer</code> solo agregue una frase como <i>Sitio web de comida xyz desarrollado por tu_nombre. Contacto: tu_correo_electrónico</i></li>  
</ul>

Otros aspectos que se deben validar:
<ul>
<li>Todas las páginas comparten el mismo <code>header</code>, <code>nav</code> y <code>footer</code></li>
<li>Apuestele a su creatividad y aproveche la mayor cantidad de resursos que pueda incluir en sus páginas (imágenes, videos, gifs animados, enlaces a contenido externo, etc.)</li>
<li>No se preocupe por el aspecto del sitio, por el momento solo nos interesa la funcionalidad y la estructura del mismo, más adelante cuando estudiemos las hojas de estilo con <i>CSS</i> tendremos oprtunidad de mejorar la parte visual</li>
</ul>

6. Cree un repositorio en su cuenta github para este ejercicio y nombrelo <i>travel_partners_agency</i>. Todos los archivos que conforman el sitio web deben ser guardados allí.
7. Publique su sitio web: Github te permite publicar tu sitio web para que pueda ser visitado. Sigue los pasos a continuación para que tu sitio quede publicado en una URL así: &lt;tu-usuario-de-github&gt;.github.io/travel_partners_agency (Ejem: https://camilocorreaudea.github.io/repo_remoto_ssh). Siga los pasos a continuación:
<p>&#8627; Asegurese de que su repositorio si tiene un archivo que se llama <code>index.html</code> sino entonces debe tener uno con ese nombre (el de página principal)</p>
<p>&#8627; Vaya a la opción <i>Settings de su repositorio</i> (la opción Settings con el ícono de la rueda dentada)</p> 

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/7a41b02c-b674-499e-8b01-8b69d0b19d6f)

<p>&#8627; De clic en <code>Pages</code> en el menú de opciones de la izquierda</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/15b6de10-f045-45a4-929b-811cb6eb31cc)

<p>&#8627; En la opción <code>Branch</code> cambie la opción <code>None</code> por <code>main</code></p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/5095e78e-7e8c-405d-a1a1-d92d2b69156e)

## Entrega: Responder al correo de entrega (que se va a enviar en el transcurso de la semana) con el enlace al sitio web (URL de Github pages)

## Ejemplo

El resultado final debería ser algo similar a lo que se observa en las siguientes imágenes

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2024_1/assets/42076547/30f4c9ee-32b2-4401-a086-a83b5c844f6a)

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2024_1/assets/42076547/ae74ae21-dae3-44c5-9ce4-aed35bcf2834)




