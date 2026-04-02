package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/RichardtJustke/my-portfolio/handlers"
)

//go:embed templates static assets
var files embed.FS

func main() {
	// Extrai os arquivos estáticos do embed
	staticFiles, err := fs.Sub(files, "static")
	if err != nil {
		log.Fatal(err)
	}

	// Extrai os arquivos de assets
	assetsFiles, err := fs.Sub(files, "assets")
	if err != nil {
		log.Fatal(err)
	}

	// Passa os templates embedados pros handlers
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"contains": strings.Contains,
	}).ParseFS(files,
		"templates/base.html",
		"templates/home.html",
		"templates/work.html",
	))
	handlers.SetTemplates(tmpl)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(assetsFiles))))
	r.Get("/", handlers.Home)
	r.Get("/work", handlers.Work)
	r.Get("/writing", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("writing — em breve"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("rodando em http://0.0.0.0:%s\n", port)
	log.Fatal(http.ListenAndServe(addr, r))
}
