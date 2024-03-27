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

### Operadores aritméticos

El operador +

Adición para operandos de tipo Number.
Concatenación para operandos de tipo String.
Operandos Number + String resulta en conversión del tipo Number a String y después hace una concatenación.

Orden de ejecucuión de operadores + es de izquierda a derecha.

<code>3+4+'5' = "75"</code>

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

### Operadores de comparación

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

### Operadores relacionales (operaciones booleanas)

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

## Estructuras de control de flujo de programa

### Condicionales

if...else if...else

```js
if(condicion){
   // bloque de sentencias primera condicion
}
else if(condicion alternativa){
   // bloque de sentencias condicion alternativa
}
else{
   // ejecucion cuando no se cumple ninguna condicion anterior
}
```

Switch case

```js
switch(opcion){
   case 'valor_1':
      // bloque de sentencias primer caso
      break
   case 'valor_2':
      // bloque de sentencias segundo caso
      break
   default:
      // bloque de sentencias caso opcional por defecto
      break
}
```

Operador ternario

```js
let resultado = condicion ? opcion1 : opcion2;
```

### Ciclos de repetición

Ciclo while

```js
while(condicion){
   // bloque de sentencias a repetir en el ciclo
   // mientras se cumpla la condicion
}
```

Ciclo do...while

```js
do{
   // bloque de sentencias a repetir en el ciclo
   // mientras se cumple la condicion
   // o al menos una vez así no se cumpla la condicion
}while(condicion);
```

Ciclo for

```js
for(let i=0; i < N; i++){
   // bloque de sentencias a repetir en el ciclo
   // para i desde 0 hasta i menor que N
}
```

Ciclo for...in

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CNYtBCoAgFAX3nuItAhOiC4QnCUGzXwjiB2sRiHcPIzezmDfPc7puuJzpjAyNVTo5QW4NvmFvIGkWcXAeI90IaacHIfWXKgKA53RxpDnyOdp%2FWYfyxdVAYyjdfs5Uq0R9AQ%3D%3D)

```js
const arreglo = ['a', 'b', 'c', 'd', 'e'];
for(let index in arreglo){
   console.log(`arreglo[${index}] = ${arreglo[index]}`)
}
```

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CHcxBCoMwEEbh%2FZziXwijID3AiCcpgiZOiyY4UN2EkLuXZPe9zYv6wNyJGZmATcAbjwQ4AbsmL2DftAt4b1IBK1OZiD7266M%2BCJpwXPU11BO8XbdFfUX79qu5893loKksmNHl2kHTUtaByh8%3D)

```js
let obj = {
  a: 'a',
  b: 'b',
  c: 'c',
  d: 'd',
  e: 'e'
};

for(let key in obj){
   console.log(`obj[${key}] = ${obj[key]}`)
}
```

Ciclo for...of

Funciona solo con iterables, por tanto no funciona con variables de tipo Object.

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CNcsxDoAgEETRnlNMt5AYL2A8ibFAHGhWNkE6490Nhc2r%2Fld2xNZY1LBikygT5BikwTmg7IvL1ryyg8qLtRss%2F2N4HIBk9TblrFb8HwX3fg%3D%3D)

```js
let arreglo = ['a', 'b', 'c', 'd', 'e'];
for(let elemento of arreglo){
   console.log(elemento)
}
```

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CJcsxCoAwDAXQvaf4OLWLFxA3LxJjqoXQSFpwEO8u4vx4Kh3ZqQlmDIs0cjdVwyagUwsTF6vScMk6TCGbR5UOPshh%2BZ%2FpDgDYajOVUW2PH6fwvA%3D%3D)

```js
let frase = "Desarrollo de aplicaciones web";
for(let char of frase){
   console.log(char)
}
```

## Funciones

Para declarar funciones en Javascript se utiliza la palabra <code>function</code>, a continuación se indica el nombre de la función y entre paréntesis se define la lista de parámetros de entrada de la función. El bloque de código de la función se agrega entre llaves.

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CSyvNSy7JzM9TKCjKzCsJTs0rSc1LTtVILEqP0FFILEqP1KzmUlBQSM7PK87PSdXLyU8HyyloK6grlOQrqCtoQ1Rx1XJxoRqhHp6ak5yfm6quo6DulViWWJxclFlQoq5pDQA%3D)

```js
function printSentence(argX, argY){
   console.log(argX + ' to ' + argY)
}

printSentence('Welcome', 'Javascript');
```

