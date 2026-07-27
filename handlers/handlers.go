package handlers

import (
	"bytes"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

var tmpl *template.Template

func SetTemplates(t *template.Template) {
	tmpl = t
}

// renderTemplate executa o template "base" com os dados e envia como HTML.
func renderTemplate(c *fiber.Ctx, data interface{}) error {
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "base", data); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.Send(buf.Bytes())
}