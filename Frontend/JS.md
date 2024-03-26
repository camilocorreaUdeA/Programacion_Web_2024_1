# Javascript: lo básico

## Variables y tipos de dato

### Declaración de variables

Para declarar variables se puede usar: <code>let</code>, <code>var</code> y <code>const</code>

La diferencia entre <code>let</code> y <code>var</code>: <code>var</code> permite re-declarar variables, mientras que con <code>let</code> ocurre un <code>SyntaxError</code>. Se prefiere el uso de <code>let</code> ya que <code>var</code> se usaba en las versiones más antiguas de Javascript.

Para declarar variables de solo lectura se utiliza <code>const</code>. Se recomienda el uso de mayúsculas (<i>UPPERCASE</i>) para las variables calificadas con <code>const</code>.

```js
const MY_BIRHTDAY = '29.02.1996'
```

Los nombres de las variables pueden empezar por letra, _ o $, nunca por número y no pueden tener caracteres especiales. Los nombres de variables son sensibles a mayúsculas y minúsculas (case-sensitive).

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2Cy0ktUUjPyU9KzAlLLMpMTMpJVbBVUIKIKJRBhZSsubjSSvOSSzLz8xRyK91K8zQ0q7kUFBQUclJLFHLyk1F1gwWQNYNUFqWWlBbloSq25qrl4krOzyvOz0nVy8lP10B1iaY1iiTUYk1rAA%3D%3D)

```js
let globalVariable = "global variable";

function myFun(){
    let localVariable = "local variable";
    return localVariable;
}

console.log(globalVariable);
console.log(myFun());
```

Las variables de tipo string se pueden inicializar con doble comilla o comilla simple. Para pasar valores de variables o expresiones a un <code>string</code> se debe crear el <code>string</code> con tildes y pasar la variable usando esta sintaxis: <code>${variable}</code>.

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CXYwxDsIwEAR7v2ITIYU0FGkt9%2FwgbZC9BNDFJ9lHgVD%2BHqFQpZ2dHaGhWkFAe6WIYtQiqWm9k30ZENCNlKgLYQp7EK93noVN591uMRtzJAKm0%2Ff3WbF8kHgruJcnc5q8c1FzVeFFdD5XK%2F0RDAfyj%2FYb)

```js
let str = "Hello World!";
let str2 = 'Welcome to the jungle!';

let sentence = `${str2} my dear friend`;

console.log(str)
console.log(str2)
console.log(sentence)
```

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CS87PK87PSdXLyU%2FXSAjJSFUoSi0uzSlRyCy2UlCpNjTVNjatTdC0BgA%3D)

```js
console.log(`The result is: ${15+35}`);
```

El estándar más reciente de ECMAScript (ECS6) define 8 tipos de datos, 7 primitivos y uno especial:

<ul>
    <li><b>Numbers</b></li>
    <li><b>String</b></li>
    <li><b>Boolean</b></li>
    <li><b>Symbol</b></li>
    <li><b>BigInt</b></li>
    <li><b>Null</b></li>
    <li><b>Undefined</b></li>
    <li><b>Object</b></li>
</ul>

### Declaración de objetos (Object)

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CVY1BDoIwFET3PcWEDZAQAyibNl7ADYkXMBQmRlL7Sam6INzdqBvcvjd54xghdsQRiwKA3nVPXiqNlI53%2BihVWmxMrdGU5YbsNWJ4cLs5aCyrWs0HqV78LI47J9dM7JgbpSQGab%2Bfni%2B0dmQfs8XL3QZqJCeZmRToJjp3G0QjOXOOgZMkBTh0g0bdrLn5T%2F%2BauXkD)

```js
let obj = {
    clave_1: 'elemento1',
    clave_2: 500,
    clave3: true,
    clave_4: {}
};
   
console.log(obj);

otroObj = new Object({nombre: "Jose", apellido: "Restrepo", edad: 25});
console.log(otroObj);
```

### Conversiones

A string