### Valores por defecto en argumentos de funciones

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CSyvNSy7JzM9TKCjKzCsJTs0rSc1LTtVILEqP0FFILEqPtFX3SixLLE4uyiwoUdes5lJQUEjOzyvOz0nVy8lPBytU0FZQVyjJV1BX0AZr0eSq5eJCNU89PDUnOT83VV3TGgA%3D)

```js
function printSentence(argX, argY='Javascript'){
   console.log(argX + ' to ' + argY)
}

printSentence('Welcome');
```

### Retorno en funciones

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CSyvNSy7JzM9TyC3NKcksyKnUSNRRSNKs5lJQUChKLSktylNI1Eqy5qrl4spJLVEoSi0uzSlRsEUoN9VRMNa05krOzyvOz0nVy8lP14Ao0gQA)

```js
function multiply(a, b){
   return a*b;
}

let result = multiply(5, 3);
console.log(result);
```

### Funciones como expresiones o funciones anónimas. No olvidar el <code>;</code> al final de la declaración.

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2Cy0ktUSgoyswr8UjNyclXsFVIK81LLsnMz9PIS8xN1azmUlBQSM7PK87PSdXLyU%2FXSICoU6kGSdcqKiomaHLVWnNxIczQUHJOLMrPycxLVNK0BgA%3D)

```js
let printHello = function(name){
   console.log(`Hello ${name}!!!`)
};

printHello("Carolina");
```

### Callbacks

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CnZFBTsMwEEX3OcWQlS1FXCALFkgsWBQhJCSWk%2BmEWjUzqSfuBvVInKIXq5KmqYoqBHg1%2Bt%2F%2B%2Fn5us1AfVIAwxgZp%2FZCF3CazDWo1ym9sx2Gh%2FrMAAAitI5U2pI95q5%2BsYU2HnK9HaQccjS%2Fthc5usSuK9lQDxZ7WbsoiFdPIt1HfXfmcGZrAUkHPYCFtGZYMiTs0vSl9%2FT3mHoU4Xo16RQOETpecALvEMg6bvP%2Baci5glC%2FYBDRYIXGaORmwwCNu0SiFrr8rq2P36ny3r%2F%2BZNDQ%2BPcWduf4Nx4D2V0k%2F0xh%2FyNcH)

```js
function callbackFunc(question, callYes, callNo){
    if(confirm(question)){
        callYes();
    } else{
        callNo();
    }
}

function ansOk(){
    console.log("Que bien, te sirve de repaso!");
}

function ansCancel(){
    console.log("Vas a poder aprender aquí");
}

callbackFunc("Sabias hacer callbacks en Javascript?", ansOk, ansCancel);

callbackFunc("Sabias hacer callbacks en Javascript?",
   function(){
       console.log("Que bien, te sirve de repaso!");
   },
   function(){
       console.log("Vas a poder aprender aquí");
   }
);
```

### Funciones flecha: Arrow functions

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2Cy0ktUcgtzSnJLMipVLBV0EjUUUjSVLC1U0jUSrLmykktUShKLS7NKVGwhSvTMNWx0LTmSs7PK87PSdXLyU%2FXgKjRtAYA)

```js
let multiply = (a, b) => a*b;
let result = multiply(5,8);
console.log(result);
```

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CVYw7CoAwEAX7PcUrEwl2ghDiXaIsIqxG4qYQ8e6Cn8J2ZhhhxVxEp1V2BJjo0FuEDgcBgLAi81ZEERCr3t80s5a8vMKDTqJf%2BA1N41rraUjLloRrSaN5Gusv)

```js
let multiply = (a, b) => {
    let result = a*b;
    return result; 
}

let result = multiply(5,8);
console.log(result);
```

## Métodos básicos asociados a los tipos de datos

### parseInt y parseFloat

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2Cy0ktUcgrzVWwVShILCpO9cwr0VA3NSioUNe05spJLVFISU2Gybnl5CeWaKgbGekZ6Zmoa3JxJefnFefnpOrl5Kdr5JXmalqjiKSkJmtaAwA%3D)

```js
let num = parseInt('50px');
let dec = parseFloat('22.2.4')

console.log(num);
console.log(dec);
```

### toString

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2Cy0ktUcgrzVWwVTAyNbXmSs7PK87PSdXLyU%2FXyCvN1SvJDy4pysxL1zA009TEI22EV9bQQFPTGgA%3D)

