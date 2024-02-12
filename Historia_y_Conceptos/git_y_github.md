## ¿Qué es Git y cómo funciona?

<p>Git es el sistema de control de versiones más popular de la actualidad, y no solo es 
  utilizado por programadores de software sino también por personas que se desempeñan en 
  otras áreas y profesiones ya que puede usarse para versionar cualquier tipo de archivo 
  y no es exclusivo de archivos de código fuente de software.</p>

<p>¿Pero qué es un sistema de control de versiones (<i>VCS</i>)? Un sistema de control de 
  versiones rastrea y registra los cambios realizados en un archivo o grupo de archivos 
  permitiendo volver a versiones anteriores en cualquier momento que sea necesario.</p>

<p>Un <i>VCS</i> permite que muchas personas puedan trabajar de forma colaborativa en un 
  mismo proyecto de manera concurrente y desde distintas ubicaciones físicas.</p>

<p>Con un <i>VCS</i> distribuido como git todos los miembros de un equipo de trabajo pueden
  tener una copia local del proyecto y su historial de versiones en sus propios 
  computadores o dispositivos de trabajo. De esta manera no es necesario que estén 
  conectados en línea para agregar cambios en los archivos del proyecto.</p>

<p>Git modela la historia de una colección de archivos y directorios, que están dentro de 
un directorio general (la carpeta que almacena todo el proyecto), como una serie de 
instantáneas o revisiones en el tiempo y como un grafo dirigido y sin ciclos.</p>

<p>Cada instantánea está precedida de un conjunto de instantáneas anteriores a ella. No 
  puede pensarse la historia de un repositorio en git como un modelo lineal ya que una 
  instantánea puede descender de múltiples instantáneas que se desarrollaron en ramas 
  simultáneas y paralelas de trabajo que luego fueron mezcladas.</p>

<p>Estas instantáneas de las que venimos hablando tienen nombre propio en git y se 
  conocen como <strong><i>commits</i></strong>.</p>

<p>Cualquier componente básico de un repositorio, sea un archivo, un directorio o un 
  commit, es tratado como un objeto en git. Y cada objeto es referenciado mediante un 
  <strong><i>hash SHA-1</i></strong> de modo que en el modelo de datos de git los objetos no contienen 
  físicamente a otros objetos sino que más bien tienen referencias a ellos a través del 
  <strong><i>hash</i></strong>.</p>

<p>Una referencia en git es un puntero a un <strong><i>commit</i></strong>, 
por ejemplo la referencia <strong><i>master</i></strong> es un puntero al commit más reciente de la rama 
principal del repositorio. Entonces, a pesar de que los objetos en git son inmutables 
las referencias si lo son y pueden ser actualizadas en cualquier momento para que apunten a otros objetos.</p>

<p>Una de las referencias más importantes es la que se denomina <strong><i>HEAD</i></strong> 
  que apunta al último <strong><i>commit</i></strong> realizado sobre la rama actual de trabajo, es a grandes 
  rasgos una referencia que señala el lugar donde se encuentra uno ubicado en el momento 
  dentro del repositorio.</p>

<p>Una vez aclaradas las definiciones de objeto y referencia se puede pasar a denominar 
  de manera más exacta lo que es un repositorio en git.</p>

<p>Un repositorio es simplemente el conjunto de objetos y referencias asociadas a un 
  directorio de trabajo y que juntos forman el grafo dirigido sin ciclos. De este modo 
  todos los comandos de git son simplemente acciones que resultan en la adición de 
  nuevos objetos y/o en la adición o actualización de referencias dentro del 
  repositorio, lo que resulta en la manipulación directa de la estructura del 
  grafo subyacente.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/1f685280-0de2-4c30-a0e2-024b47827239)

# Los 3 estados en git y el “git workflow”

<p>Git tiene 3 estados en los cuales se puede encontrar cualquier archivo, esos estados 
son: <strong><i>modified</i></strong> (modificado), <strong><i>staged</i></strong> (en zona de staging) y <strong><i>commited</i></strong> (agregado a una 
instantánea).</p>

<p><strong><i>Modified</i></strong>, quiere decir que se han hecho cambios al archivo 
  pero esos cambios no se han agregado a ninguna instantánea.</p>

