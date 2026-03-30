package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Essa linha é o que tava faltando — serve os arquivos da pasta static/
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(
			"templates/base.html",
			"templates/home.html",
		))
		tmpl.ExecuteTemplate(w, "base", nil)
	})

	r.Get("/writing", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("writing — em breve"))
	})

	log.Println("rodando em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
