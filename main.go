package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/RichardtJustke/my-portfolio/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Get("/", handlers.Home)

	r.Get("/writing", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("writing — em breve"))
	})

	log.Println("rodando em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