<p><strong><i>Staged</i></strong>, significa que los cambios se han marcado para ser 
  agregados a la próxima instantánea (commit).</p>

<p><strong><i>Commited</i></strong>, indica que los cambios se han almacenado con 
  éxito en una instantánea (commit) del repositorio.</p>

<p>Los 3 estados direccionan a las 3 secciones principales de un proyecto en git: 
  el <strong><i>working tree</i></strong>, el <strong><i>área de staging</i></strong> 
  y el <strong><i>directorio de git</i></strong> o estado actual del repositorio.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/3e6810ee-6ee8-4c4b-ba68-7cc3b17cfda1)

<p>El <strong><i>working tree</i></strong> es la sección en la que se modifican los archivos, ya que allí se tiene 
una copia de la versión en curso del repositorio.</p>

<p>En el <strong><i>área de staging</i></strong> se marcan los cambios que se desea hagan parte del próximo <strong><i>commit</i></strong>.</p>

<p>Cuando se hace un <strong><i>commit</i></strong> se toman los cambios que están en el área de staging y se 
almacenan en una nueva instantánea en el <strong><i>directorio de git</i></strong>.</p>

# Github

<p>Github es un servicio de almacenamiento en la nube que ayuda en la gestión ágil y 
  eficiente de repositorios de proyectos de software.</p>

<p>Github es una tecnología complementaria a git, pero no necesaria para trabajar con 
  git, mientras que github si depende de git.</p>

<p>Github permite que los desarrolladores de software puedan crear repositorios remotos 
  en la nube. Una vez se crea el repositorio remoto es posible obtener una copia local 
  de forma que se puedan hacer modificaciones a los archivos localmente y luego 
  “empujar” esos cambios de vuelta al repositorio remoto donde los cambios quedan 
  públicos y visibles a los demás miembros del equipo.</p>

<p>Aparte de tener una interfaz gráfica bastante intuitiva y amigable, una de las 
  mayores ventajas al usar github es la posibilidad de poder usarlo como una plataforma 
  colaborativa y de relacionamiento social con otros desarrolladores ya que los 
  usuarios de github tienen perfiles asociados a sus cuentas en los que pueden mostrar 
  sus proyectos, contribuciones y su actividad reciente.</p>

<p>Github como plataforma social alienta a los desarrolladores a explorar y contribuir 
  en proyectos open-source de cualquier tipo. Es bastante sencillo hacer un fork de un 
  proyecto, agregar cambios o nuevas funcionalidades y luego hacer un pull request de 
  modo que el dueño del proyecto pueda tomar una decisión no solo desde la revisión de 
  código sino basándose además en la información relevante que puede obtener al revisar 
  el perfil del colaborador (contribuciones pasadas, lenguajes en los que trabaja, 
  etc).</p>

<p>Para desarrolladores en búsqueda de empleo se convierte en una herramienta poderosa 
  ya que con github pueden mostrar sus proyectos a modo de portafolio profesional para 
  que los empleadores puedan tener de primera mano una validación de las habilidades 
  técnicas del aspirante.</p>

## Trabajando con git y github (ejemplo práctico)

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/04ffed62-b3d1-4193-aa58-29a8ed890596)

### Haciendo push desde un repositorio local de git a uno remoto en github

<p>Lo primero es crear un directorio para el repositorio en local y convertirlo en 
  un repositorio de git con el comando <strong><i>git init</i></strong></p>

<p>Para este ejemplo creamos un repositorio local en la carpeta repo_local, y con el 
  comando git init hacemos que el repositorio local se pueda versionar con git.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/de517f7a-df05-4181-849e-d90095384b69)

<p>Añadimos un nuevo archivo al repositorio local, para efectos de este ejemplo vamos a 
agregar un archivo vacío llamado index.html</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/6571d18f-a931-4de2-a2ca-34e9a6b64652)

<p>Verificamos el estado de git, es decir que el nuevo archivo haya sido añadido al 
working tree, con el comando <strong><i>git status</i></strong></p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/560679a2-4476-42d9-8d65-d5f17b949e76)

<p>Git indica que hay archivos nuevos a los cuales aún no se les hace seguimiento de 
  los cambios por tanto no pueden ser incluidos en ningún commit.</p>

