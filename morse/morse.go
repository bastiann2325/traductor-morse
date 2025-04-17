package morse

import "strings"

// Almacenar mapa inverso
var reverseCode map[string]string

// Reconstruir mapa inverso
func init() {
	reverseCode = make(map[string]string)
	for letter, codigo := range morseCode {
		reverseCode[codigo] = string(letter)
	}
}

// Alfabeto Morse
var morseCode = map[rune]string {
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
	// Variable para almacenar Morse equivalente a cada letter del texto
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
	var result []string
	words := strings.Split(morse, " ")

	// Diccionario inverso
	reverse := make(map[string]string)
	for letter, codigo := range morseCode {
		reverse[codigo] = string(letter)
	}

	for _, symbol := range words {
		if letter, ok := reverseCode[symbol]; ok {
			result = append(result, letter)
		} else {
			result = append(result, "?")
		}
	}

	return strings.Join(result, "")
}