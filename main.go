package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/RichardtJustke/my-portfolio-v2/handlers"
)

//go:embed templates static assets
var files embed.FS

func main() {
	staticFiles, err := fs.Sub(files, "static")
	if err != nil {
		log.Fatal(err)
	}

	assetsFiles, err := fs.Sub(files, "assets")
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"contains": strings.Contains,
	}).ParseFS(files,
		"templates/base.html",
		"templates/home.html",
		"templates/work.html",
		"templates/resume.html",
		"templates/clock.html",
	))
	handlers.SetTemplates(tmpl)

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	app.Use(logger.New())

	app.Get("/static/*", func(c *fiber.Ctx) error {
		path := strings.TrimPrefix(c.Params("*"), "/")
		if path == "" {
			return c.SendStatus(fiber.StatusNotFound)
		}
		f, err := staticFiles.Open(path)
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		defer f.Close()
		info, err := f.Stat()
		if err != nil || info.IsDir() {
			return c.SendStatus(fiber.StatusNotFound)
		}

		// Define Content-Type baseado na extensão do arquivo
		ext := filepath.Ext(path)
		if mimeType := mime.TypeByExtension(ext); mimeType != "" {
			c.Set("Content-Type", mimeType)
		}

		_, err = io.Copy(c.Response().BodyWriter(), f)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return nil
	})

	app.Get("/assets/*", func(c *fiber.Ctx) error {
		path := strings.TrimPrefix(c.Params("*"), "/")
		if path == "" {
			return c.SendStatus(fiber.StatusNotFound)
		}
		f, err := assetsFiles.Open(path)
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		defer f.Close()

		// Define Content-Type baseado na extensão do arquivo
		ext := filepath.Ext(path)
		if mimeType := mime.TypeByExtension(ext); mimeType != "" {
			c.Set("Content-Type", mimeType)
		}

		_, err = io.Copy(c.Response().BodyWriter(), f)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return nil
	})

	app.Get("/sw.js", func(c *fiber.Ctx) error {
		content, err := fs.ReadFile(files, "static/sw.js")
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		c.Set("Content-Type", "application/javascript")
		c.Set("Service-Worker-Allowed", "/")
		return c.Send(content)
	})

	app.Get("/manifest.json", func(c *fiber.Ctx) error {
		content, err := fs.ReadFile(files, "static/manifest.json")
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		c.Set("Content-Type", "application/json")
		return c.Send(content)
	})

	app.Get("/", handlers.Home)
	app.Get("/work", handlers.Work)
	app.Get("/resume", handlers.Resume)

	app.Get("/writing", func(c *fiber.Ctx) error {
		return c.SendString("writing — em breve")
	})

	app.Get("/clock", handlers.Clock)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("rodando em http://%s\n", addr)
	log.Fatal(app.Listen(addr))
}