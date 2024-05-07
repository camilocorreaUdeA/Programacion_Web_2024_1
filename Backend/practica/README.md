![cover](assets/cover.png)

<div align="center">
	<h2>Superhero App</h2>
</div>

> [Live Preview](https://rapidapi-example-superhero-app.vercel.app/)

### 🛠️ Installation Steps

1. Clona los archivos de este repositorio

2. Abra la carpeta donde se alojaron los archivos

3. Instale las dependencias

```bash
npm install
```

4. Crea una aplicación de backend que exponga un endpoint para obtener los datos de algunos superhéroes.

a. El servidor debe exponer un único endpoint <code>/api/superhero</code> que solo recibe peticiones de tipo <code>GET</code> en el puerto 8080 de localhost<br>
b. El nombre del súperheroe se pasa en la solicitud en un <code>query parameter</code> denominado <code>hero</code> (/api/superhero?hero={nombre})<br>
c. Tenga disponibles datos en un almacenamiento en memoria para devolver información de por lo menos 5 súperheroes distintos.<br>
d. El formato del body de la respuesta de la API es el siguiente:

```json
{
    "name": "Wolverine",
    "biography": {
        "fullName": "John Logan"
    },
    "powerstats": {
        "intelligence": 63,
        "strength": 32,
        "speed": 50,
        "durability": 100,
        "power": 89,
        "combat": 100
    },
    "images": {
        "xs": "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg",
        "sm": "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg",
        "md": "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg",
        "lg": "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg"
    }
}
```

5. Ponga en ejecución la aplicación de backend

```bash
go run main.go
```
6. Ponga en ejecución esta app

```bash
npm run dev
```
En este punto debería estar listo para ver la app funcionando Abra el navegador en [localhost:3000](http://localhost:3000/) para que acceda a la interfaz de usuario.
