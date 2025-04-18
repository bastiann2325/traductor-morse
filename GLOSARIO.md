# Glosario para el archivo morse.go

- map[]: vector asociativo, es el equivalente a un diccionario o arreglo asociativo en otros lenguajes de programación.
- range: cláusula utilizada para recorrer colecciones como slices (arreglos).
- rune:
- " _ ": el caracter _ significa que es una variable de descarte. Lo verás en los ciclos for de este programa, en esos casos significará que se va ignorar el índice (porque no se necesita).
- ToUpper(): método que se utiliza para convertir una cadena de texto a mayúsculas.
- Join(): método que se utiliza para combinar elementos de una lista (slice) de cadenas en una sola cadena de texto. 
- Split(): método utilizado para dividir una cadena de subcadenas con base en un delimitador específico (el pasado por parámetro).
- append(): función para agregar elementos al final de un slice.

# Glosario para el archivo main.go
El archivo main.go es el punto de entrada al programa, está encargado de levantar un servidor web. Tiene como función traducir entre texto y código Morse, usanso el paquete creado en morse.go.

- encoding/json: paquete estándar de Go, permite leer y escribir datos en formato JSON.
- html/template: paquete para trabajar con archivos HTML como plantillas.
- log: paquete para imprimir mensajes o errores en consola.
- net/http: paquete que permite crear servidores, manejar rutas, respuestas HTTP, etc.
- morse: paquete personalizado que creamos, corresponde al archivo morse.go y contiene las funciones TextToMorse y MorseToText.