```js
let num = 255;
console.log(num.toString(16));
console.log(num.toString(2));
console.log(num.toString(10));
```

### length, charAt y operador []

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2Cy0ktUSguKVKwVVBySS1OLCrKz8nJV0hJVXAsKChWCE9NUrLmSs7PK87PSdXLyU%2FXKC4pijaI1cQQ1EvOSCxyLNEw1cSUiwbJ56TmpZdkKOgqGMZqAgA%3D)

```js
let str = "Desarrollo de Apps Web";
console.log(str[0]);
console.log(str.charAt(5));
console.log(str[str.length - 1]);
```

### uppercase y lowercase

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2Cy0ktUSguKVKwVVAvd81x9s91VcjI93VV50rOzyvOz0nVy8lP1yguKdIryQ8tKEgtck4sTtXQ1LTGIu%2BTX46QBwA%3D)

```js
let str = 'wElCOmE hoME'
console.log(str.toUpperCase());
console.log(str.toLowerCase());
```

### indexOf

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2Cy0ktUSguKVKwVVBySS1OLCrKz8nJV0hJVXAsKChWCE9NUrLmSs7PK87PSdXLyU%2FXKC4p0svMS0mt8E%2FTUEfoUNfUxKMOZBZ%2BFeWpSSAFAA%3D%3D)

```js
let str = "Desarrollo de Apps Web";
console.log(str.indexOf('Desarrollo'));
console.log(str.indexOf('Apps'));
console.log(str.indexOf('web'));
```

### includes

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2Cy0ktUSguKVKwVVBySS1OLCrKz8nJV0hJVXAsKChWCE9NUrLmSs7PK87PSdXLyU%2FXKC4p0svMS84pTUkt1lAPT01S19TEp6IcogIA)

```js
let str = "Desarrollo de Apps Web";
console.log(str.includes('Web'));
console.log(str.includes('web'));
```

### startsWith y endsWith

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2Chc7BCsIwEATQe79i6CXJxQ%2BwiAj6B0LP23bVwDYJuxHBr5ei9CR6nRmYJ1xByR6s2MG3RzZSzSIZE%2BNQiqHnoQ0bq6TV%2Blhv3i2xC9jDu5NBCIWEBiXEFMdI4gK2SzeXyE%2FCmBNy1XXmQteMOVkW3ki%2B%2Bvd96JpGuELZyp2t0m8Pp%2Bmj6Xn4irnEtFLOrHNM%2Fynreehe)

```js
let answer = ("Desarrollo de Apps Web").startsWith('Apps') ? ('Es la palabra inicial') : ('Empieza con otra palabra');
console.log(answer);

let respuesta = ("Desarrollo de Apps Web").endsWith('Web') ? ('Es la palabra final') : ('Termina con otra palabra');
console.log(respuesta);
```

### substr, substring y slice

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2Cy0ktUSguKVKwVVB3LMjJTE5MzszPSy1Wt%2BZKzs8rzs9J1cvJT9coLinSKy5NKi4p0jDRUTDT1MQlnZmXDlJhaIBfiaGBjoIJViU5mcmpcBMA)

```js
let str = 'Aplicaciones';
console.log(str.substr(4, 6)); // 6 caracteres a partir de la posicion 4
console.log(str.substring(4, 10)); // fragmento desde la posicion 4 a la 9
console.log(str.substring(10, 4)); // fragmento desde la posicion 4 a la 9
console.log(str.slice(4, 10)); // fragmento desde la posicion 4 a la 9
```
## Métodos para manipular arreglos

### Operador [], indexOf y length

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CS87PKy5RSCwqSqxUsFWIVg9PTVLXUVBPLCjIyUxOLMnMzysG8VNSy1Jz8gtyU%2FNK1GOtuZLz84rzc1L1cvLTNcB6ow1jNbEI62XmpaRW%2BKdpgM3VxKokJzUvvSRD0xoA)

```js
const array = ['Web', 'applications', 'development'];
console.log(array[1]);
console.log(array.indexOf('Web'));
console.log(array.length);
```

### arreglo a String

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CS87PKy5RSCwqSqxUsFWIVg9PTVLXUVBPLCjIyUxOLMnMzysG8VNSy1Jz8gtyU%2FNK1GOtuZLz84rzc1L1cvLTNYJLijLz0jXARmhqWgMA)

```js
const array = ['Web', 'applications', 'development'];
console.log(String(array));
```

### push, pop, shift y unshift







