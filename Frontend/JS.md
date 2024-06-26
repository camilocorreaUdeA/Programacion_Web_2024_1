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

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CXY1BCsMgEEX3nmJ2o1ByAck5uihdaDJtLBMVnST09sUkFNrd%2F%2B8%2F%2BEwCrhT3hh5ueCWPF0CXM4fBSUixtj7SSpzyTFHwbtXud3mpk8aXW10dSsjSxElmRmMVk0DL0MMpp6xPvpH%2F4jqFh7ThqEs8AG7k0aghxZqYOk5PvQvG%2FrD28Ic28sZ%2BAA%3D%3D)

```js
let array = ['Web', 'applications', 'development'];
array.push('javascript', 'html');
let html = array.pop();
let web = array.shift();
array.unshift('web')
console.log(array);
console.log(html);
console.log(web);
```

### splice, join y split

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CXY3BCsIwEETv%2BYr1lAZqQa%2Bh3%2BAHiIcl3crCNhvcFATx38UGFbwNM29mhCqQVZxYDUY4e6yY0Vh9D76gpRXlLReqNy0qXDFvWcKlcG6ULGiU0F%2BiS5qtAueJ7jB%2Bl4fNOM3d306IjmfoGr4bYX8IDwfw61kRTtSAHo4huud2oUKD6LX7gCG%2BAA%3D%3D)

```js
let estadios = ['atanasio', 'pascual', 'metropolitano', 'campin', 'palmaseca'];
const index = estadios.indexOf('metropolitano');
if (index != -1){
   estadios.splice(index, 2); // remueva dos elementos a partir del indice
}
console.log(estadios);
```

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CXYxBCsMwDATvfYXwJQ44%2FkDIK%2FoCYYuiIlumUv9fXJJLjrs7O0IOZI6V1eCAgI4djTXBQCtflASN%2FKNDhR27JijYBve5S0OjgmF%2FFO3mYF6n49JlG8IelwTLeiFv5f70CscfzjPGsG3hBFQoi77iHG%2FV%2BVz3Hw%3D%3D)

```js
let estadios = "atanasio, pascual, metropolitano, campin, palmaseca";
const stds = estadios.split(', ');
const joinStd = stds.join("--");
console.log(stds);
console.log(joinStd);
```

### find, filter, map y some

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CndCxasMwEAbgXU9xnRxDMKTtZKMMbQoZ8gahg7AvieAkBetMh9B3D5Kj1rWdUqpBcJL49N8RMiirjSIPEvYCAOAS97B0U8Jq%2BVVaZbCE7FVx9n3oWXHnS8gaZ9CzrhVjk8Xrz%2BWM9zjx3gjPJ2Vn0Q9Nv2FPE2yLrbP%2FkJ4n0ktH9Mc%2BxXslBCFD3aLirkWQaarFQdtm0Rcg17fjIvwBUvbDzCtRO%2BsdYUHuuEhIfkND8p0%2B%2FESJsZ2yfcgIx35HcpKSHFL4Hdojnwa4Uef5wAXFpyN0gCTXO4ObwZgGeLi6m%2FvhTu6xl1dX)

```js
let animals = [
    {
        id: 1,
        name: 'Cat',
        status: 'domesticated'
    },
    {
        id: 2,
        name: 'Elephant',
        status: 'wild'
    },
    {
        id: 3,
        name: 'Heron',
        status: 'wild'
    },
    {
        id: 4,
        name: 'Bull',
        status: 'domesticated'
    }
];

let creature = animals.find(animal => animal.name == 'Cat');
console.log(creature);

let wildLife = animals.filter(animal => animal.status == 'wild');
console.log(wildLife);

let namesLength = animals.map(animal => animal.name.length);
console.log(namesLength);

let someDomesticated = animals.some(animal => animal.status != 'wild');
console.log(someDomesticated);
```

## Maps y Sets

Los mapas son colecciones conformadas por elementos que son pares de clave y valor. Donde las claves son únicas y pueden ser de cualquier tipo.

### Métodos principales asociados a los mapas

