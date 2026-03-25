package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/yuin/goldmark"
)

// Post representa um artigo do blog
type Post struct {
	Title   string
	Slug    string
	Date    time.Time
	Content template.HTML // HTML já renderizado do markdown
}

// WritingList lista todos os posts
func WritingList(w http.ResponseWriter, r *http.Request) {
	posts, err := loadAllPosts()
	if err != nil {
		http.Error(w, "erro ao carregar posts", http.StatusInternalServerError)
		return
	}

	// Ordena do mais recente pro mais antigo
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/writing.html",
	))

	tmpl.ExecuteTemplate(w, "base", posts)
}

// WritingPost exibe um post individual pelo slug
func WritingPost(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	post, err := loadPost(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/post.html",
	))

	tmpl.ExecuteTemplate(w, "base", post)
}

// loadPost lê um arquivo markdown e converte pra HTML
func loadPost(slug string) (*Post, error) {
	path := filepath.Join("content", "writing", slug+".md")

	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Converte markdown → HTML
	var buf strings.Builder
	if err := goldmark.Convert(raw, &buf); err != nil {
		return nil, err
	}

	return &Post{
		Title:   slug, // por enquanto usa o slug; depois vamos adicionar frontmatter
		Slug:    slug,
		Date:    time.Now(), // placeholder
		Content: template.HTML(buf.String()),
	}, nil
}

// loadAllPosts carrega todos os .md da pasta content/writing
func loadAllPosts() ([]Post, error) {
	files, err := filepath.Glob("content/writing/*.md")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range files {
		slug := strings.TrimSuffix(filepath.Base(f), ".md")
		post, err := loadPost(slug)
		if err != nil {
			continue
		}
		posts = append(posts, *post)
	}

	return posts, nil
}
