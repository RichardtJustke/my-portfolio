package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/richardtjustke/site/handlers"
)

func main() {
	r := chi.NewRouter()

	// Middleware: loga cada requisição e recupera de panics
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Arquivos estáticos (CSS, JS)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Rotas
	r.Get("/", handlers.Home)
	r.Get("/writing", handlers.WritingList)
	r.Get("/writing/{slug}", handlers.WritingPost)

	log.Println("servidor rodando em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
