# Javascript: lo básico

## Variables y tipos de dato

### Declaración de variables

Para declarar variables se puede usar: <code>let</code>, <code>var</code> y <code>const</code>

La diferencia entre <code>let/code> y <code>var</code>: <code>var</code> permite re-declarar variables, mientras que con <code>let</code> ocurre un <code>SyntaxError</code>. Se prefiere el uso de <code>let</code> ya que <code>var</code> se usaba en las versiones más antiguas de Javascript.

Para declarar variables de solo lectura se utiliza <code>const</code>. Se recomienda el uso de mayúsculas (<i>UPPERCASE</i>) para las variables calificadas con <code>const</code>.

Los nombres de las variables pueden empezar por letra, _ o $, nunca por número y no pueden tener caracteres especiales. Los nombres de variables son sensibles a mayúsculas y minúsculas (case-sensitive).

```js
let globalVariable = "global variable";

function myFun(){
    let localVariable = "local variable";
    return localVariable;
}

console.log(globalVariable);
console.log(myFun());
```

[example](https://codapi.org/embed/?engine=browser&sandbox=javascript&code=data%3A%3Bbase64%2Cy0ktUUjPyU9KzAlLLMpMTMpJVbBVUIKIKJRBhZSsubjSSvOSSzLz8xRyK91K8zQ0q7kUFBQUclJLFHLyk1F1gwWQNYNUFqWWlBbloSq25qrl4krOzyvOz0nVy8lP10B1iaY1iiTUYk1rAA%3D%3D)

```js
const MY_BIRHTDAY = '29.02.1996'
```

Las variables de tipo string se pueden inicializar con doble comilla o comilla simple. Para pasar valores de variables o expresiones a un <code>string</code> se debe crear el <code>string</code> con tildes y pasar la variable usando esta sintaxis: <code>${variable}</code>.