<ul>
    <li><b>new Map(): </b>función para crear un mapa.</li>
    <li><b>set(clave, valor): </b>permite almacenar un nuevo par clave-valor en el mapa.</li>
    <li><b>get(clave): </b>retorna el valor asociado a la clave. <code>undefined</code> si la clave no existe en el mapa.</li>
    <li><b>has(clave): </b>verifica si la clave existe en el mapa.</li>
    <li><b>delete(clave): </b>elimina un par clave-valor del mapa.</li>
    <li><b>clear(): </b>elimina todos los elementos del mapa.</li>
    <li><b>keys(): </b>retorna un arreglo con todas las claves del mapa.</li>
    <li><b>values(): </b>retorna un arreglo con todos los valores del mapa.</li>
    <li><b>size: </b>retorna la cantidad de elementos almacenados en el mapa.</li>
</ul>

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CfY5NCoMwEIX3OcXskkAIycKVeAQPMdipFoZEmkihpy%2BJFrVIVwNv3s%2FHlGEJPc4IHQR6QY%2Bz0q1YNZsoK%2BmdNCCnuCSS55d3BprmrN2RExmoR7diiCFFJstxVF%2FX410%2BgilDbYVuY7DjtvcTrK4SuWibMK2beue4EVMm5d01QIl4d%2FAPTPhU%2F2g%2F)

```js
let unMapa = new Map();
unMapa.set('10', 'house');
unMapa.set(10, 55);
unMapa.set(false, false);
console.log(unMapa.size);

let house = unMapa.get('10');
console.log(house);

console.log(unMapa.has(false));
unMapa.delete(10);
console.log(unMapa.has(10));
unMapa.clear();
console.log(unMapa.size);
```

### Ciclos iterativos sobre mapas

<b>for...of</b>

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CbY5BCsIwEEX3OcXsmkApsSgixa07T1C6GMq0qGmmJKmi4t0lFi2lbh%2FvvxlDAXpH9Yn9wQ0BPezB0g2O2MtSAECZdGgfaDFJYbXVukpHatGhPUe62%2Fxgh7blKOaT2LGLVr7WuhIAlVCFEA07aShAE28CN%2FMfsgvdvVTqGQs1W8%2BGMsOt%2FOiqEK%2BpMA6XiSuagf5GRm9eIUMd2bDsLOdfNQbe)

```js
let preciosFrutas = new Map([
   ['manzana', 1700],
   ['naranja', 950],
   ['mango', 1200],
   ['mora', 2400]
  ]
);

for(let fruta of preciosFrutas.keys()){
   console.log(fruta);
}

for(let precio of preciosFrutas.values()){
   console.log(precio);
}

for(let elemento of preciosFrutas){
   console.log(elemento);
}
```

<b>forEach</b>

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CVY3RCoIwGIXv9xT%2FheAGIkuKKLG7uusJRPBnLDPmftnUIPHdYwZFt9%2F5zjlGD9A7rVryFzcO6KEAq59wxZ6XDADKuEP7QotxApu9lFXyoRYd2kegh90XdmgbCmL2Eztywcq2UlYMoGIiZ%2BzvMr2RO6O6cz6hIZeAMjhpAcUJ5rChyHoyOjXU8Dqa13QBNWo%2F4BGieW0ttcjZIt4%3D)

```js
let preciosFrutas = new Map([
   ['manzana', 1700],
   ['naranja', 950],
   ['mango', 1200],
   ['mora', 2400]
  ]
);

preciosFrutas.forEach((valor, clave) => {
   console.log(`${clave} cuesta: ${valor}`);
})
```

Los <b>sets</b> son colecciones donde cada elemento es único y no se repite dentro de la colección.

### Métodos principales asociados a los sets

<ul>
    <li><b>new Set(): </b>función que permite crear un <code>set</code>. Se le puede pasar como argumento de entrada un arreglo o cualquier otro iterable.</li>
    <li><b>add(valor): </b>agrega un nuevo valor a la colección.</li>
    <li><b>delete(value): </b>elimina un valor de la colección.</li>
    <li><b>has(valor): </b>verifica si la colección tiene a valor.</li>
    <li><b>clear(): </b>elimina todos los elementos de la colección.</li>
    <li><b>size: </b>retorna la cantidad de elementos de la colección.</li>