<p>Añadimos el archivo al área staging para que git pueda tenerlo en cuenta en el 
  próximo commit. Esto se hace con el comando <strong><i>git add [nombre_archivo]</i></strong></p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/5472a153-fc7a-417d-a859-d69f7ea87051)

<p>Git indica que en staging hay un nuevo archivo que puede ser incluído en el 
  próximo commit.</p>

<p>Hacemos un commit para almacenar una instantánea o revisión del cambio, en este caso 
  en la rama principal del repositorio local. Damos un mensaje relevante para tener un control adecuado 
  de los cambios entre versiones. Lo anterior se logra con el comando <strong><i>git commit -m “[mensaje_del_commit]”</i></strong></p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/0b788485-e571-407c-b173-1fcda57a9cc7)

<p>Con el comando <strong><i>git log</i></strong> podemos revisar la historia de commits 
de una rama en particular, en este caso la rama master. <strong><i>git log</i></strong> indica el hash del 
commit, hacia donde apunta HEAD, el autor, la fecha y el mensaje descriptivo del 
commit.</p>

<p>Para subir los cambios a un repositorio remoto en github primero se debe crear un 
  nuevo repositorio en github.</p>

<p>Accedemos a nuestra cuenta de github, damos clic en el símbolo (+) que está en la 
  barra de navegación y luego clic en Nuevo Repositorio en el menú desplegado.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/a347d811-4060-486f-8773-82fba0d55e70)

<p>Le damos un nombre al repositorio remoto, y podemos dejar las siguientes opciones 
  tal cual como están o completarlas según nuestra conveniencia. Luego damos clic en 
  el botón Crear repositorio</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/65b70162-2ac0-4de9-904e-c39eb53a29a3)

<p>Una vez creado el repositorio remoto github ofrece varias opciones para empezar a 
  usarlo, para los fines de este ejemplo en particular nos interesa subir los archivos 
  y commits de nuestro repositorio local al remoto.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/297aa500-32c0-4270-a59f-238b1d50ef41)

<p>Copiamos los comandos que nos proporciona github y los ejecutamos en la terminal.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/4965dd6c-b540-433b-98fc-cae4b97ad944)

<p>Al parecer las cosas no ocurrieron como esperábamos. Esto tiene explicación, y es que 
no tenemos las credenciales de autenticación a github en nuestro repositorio local que 
nos permitan interactuar de manera segura con el repositorio remoto alojado en github.</p>

<p>Llegados a este punto tenemos 2 opciones:</p>

<ol>
  <li>Autenticarse usando un token de autenticación generado en github. Con dicho 
    token podemos interactuar con el repositorio utilizando peticiones del protocolo HTTP.</li>
  <li>Autenticarse usando llaves públicas y privadas, para interactuar con el repositorio 
    remoto a través del protocolo SSH.</li>
</ol>

<p>Para la primera opción debemos crear el token de autenticación en github. 
  Tutorial adicional en este enlace: https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens</p>

<p>Vamos a dar clic en el avatar de nuestra cuenta de github, ubicado en la esquina 
  superior derecha de la barra de navegación. Y luego clic en la opción Settings 
  (configuración) del menú desplegado.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/165d5bfc-cd8b-43bb-afb0-aabb5f5ffeb6)

<p>Da clic en la opción <strong>Developer settings</strong> (configuraciones para 
  desarrolladores), que está al final del menú en la izquierda.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/179ba832-7d5f-4e7f-9358-e19d25b11bbc)

<p>Una vez allí escoja la opción <strong>Personal acces tokens</strong> y luego la opción 
  <strong>Tokens (classic)</strong></p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/c35381e6-73a6-494e-85df-0c5bb0e8c0b6)

<p>Luego clic en <strong>Generate new token</strong>, y después en 
  <strong>Generate new token (classic)</strong></p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/0c718235-a510-4b6c-a705-caba41f74e35)

<p>Agregue una nota descriptiva y un tiempo de vencimiento al token. También asegúrese de marcar la opción repo en los scopes (alcances) del token.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/5f01e040-e002-4699-a92f-9e3beb582016)

<p>Luego clic en generar token que se encuentra al final</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/02251622-f91b-484c-a02c-185dce262241)

