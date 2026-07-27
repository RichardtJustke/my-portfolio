package handlers

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
)

func Clock(c *fiber.Ctx) error {
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "clock", nil); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.Send(buf.Bytes())
}
