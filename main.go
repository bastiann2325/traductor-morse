package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"github.com/bastiann2325/traductor-morse/morse" 
)

func main() {
	// Cargar archivo html como plantilla web
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))

	// Ruta principal (sirve el HTML), muestra el contenido de index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Escribe el html en la respuesta web
		tmpl.Execute(w, nil)
	})

	// Servir archivos estáticos (CSS, JS, etc.)
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/web/static/", http.StripPrefix("/web/static/", fs))

	// Definir ruta de tipo API para recibir datos y responder con una traducción
	http.HandleFunc("/traducir", traducirHandler)

	// Iniciar servidor en el puerto 8080
	log.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	// Si occurre algún error al iniciar el servidor, se muestra aquí
	if err != nil {
		log.Fatal("Error iniciando servidor:", err)
	}
}

// Handler para procesar traducciones, esta función se ejecuta cuando alguien hace una petición POST a /traducir
func traducirHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar que se usa el método POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Leer JSON enviado desde el navegador
	var datos struct {
		Texto string `json:"texto"`
		Morse string `json:"morse"`
	}

	err := json.NewDecoder(r.Body).Decode(&datos)
	if err != nil {
		http.Error(w, "Error al leer JSON", http.StatusBadRequest)
		return
	}

	var resultado string
	// Si los datos que llegaron son texto, convertir a morse
	if datos.Texto != "" {
		resultado = morse.TextToMorse(datos.Texto)
	// Si los datos que llegaron son morse, convertir a texto	
	} else if datos.Morse != "" {
		resultado = morse.MorseToText(datos.Morse)
	}

	// Responder con un JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"resultado": resultado,
	})
}