<p>Copie el token, y de ser posible guárdelo en un archivo de texto en su computador 
  porque si lo olvida github nunca le permitirá verlo de nuevo.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/3962f792-40ef-4ac6-9eb2-24671df7e39a)

<p>De vuelta en su terminal de línea de comandos ejecute el comando <strong><i>git remote add o
rigin https://[usuario-git]:[el-token-recien-creado]@github.com/[usuario-git]/[nombre-repositorio].git</i></strong>
</p>

<p>De nuevo <strong><i>git branch -M main</i></strong> y <strong><i>git push -u origin main</i></strong></p>

<p>Si recibe un error indicando que el remote origin ya existe, entonces ejecute el 
  siguiente comando: <strong><i>git remote rm origin</i></strong> y vuelva a ejecutar los 3 anteriores de nuevo</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/d1f86c41-80e6-481e-9ffb-d46aafc2d07f)

<p>Vaya al repositorio remoto en github y verifique que los cambios realizados en 
  local ahora están también en el remoto.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/b35eebdf-9736-4beb-81d7-9650389254c4)

<p>Para la segunda opción debemos crear un par de llaves pública y privada con SSH y 
  luego asociarlas a nuestra cuenta de github.</p>

<p>Los comandos que se muestran a continuación solo funcionan en sistemas operativos 
  Linux y macOS. Para ejecutarlos en Windows debe hacer uso de un terminal que 
  permita ejecutar comandos de linux (puede utilizar el git bash que fue incluido en 
  su instalación de git) Otra alternativa en Windows es utilizar el software Putty 
  (Puede ver un tutorial aquí: https://www.ibm.com/docs/es/flashsystem-5x00/8.5.x?topic=host-generating-ssh-key-pair-using-putty)</p>

<p>Ejecute el comando <strong><i>ssh-keygen -t rsa -b 4096 -C "[usuario/email-asociado-a-github]"</i></strong>. 
  Agregue un nombre de archivo de su preferencia para las llaves y sino de enter, 
  si desea agregar un password (passphrase) adicional a la llave agréguelo y repítalo 
  de nuevo, sino de enter y luego enter de nuevo.</p>

<p>En los siguientes enlaces puede ver tutoriales para la creación de llaves SSH y cómo asociarlas a su cuenta de github:
https://www.atlassian.com/git/tutorials/git-ssh
https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/f57de25c-1a0d-490b-ad58-e55b2115fccc)

<p>Ejecute el comando <strong><i>ssh-add [ruta-archivo-llave]</i></strong> para agregar las nuevas llaves al agente ssh de su máquina</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/d66c342b-6e97-4ea2-9496-656c0c222b96)

<p>Copie el contenido de la clave pública que está en el archivo con extensión .pub</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/550ad4e9-559a-4903-979f-4534ee0f896d)

<p>Ahora vaya a la opción settings en su cuenta de github y allí de clic en la opción SSH and GPG keys</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/b0ff8deb-0314-41ed-9244-66029b27cc69)

<p>Y luego clic en New SSH key</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/5798b377-0b55-41bb-8743-35513f9b26e8)

<p>Allí le damos un nombre a la llave y pegamos la llave pública que habíamos copiado 
  en la terminal en el espacio para Key, dejamos el tipo de llave en autenticación y 
  luego clic en Add SSH key</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/cfdf3e71-ea2c-4d07-9d4f-76fe1c3cf852)

<p>Luego ejecute los siguientes comandos en la terminal</p>

<p><strong><i>git remote add origin git@github.com:[usuario-git]/[nombre-repositorio].git</i></strong></p>
<p><strong><i>git branch -M main</i></strong></p>
<p><strong><i>git push origin main</i></strong></p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/38532789-ef3e-4de8-b366-b83d85d729ff)

<p>Vaya al repositorio remoto en github y verifique que los cambios realizados en 
  local ahora están también en el remoto.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/1c485aaa-d5e4-4143-b87e-4e045d9c05ac)

### Clonando una copia local de un repostiorio remoto alojado en github

<p>Usted también puede crear primero el repositorio en github para luego clonar una 
  copia en su computador.</p>

