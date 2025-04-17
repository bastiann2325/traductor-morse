package main

import (
	"net/http"
	"log"
	"html/template"
)

func main() {
	// Cargar templates
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))

	// Ruta para servir el formulario
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// Servir archivos estáticos (Css, Js, imágenes, etc.)
	http.Handle("/web/static/",http.StripPrefix("/web/static", http.FileServer(http.Dir("web/static"))))

	// Levantar servidor en el puerto 8080
	log.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}