```js
let a = 250; // Variable 'a' de tipo Number 
let a_string = String(a); // Conversion de 'a' a string
let oneHundred = String(100); // Conversion de un literal numerico a string
let sum = String(20+30); // Conversion del resultado de una expresion a string
```

A bool

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CS87PK87PSdXLyU%2FXcMrPz0lNzNMw1NS05krGImGAS0IpIzUnJ18JpzRIBgA%3D)

```js
console.log(Boolean(1));
console.log(Boolean(0));
console.log(Boolean("hello"));
console.log(Boolean(""));
```
De strings a números

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CfY5NCoNADEb3c4qPWenGUYugiFdw1wNYTUXQROan0NsXtS4K0l3gvZekF3YyUzLLGLVheZCNdJHkhY7jGsZgm1V%2FIeE00kv8n1YVyvJU2q5VyhjcGbKS7Qax4LCQnXpBL%2FyayHo6GA%2FiNncgiLfi4KdVHLojEHd5rUg1DHSWbheVmsnDeYsGOstvuv5p%2FHsleW78eM55O%2FG4NxwWNPgu3YWrksNylLyLHw%3D%3D)

```js
console.log(Number("5.25")); // 5.25
console.log(Number(" ")); // 0
console.log(Number("")); // 0
console.log(Number("99 88")); // NaN

// Un operador numerico convierte operandos
// de otros tipos a numeros
console.log(Number("50" / "10"));

let str = "123";
console.log(typeof str); // string
let num = Number(str);
console.log(typeof num); // number
```

## Operadores en Javascript

### Operador +

Adición para operandos de tipo Number.
Concatenación para operandos de tipo String.
Operandos Number + String resulta en conversión del tipo Number a String y después hace una concatenación.

Orden de ejecucuión de operadores + es de izquierda a derecha.

3+4+'5' = "75"

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2Cy0ktUcgvULBVMNY20lY3NFC35krOzyvOz0nVy8lP18gv0EQVKKksSM1PUwCJAwA%3D)

```js
let op = 3+2+'10';
console.log(op); // '510'
console.log(typeof op); // String
```

El operador + también permite convertir datos de tipos no numéricos en números.

```js
console.log(+true); // 1
console.log(+""); // 0
```

Los demás operadores aritméticos no realizan conversiones al tipo String, solo al tipo Number.

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CS87PK87PSdXLyU%2FXMFLQVVA3VNe05kpGElU3U1fQV1A3Aolz5aSWKCTnl%2BaVpBYp2CooGRoooSouqSxIzU%2BDKQGbBGZpa%2BNXBwA%3D)

```js
console.log(2 - '1'); // 1
console.log('6' / '2'); // 3

let counter = "10";
console.log(typeof counter); // string
counter++;
console.log(typeof counter); // number
```

Los operadores de comparación: &lt;, &gt;, &lt;=, &gt;=, !=, ==. Si son aplicados a operandos de distintos tipos homogenizan los operandos a un mismo tipo.

```js
console.log(1 != 2); // true
console.log(2 != '2'); // false
console.log(1 != true); // false
console.log(0 == false); // true
```

La comparación entre datos de tipo string es lexicográfica, es decir, letra por letra.

```js
console.log("Z" > "A"); // true
console.log("Want" > 'Walk'); // true
console.log('TO' > 'To'); // false
```

Los operadores de igualdad "estricta" no solo comparan los valores sino que también verifican los tipos de datos.

```js
console.log(1 == "1"); // true
console.log(1 === "1"); // false
console.log('Walk' !== 'Walk'); // false
```

Opeadores de operaciones booleanas

<ul>
    <li><b>||: </b>Operación OR</li>
    <li><b>&&: </b>Operación AND</li>
    <li><b>!: </b>Operación NOT</li>
    <li><b>!!: </b>Conversión a Boolean y operación NOT</li>
</ul>

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CS87PK87PSdXLyU%2FXUFQ01LTmSkYRUTJQ0rQGAA%3D%3D)

```js
console.log(!!1); // true
console.log(!!"0"); // true
```