<p>Luego de crear el repositorio (cómo ya se vió anteriormente) y dependiendo 
  del método de autenticación que usted haya elegido para trabajar con github va a 
  tener que utilizar uno de los siguientes comandos en su terminal.</p>

<p>Si ha elegido autenticación con token:</p>

<p><strong><i>git clone https://[usuario]:[token]@github.com/[usuario]/[nombre-repositorio].git</i></strong></p>

<p>En usuario va su nombre de usuario de github y en token va el que usted generó 
  para la autenticación</p>

<p>Si ha elegido autenticación con ssh:</p>

<p><strong><i>git clone git@github.com:[usuario]/[nombre-repositorio].git</i></strong></p>

<p>En usuario va su nombre de usuario de github que debe estar asociado a una llave 
  ssh en github y en su computador.</p>

<p>Los enlaces anteriores para clonar el repositorio pueden obtenerse en la página 
  principal del repositorio remoto en github, dando clic en el botón Code</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/48a68f88-a5a2-4407-b22f-76691fd44744)

## Flujo de trabajo con git y github

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/b081827a-c7bc-4056-8c25-9ce0c3c052ea)

<p>Aquí listamos algunas de las principales ventajas de utilizar git y github en 
  proyectos de software:</p>

<ol>
  <li>Tener un sistema de control de versiones del repositorio, que no necesariamente 
    tiene que ser de un proyecto de software</li>
  <li>Permitir que varias personas al tiempo puedan trabajar y subir cambios al mismo 
    repositorio remoto (cada uno clona el repositorio remoto y hace cambios de manera 
    local) de modo que el repositorio se pueda ir actualizando paulatinamente con los 
    cambios que las personas suben desde su copia local al repositorio remoto</li>
  <li>Permitir el desarrollo de ramas de trabajo en un repositorio, de este modo los 
    cambios se pueden realizar sobre una copia de la versión estable del repositorio 
    pudiendo ser luego descartados para mantener la versión actual repositorio, o 
    bien, ser aprobados para agregarse a la rama principal del repositorio en una 
    nueva versión estable del repositorio</li>
  <li>Github proporciona una metodología de revisión de cambios mediante la cual otros 
    miembros del proyecto pueden participar en la revisión de los cambios que se 
    pretenden combinar con la rama principal del repositorio para obtener una nueva 
    versión estable</li>
</ol>

<p>Hagamos un recuento de los comandos de git que hemos visto hasta el momento:</p>

<p><strong>git init</strong>: Se utiliza para indicar a git que el directorio actual 
  va a estar sujeto a un control de versiones con git.</p>

<p><strong>git clone [string_conexión_repositorio_remoto]</strong>: Para clonar un 
  repositorio remoto alojado en github y tener una copia local para trabajar. 
  Vimos cómo clonar con autenticación por  HTTP y SSH.</p>

<p><strong>git status</strong>: Sirve para verificar el estado de la rama en la que 
  se está trabajando en el momento. Permite ver qué archivos están por fuera del 
  staging, cuáles ya hacen parte del mismo y cuáles ya están agregados a un commit.</p>

<p><strong>git add [nombre_archivo]</strong>: Este comando permite agregar a la zona 
  de staging un archivo nuevo o un archivo que ha sido modificado para ser tenido en 
  cuenta en el próximo commit del repositorio.</p>

<p><strong>git add .</strong>: Igual que el comando anterior pero agrega todos los 
  archivos nuevos y/o modificados a la zona de staging.</p>

<p><strong>git commit -m “[mensaje_descriptivo]”</strong>: Con este comando se 
  persisten o se almacenan en la copia local del repositorio los cambios y archivos 
  que estaban en la zona de staging. Cada commit es una nueva instantánea o revisión 
  de una rama del repositorio.</p>

<p>Para trabajar en proyectos de software con repositorios en git y github existe una 
  metodología conocida como el flujo de trabajo en git (git workflow) que consiste en 
  una serie de pasos en los cuales con la ayuda de comandos de git se facilita y 
  agiliza el trabajo colaborativo y distribuido en el proyecto.</p>

<p>La metodología es bastante simple ya que los comandos de git utilizados son pocos 
  y sencillos y además se repiten durante todo el ciclo de vida del proyecto de 
  software.</p>

