async function getData(endpoint){
    try{
        const response = await axios.get(endpoint);
        const data = response.data;
        return data;
    } catch(error){
        console.error(`fallo la peticion: ${error}`);
    }
}

(async() => {
    try{
        const datos = await getData('https://reqres.in/api/users?page=2');
        const data = datos.data;
        const main = document.querySelector("main");
        const contenedor = document.createElement("div");
        contenedor.setAttribute("class", "container");
        for(dato of data){
            const tarjeta = document.createElement("div");
            tarjeta.setAttribute("class", "user");
            const imagen = document.createElement("div");
            imagen.setAttribute("class", "user-pic");
            const img = document.createElement("img");
            img.setAttribute("src", dato.avatar);
            const userName = document.createElement("div");
            userName.setAttribute("class", "user-name");
            const userEmail= document.createElement("div");
            userEmail.setAttribute("class", "user-email");
            userName.innerHTML = `<p>Nombre: ${dato.first_name} ${dato.last_name}</p>`
            userEmail.innerHTML = `<p>Correo: ${dato.email}</p>`
            imagen.appendChild(img);
            tarjeta.appendChild(imagen);
            tarjeta.appendChild(userName);
            tarjeta.appendChild(userEmail);
            contenedor.appendChild(tarjeta);
        }
        main.appendChild(contenedor);
    }catch(error){
        console.error(`fallo al obtener los datos: ${error}`);
    }
})();