</ul>

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CfY29CgIxEIT7PMV0SUDuBcLZ6QtYqsV6bCKS24VcDlHx3cXgX2U3zHwzk7li4rouc6UJPYTP2HB1WzuSXEnILmCFCsmpyZEkaRNampFmutDhnbXG3gcTtbjMFfE5DI3fE38zADCoTJq5y5pcg3wwd%2FPr2p1YH8yn2EUtKxqO7sWjX%2BLPlg8P)

```js
let setFrutas = new Set(['manzana', 'naranja', 'mango', 'mora', 'guayaba', 'manzana']);
for(let fruta of setFrutas){
    console.log(fruta);
}
console.log('\n');
setFrutas.forEach((fruta) => {
    console.log(fruta);
});
```

## Clases y objetos en Javascript

Definición de una clase

```js
class ClassName{
   // variables estaticas de la clase
   static variable = "valor"
   // metodos de la clase
   constructor(args...){
   }
   classMethod1(args...){
   }
   classMethod2(args...){
   }
   // ...
}
```
Declaración de objetos de una clase

```js
let obj = new ClassName(argumentos_constructor...);
obj.classMethod1(args...);
```

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2ChVE9a8QwDN39K7TZgZDh1nBD6d6lY7lBdYUvxbEPy%2BaGI%2F%2B9KB%2BmTUnrwcZ6T3rvIeuRGV5z%2BaCQHwoAwMbAORWbYzIBR2oBHbVgY0lM3CwkOfk6cCcMOIM8%2FU8EnQDo9vV1EJy3kQs%2BzTdTfsGRZuF%2FpWrLkyOD7lfDzkGlPy%2B65iDSnwbdarB5QKJcUvjmrTLEz44gXiq%2BGdhxNuFJTUp5ysDLXiQ03bctGf0Zr0G3cDq18KZvKbqEI9ohSlHjzQ92%2FhHDnd6lZtHb4qO%2BNL2S9UZPnY%2FOrPO7GuoYnyMdwzVR038B)

```js
class Student{
    constructor(name, age, courses){
        this.name = name;
        this.age = age;
        this.courses = courses;
    }
    setName(name){
        this.name = name;
    }
    setAge(age){
        this.age = age;
    }
    setCourses(courses){
        this.courses = courses;
    }
    getName(){ return this.name }
    getAge(){ return this.age }
    getCourses(){ return this.courses }
}

let student = new Student('john', 22, ['programacion', 'aplicaciones web', 'calculo']);
console.log(student.getName());
console.log(student.getAge());
console.log(student.getCourses());
```
Herencia entre clases en Javascript

Para heredar de otra clase se utiliza la palabra <code>extends</code>. Y para utilizar el <code>constructor</code> de la clase base se utiliza el método <code>super</code> con los argumentos de entrada correspondientes.

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CdZDNasMwDIDvfgphdnCg5AVML9t57BWqOqIrOPaQFdqu5N2HREYW2vlgg%2F78fUoZW4M3ZLg7AEi1NOEpSeVwZCxDZ2E98nluvcVgD%2FZGTc16nUheNRLWciaZuPzpWqpn51wmgYQMeyh00c%2BDH%2FF7QN9FpwQ1U5%2FrKSTkfh3dRedcMt73mxLTVagMRv8cfgcFR1qZ2vRFv15x41UvhYwHR9pofWjiUevwcl%2FNZpPREQMcb7CkbOR8eLDms9TF3DSCZyo4ZfE78CM92YF2bPfwT8EC28Uf)

```js
class Car {
   constructor(brand){
       this.brand = brand;
   }
   getBrand(){
       return this.brand;
   }
}


let car = new Car("mazda");
console.log(car.getBrand());


class MyCar extends Car{
   constructor(brand, name){
       super(brand);
       this.owner = name;
   }
   getOwner(){
       return `${this.brand} car owned by ${this.owner}`
   }
}


let carrito = new MyCar("renault", "me");
console.log(carrito.getBrand());
console.log(carrito.getOwner());
```

## Try, catch y throw. Manejo de excepciones en Javascript

La sentencia <code>try</code> permite definir un bloque de código que se prueba contra errores al tiempo que se va ejecutando. Mientras que la sentencia <code>catch</code> permite definir el bloque de código que se va a ejecutar en caso de que haya ocurrido un error (excepción) cuando se ejecutaba el bloque definido por <code>try</code>.