<p>En este flujo de trabajo se pretende que en primera instancia el desarrollador 
  tenga almacenada y una copia local actualizada del repositorio remoto. Una vez vaya 
  a trabajar en cambios o en agregar nuevas funcionalidades debe crear una nueva rama 
  derivada de aquella que contiene la versión estable sobre la que quiere trabajar 
  (por lo general es la rama main/master del repositorio).</p>

<p>Una vez los cambios o adiciones han sido efectuados sobre la nueva rama (los 
  cambios pasaron por el staging y fueron agregados en un commit), esta pasa a ser 
  probada y revisada por otros integrantes del grupo de trabajo (otros 
  desarrolladores).</p>

<p>Y una vez el proceso de revisión da el visto bueno para que los cambios puedan 
  mezclarse a la rama principal del proyecto se procede a crear una nueva versión 
  estable que contiene los cambios recién realizados.</p>

<p>Se debe actualizar el repositorio para obtener los cambios que otros 
  desarrolladores hayan incluido en la rama principal del repositorio.</p>

<p>Flujo de trabajo con git (git workflow)</p>

<ol>
  <li>
    <p>Clona o almacena en tu computador una copia local del repositorio remoto 
      alojado en github</p>
    <p><strong><i>git clone git@github.com:[usuario]/[nombre-repositorio].git</i></strong></p>
  </li>
  <li>
    <p>Si ya tenías una copia del repositorio entonces solo necesitas actualizarlo 
      para que baje los cambios que hayan sido incluidos por otros desarrolladores</p>
    <p><strong><i>git pull</i></strong></p>
  </li>
  <li>
    <p>Crea una nueva rama para trabajar en tus cambios sin necesidad de afectar la 
      versión estable del repositorio</p>
    <p><strong><i>git checkout -b [nombre_nueva_rama]</i></strong></p>
  </li>
  <li>
    <p>Trabaja en tus cambios y una vez hayas finalizado agrega al staging los cambios 
      que quieras tener en cuenta para el commit y efectúa el mismo</p>
    <p><strong><i>git add [nombre_archivo]  /  git add .</i></strong></p>
    <p><strong><i>git commit -m “[mensaje_del_commit]”</i></strong></p>
  </li>
  <li>
    <p>A partir de este paso tienes dos opciones que dependen de si en el proyecto 
      estás trabajando solo o con un grupo de desarrolladores</p>
  </li>
</ol>

<p>Si estás trabajando solo basta con que te asegures de que tus cambios están bien y 
  que la funcionalidad es exactamente la esperada. Luego de eso puedes pasarte a la 
  rama principal del repositorio y mezclar los cambios de la rama de trabajo. Y 
  posteriormente subir la nueva versión a tu repositorio remoto para tenerlo 
  actualizado.</p>
<p><strong><i>git checkout [nombre_rama_principal_main_o_master]</i></strong></p>
<p><strong><i>git merge [nombre_rama_trabajo]</i></strong></p>
<p><strong><i>git push</i></strong></p>

<p>Si estás trabajando con otros compañeros encargados de revisar tus cambios debes 
  subir tu rama de trabajo al repositorio remoto, elevar un <strong>pull request</strong> en github, 
  escoger la lista tentativa de revisores y esperar por el resultado de la revisión 
  que bien puede ser sugerencias de correcciones y cambios adicionales o la aprobación 
  de los mismos.</p>

<p>Una vez aprobados se pueden mezclar con la rama principal del repositorio allí 
  mismo en github. Luego solo bastará actualizar tu copia local a la nueva versión 
  con tus cambios que ya está en el repositorio remoto.</p>

<p><strong><i>git push --set-upstream origin [nombre_rama_trabajo]</i></strong></p>
<p><strong><i>git checkout [nombre_rama_principal_main_o_master]</i></strong></p>
<p><strong><i>git pull</i></strong></p>

<p>A continuación se deja una lista con otros comandos de git que podrían ser de 
  utilidad en situaciones específicas dentro del flujo de trabajo:</p>

<p><strong><i>git diff [nombre_archivo]</i></strong>: permite revisar las 
  modificaciones que se han hecho a un archivo.</p>

<p><strong><i>git restore [nombre_archivo]</i></strong>: descarta los cambios hechos 
  al archivo y lo restaura a su estado original antes de efectuados los cambios.</p>

