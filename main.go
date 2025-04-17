package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"github.com/bastiann2325/traductor-morse/morse" 
)

func main() {
	// Cargar template principal
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))

	// Ruta principal (sirve el HTML)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// Servir archivos estáticos (CSS, JS, etc.)
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/web/static/", http.StripPrefix("/web/static/", fs))

	// Endpoint para traducir
	http.HandleFunc("/traducir", traducirHandler)

	// Iniciar servidor
	log.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error iniciando servidor:", err)
	}
}

// Handler para procesar traducciones
func traducirHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

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
	if datos.Texto != "" {
		resultado = morse.TextToMorse(datos.Texto)
	} else if datos.Morse != "" {
		resultado = morse.MorseToText(datos.Morse)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"resultado": resultado,
	})
}