```js
try{
   // Bloque de codigo a ejecutar y probar contra errores
}
catch(err){
   // Bloque de codigo a ejecutar cuando ocurre algun error
}
```

Para crear errores personalizados (definidos por el programador) se puede utilizar la sentencia <code>throw</code>. Esta sentencia permite lanzar un error o excepción que puede ser de alguno de los siguientes tipos: String, Number, Boolean, Object o un objeto de alguno de los tipos de error nativos de Javascript (Error, SyntaxError, ReferenceError, TypeError, EvalError, RangeError, URIError).

```js
throw "ha ocurrido un error";
throw 404;
throw true;
throw {
   code: 404,
   message: "not found"
}
throw new Error("algo inusual acaba de ocurrir!");
```

[Ejemplo:](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2CLYoxDgIhFAX7f4oXqt3GYGGFlh6EsOxCQnjmgzFmw93NqlPOTGBtHSuJGy7WiXR9YxcAyCumI1xxtnb%2BuYOelC%2BYuyrVuK8eMiT4HtIUVf9rYG0s8VS4TcaXjcg1tkdUvxDJg%2BGpmhea2cn4AA%3D%3D)

```js
const foo = 50;

try {
    if (foo < 100){
        throw "Error";
    }
}
catch(err){
    console.log("algo inesperado ha ocurrido");
}
```

## Ejecución asincrónica: async, await y promises

La ejecución asincrónica de una aplicación se refiere a la habilidad de poder correr tareas que requieren de mucho tiempo de procesamiento sin necesidad de bloquear la ejecución normal del programa. Es decir, que a pesar de que se inicie la ejecución de una tarea que toma mucho tiempo en completar se puedan correr otras tareas y atender otros eventos mientras que la tarea de larga ejecución se completa.

Las funciones asincrónicas en Javascript permiten:

<ol>
    <li>Inicializar una tarea u operación que toma mucho tiempo mediante el llamado de una función.</li>
    <li>Permitir que la funcion inicie la tarea y retorne inmediatamente para que el programa pueda ejecutar otras funciones o eventos.</li>
    <li>La función ejecuta la tarea sin necesidad de bloquear el hilo principal de ejecución de la aplicación.</li>
    <li>La función notifica acerca del resultado de la tarea una vez esta se completa (eventualmente).</li>
</ol>

### Promise

La programación asincrónica en Javascript orbita alrededor del concepto de <code>Promise</code> (promesa), que es simplemente un objeto retornado por una función asincrónica y sirve para representar el estado en el que se encuentra la tarea de larga ejecución que fue iniciada por la función. El objeto <code>Promise</code> debe proporcionar métodos que permitan atender el eventual éxito o falla al completar la ejecución de la tarea.

Al usar <code>Promise</code> la función asincrónica da inicio a la tarea e inmediatamente retorna un objeto <code>Promise</code>. Luego, se pueden asociar <code>handlers</code> a ese objeto <code>Promise</code> para que sean ejecutados una vez la tarea se complete exitosamente o bien falle.

Handlers más comunes para trabajar con objetos <code>Promise</code>:

<ul>
    <li><code>.then()</code>: Permite manejar el resultado de la tarea asincrónica cuando el resultado de la ejecución fue exitoso.</li>
    <li><code>.catch()</code>: Permite manejar el resultado de la tarea asincrónica cuando la promesa fue rechazada, es decir, la operación lanzó un error.</li>
    <li><code>.all()</code>: Permite manejar en una sola promesa todos los resultados de varias promesas asociadas a distintas funciones asincrónicas.</li>
    <li><code>.race()</code>: Permite manejar a través de una promesa el resultado de la primera promesa en ser resuelta o rechazada de un conjunto de promesas asociadas a distintas funciones asincrónicas.</li>
</ul>

Una promesa, <code>Promise</code>, es algo que va a tomar cierto tiempo en ser completado y tiene 3 estados posibles: pendiente, completada o rechazada. Por tanto, una tarea asincrónica estará es progreso mientras la promesa no se haya completado. De allí que una promesa que alcanza el estado completado indica la ejecución exitosa de una tarea asincrónica.

