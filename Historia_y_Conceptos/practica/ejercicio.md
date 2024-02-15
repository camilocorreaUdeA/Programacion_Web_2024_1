# A ponernos finos con git y github

1. Descargue e instale git
2. Descargue e instale VSCode o el IDE de su preferencia
3. Abra una cuenta en github (ignore este paso si ya tiene una)
4. Asocie una llave criptografica para autenticación por SSH a su cuenta de github
5. Crea un nuevo repositorio remoto llamado <i>my_web_page</i>
6. Agregue un archivo <i>README.md</i> (Use luego este archivo para consignar todos los comandos de git que utilice en el desarrollo de este ejercicio)
7. Clone el repositorio en su máquina para tener una copia local
8. Abra una nueva rama de trabajo en el repositorio llamada <i>new_features</i>
9. Con ayuda del IDE instalado en su máquina cree un archivo con nombre <i>aboutme.html</i> y agregue allí el código disponible [aquí](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/blob/main/Historia_y_Conceptos/practica/aboutme_html.txt)
10. Haga commit del cambio que acaba de realizar, agregue un mensaje que describa claramente lo que se hizo.
11. De nuevo utilice el IDE para agregar un directorio al proyecto, con el nombre <i>styles</i>, dentro de este directorio agregue un archivo de nombre <i>aboutme.css</i> con el código disponible [aquí](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/blob/main/Historia_y_Conceptos/practica/aboutme_css.txt)
12. Haga commit del cambio que acaba de realizar, agregue un mensaje que describa claramente lo que acaba de hacer.
13. En el archivo <i>aboutme.html</i> modifique el contenido de los textos de las líneas 12, 16, 17, 21, 26, 28, 30, 32 con información relevante según sea el caso.
14. Haga commit de esos cambios que acaba de realizar, agregue un mensaje que describa lo que hizo.
15. Agregue un nuevo archivo <i>index.html</i> con el contenido disponible [aquí](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/blob/main/Historia_y_Conceptos/practica/index_html.txt)
16. Haga commit del paso anterior
17. Reemplace el contenido de la línea 7 con lo siguiente (en el archivo aboutme.html):
    ```html
     <link rel="stylesheet" href="styles/aboutme.css">
    ```
19. Luego de la línea 35 agregue las siguientes 5 líneas (en el archivo aboutme.html)
    ```html
     <ul class="social">
        <li><a class="css-is-deranged" href="_TU_DIR_DE_GITHUB_">GitHub</a></li>
        <li><a class="css-is-deranged" href="_TU_DIR_DE_TWITTER_">Twitter</a></li>
        <li><a class="css-is-deranged" href="_TU_DIR_DE_INSTAGRAM_">Instagram</a></li>
    </ul>
    ```
20. Haga un commit para consignar estos dos últimos cambios
21. Es hora de enviar nuestros cambios al repositorio remoto, pero antes asegúrese de descartar el commit con el cambio introducido en el punto 15.
22. Ahora si haga push de su rama main/master con sus cambios al remoto (no olvide hacer antes un merge en main/master de su rama new_features)
