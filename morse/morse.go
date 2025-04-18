// Definir el paquete morse
package morse

// Importar el paquete strings para trabajar con cadenas de texto
import "strings"

// reverseCode almacena un mapa inverso del alfabeto Morse
var reverseCode map[string]string

// Inicializar mapa inverso usando datos de morseCode
func init() {
	reverseCode = make(map[string]string)
	for letter, code := range morseCode {
		reverseCode[code] = string(letter)
	}
}

// Diccionario Morse
var morseCode = map[rune]string{
	'A': ".-",
	'B': "-...",
	'C': "-.-.",
	'D': "-..",
	'E': ".",
	'F': "..-.",
	'G': "--.",
	'H': "....",
	'I': "..",
	'J': ".---",
	'K': "-.-",
	'L': ".-..",
	'M': "--",
	'N': "-.",
	'Ñ': "--.--",
	'O': "---",
	'P': ".--.",
	'Q': "--.-",
	'R': ".-.",
	'S': "...",
	'T': "-",
	'U': "..-",
	'V': "...-",
	'W': ".--",
	'X': "-..-",
	'Y': "-.-",
	'Z': "--..",
	'0': "-----",
	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",
	' ': "/",
}

// Función para convertir texto a Morse
func TextToMorse(text string) string {
	// Convertir texto a mayúsculas
	text = strings.ToUpper(text)
	// slice para almacenar Morse equivalente a cada letra
	var result []string
	// Recorrer cada caracter del texto
	for _, char := range text {
		// Buscar valor del caracter en el diccionario
		if code, ok := morseCode[char]; ok {
			// Si existe en el diccionario, agregar al código Morse
			result = append(result, code)
		} else {
			// No existe en el diccionario, reemplazar el caracter por un '?'
			result = append(result, "?")
		}
	}
	// Unir códigos morse almacenados en "result" y separarlos por espacios
	return strings.Join(result, " ")
}

// Función para convertir Morse a texto
func MorseToText(morse string) string {
	// slice para almacenar letra/número equivalente a cada código Morse 
	var result []string
	// Dividir cadena de texto Morse para almacenar cada código Morse por separado
	words := strings.Split(morse, " ") // La divide en partes usando el espacio " " como separador

	// Recorrer cada código morse en la lisa (slice) words
	for _, symbol := range words {
		// Verificar si el morse existe en el mapa
		if letter, ok := reverseCode[symbol]; ok {
			// Existe, convertirlo en letra y agregarlo al slice result
			result = append(result, letter)
		} else {
			// No existe, reemplazar el/los caracteres ingresados por "?"
			result = append(result, "?")
		}
	}
	// Unir letras almacenadas en "result" y NO separarlos, por eso el ""
	return strings.Join(result, "")
}
