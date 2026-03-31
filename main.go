package main

import (
    "embed"
    "html/template"
    "io/fs"
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"

    "github.com/RichardtJustke/my-portfolio/handlers"
)

//go:embed templates static
var files embed.FS

func main() {
    // Extrai os arquivos estáticos do embed
    staticFiles, err := fs.Sub(files, "static")
    if err != nil {
        log.Fatal(err)
    }

    // Passa os templates embedados pros handlers
    tmpl := template.Must(template.ParseFS(files,
        "templates/base.html",
        "templates/home.html",
    ))
    handlers.SetTemplates(tmpl)

    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))
    r.Get("/", handlers.Home)
    r.Get("/writing", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("writing — em breve"))
    })

    log.Println("rodando em http://localhost:3000")
    log.Fatal(http.ListenAndServe(":3000", r))
}