Las promesas <code>Promise</code> se pueden crear utilizando el constructor de <code>Promise</code> de Javascript. Este constructor es una función que toma como argumento de entrada a otra función que se conoce como el <i>executor</i>.

El <i>executor</i> es una función que tiene dos argumentos de entrada llamados <i>resolve</i> y <i>reject</i> que también son funciones. Entonces, en la lógica del <i>executor</i> se hace el llamado de la función asincrónica. Si la función se ejecuta exitosamente entonces se llama la función resolve y si no entonces se llama a la función reject. De hecho, si la función asincrónica lanza una excepción entonces la función reject es llamada de forma automática.

```js
let promise = new Promise((resolve, reject) => {
   // este es el executor y aqui se debe llamar la funcion asincronica
});
```

Ejemplo:

```js
function alarma(nombre, tiempo) {
  return new Promise((resolve, reject) => {
    if (tiempo < 0) {
      reject("el tiempo para la alarma no puede ser negativo");
    }
    setTimeout(() => {
      resolve(`Beeep beeep beeep despierta, ${nombre}!`);
    }, tiempo);
  });
}
```
O lanzando un error (excepción):

```js
function alarma(nombre, tiempo) {
  return new Promise((resolve, reject) => {
    if (tiempo < 0) {
      throw new Error("el tiempo para la alarma no puede ser negativo");
    }
    setTimeout(() => {
      resolve(`Beeep beeep beeep despierta, ${nombre}!`);
    }, tiempo);
  });
}
```

Para manejar el objeto de la promesa se utilizan los handlers que vimos anteriormente (se puede hacer encadenamiento de handlers):

```js
alarma('estudiante', 5000)
   .then((respuesta) => console.log(respuesta))
   .catch((err) => console.log(`error: ${err}`));
```

### async y await

<code>async</code> y <code>await</code> son idiomas de Javascript que permiten escribir código asincrónico de manera que luzca como código sincrónico, dicho de otro modo, con esas palabras de Javascript se puede escribir código que parece que se ejecuta de forma secuencial cuando en realidad lo hace de manera asincrónica.

<code>async</code> y <code>await</code> simplifican la lectura y escritura de operaciones asincrónicas evitando las confusiones inherentes de utilizar handlers para atender las promesas de una función asincrónica.

La palabra <code>async</code> se utiliza para indicar que el bloque de código de una función se debe ejecutar de manera asincrónica. Por su parte, la palabra <code>await</code> solo puede usarse dentro de una función construida con <code>async</code> e indica que se debe esperar por la ejecución de cierta línea de código, en particular la ejecución de una función asincrónica que retorna una promesa, dicho de otro modo <code>await</code> pausa la ejecución hasta que se complete o rechace la promesa.

Reescribiendo el ejemplo anterior con <code>async</code> y <code>await</code>:

```js
(async () => {
    try{
        const res = await alarma('estudiante', 5000);
        console.log(res);
    }
    catch(err){
        console.log(err);
    }
})();
```

## Recursos recomendados

[https://htmlcheatsheet.com/js/](https://htmlcheatsheet.com/js/)<br>
[https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/await](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/await)<br>
[https://www.w3schools.com/js/js_promise.asp](https://www.w3schools.com/js/js_promise.asp)<br>
[https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Asynchronous/Introducing](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Asynchronous/Introducing)<br>
[https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Asynchronous/Promises](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Asynchronous/Promises)<br>
[https://semaphoreci.com/blog/asynchronous-javascript](https://semaphoreci.com/blog/asynchronous-javascript)<br>
[https://www.scaler.com/topics/types-of-errors-in-javascript/](https://www.scaler.com/topics/types-of-errors-in-javascript/)<br>
[https://developer.mozilla.org/en-US/docs/Learn/JavaScript](https://developer.mozilla.org/en-US/docs/Learn/JavaScript)<br>
[https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide)<br>
[https://learnjavascript.online/](https://learnjavascript.online/)<br>
[https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/JavaScript_basics](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/JavaScript_basics)<br>
[https://eloquentjavascript.net/](https://eloquentjavascript.net/)<br>
[https://www.javascripttutorial.net/](https://www.javascripttutorial.net/)<br>
[https://javascript.info/](https://javascript.info/)