<p><strong><i>git branch</i></strong>: ver la lista de ramas existentes en el 
  repositorio.</p>

<p><strong><i>git branch -D [nombre_rama]</i></strong>: eliminar una rama del repositorio.</p>

<p><strong><i>git branch -m [nombre_actual_rama] [nombre_nuevo_rama]</i></strong>: 
  permite cambiar el nombre a una rama.</p>

<p><strong><i>git fetch</i></strong>: permite bajar todos los cambios de una rama de 
  otro repositorio (el remoto por ejemplo) sin necesariamente mezclarlos al 
  repositorio actual.</p>

<p><strong><i>git log</i></strong>: se usa para revisar el historial de cambios de 
  un repositorio. Proporciona una lista con los commits que se han mezclado a la 
  rama durante su ciclo de vida.</p>

<p><strong><i>git log --all --graph --decorate</i></strong>: ver el historial de 
  manera gráfica (como un grafo direccionado sin ciclos).</p>

<p><strong><i>git reset HEAD [nombre_archivo]</i></strong>: este comando es útil cuando se 
  necesita sacar del staging un archivo sin perder los cambios realizados en el.</p>

<p><strong><i>git clean -i</i></strong>: con ayuda de una interfaz interactiva 
  permite restaurar (o eliminar en caso de archivos nuevos) todos los archivos en 
  los que se han hecho modificaciones pero que no se han tenido en cuenta para ser 
  llevados a staging.</p>

<p><strong><i>git commit --amend -m “[nuevo_mensaje_commit]”</i></strong>: permite corregir o 
  cambiar el mensaje descriptivo del commit.</p>

<p><strong><i>git commit --amend --no-edit</i></strong>: permite agregar o eliminar 
  archivos de un commit, no cambia el mensaje descriptivo del commit.</p>

<p><strong><i>git revert HEAD~[orden_del_commit_desde_HEAD]</i></strong>: permite 
  eliminar los cambios introducidos en un commit. Es útil cuando se está trabajando 
  sobre una rama y se van agregando commits conforme se hacen cambios, pero si en 
  algún momento se desea descartar los cambios introducidos en algún commit basta 
  con saber a cuántos commits del puntero HEAD está ese commit y listo, o bien se 
  puede usar directamente el SHA del commit <strong><i>git revert [SHA_commit]</i></strong>.</p>

<p><strong><i>git config user.name [nombre_usuario_github]</i></strong>: permite 
  configurar la identidad de quien está trabajando en el repositorio local. Puede 
  ser de forma global para todos los repositorios agregando la bandera --global. 
  Debe utilizarse en conjunto con el comando a continuación.</p>

<p><strong><i>git config user.email [email_registrado_github]</i></strong>: permite 
  configurar la identidad de quien está trabajando en el repositorio local. Puede 
  ser de forma global para todos los repositorios agregando la bandera --global. 
  Debe utilizarse en conjunto con el comando anterior.</p>

## Si quieres aprender más visita estos enlaces:

<p>https://product.hubspot.com/blog/git-and-github-tutorial-for-beginners</p>
<p>https://git-scm.com/docs</p>
<p>https://www.freecodecamp.org/news/git-and-github-for-beginners/</p>
<p>https://www.datacamp.com/blog/all-about-git</p>
<p>https://blog.hubspot.com/website/what-is-github-used-for</p>
<p>https://phoenixnap.com/kb/how-to-use-git</p>
<p>https://developer.ibm.com/tutorials/d-learn-workings-git/</p>
<p>https://www.kodeco.com/books/advanced-git/v1.0/chapters/1-how-does-git-actually-work</p>
<p>https://luisiblogdeinformatica.com/git-github-y-vcs/</p>
<p>https://experienceleague.adobe.com/docs/contributor/contributor-guide/setup/github-signup.html?lang=es</p>
<p>https://git-scm.com/book/en/v2/Getting-Started-What-is-Git%3F</p>
<p>https://missing.csail.mit.edu/2020/version-control/</p>
<p>https://eagain.net/articles/git-for-computer-scientists/</p>
<p>https://xosh.org/explain-git-in-simple-words/</p>
