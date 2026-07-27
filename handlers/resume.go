package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type ResumeData struct {
	Page string
}

func Resume(c *fiber.Ctx) error {
	data := ResumeData{
		Page: "resume",
	}
	return renderTemplate(c, data)
}