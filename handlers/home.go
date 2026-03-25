package handlers

import (
	"html/template"
	"net/http"
)

// HomeData é o que passamos pro template da home
type HomeData struct {
	Name     string
	Role     string
	Bio      string
	Location string
	Email    string
	GitHub   string
	LinkedIn string
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := HomeData{
		Name:     "Richard Justke",
		Role:     "Backend Engineer",
		Bio:      "Construo sistemas e ferramentas. Apaixonado por Go, infraestrutura e código simples que resolve problemas reais.",
		Location: "São Paulo, BR",
		Email:    "hi@richardtjustke.dev",
		GitHub:   "https://github.com/richardtjustke",
		LinkedIn: "https://linkedin.com/in/richardtjustke",
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/home.html",
	))

	tmpl.ExecuteTemplate(w, "base", data)
}
