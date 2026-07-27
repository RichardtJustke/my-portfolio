package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type HomeData struct {
	Name     string
	Role     string
	Bio      string
	Location string
	Email    string
	GitHub   string
	LinkedIn string
	Page     string
}

func Home(c *fiber.Ctx) error {
	data := HomeData{
		Name:     "Richard Justke",
		Role:     "BACKEND ENGINEER",
		Bio:      "I build systems and tools. Passionate about Go, infrastructure, and simple code that solves real problems.",
		Location: "Brasília, BR",
		Email:    "rj.justke@gmail.com",
		GitHub:   "https://github.com/RichardtJustke",
		LinkedIn: "https://linkedin.com/in/richardtjustke",
		Page:     "home",
	}
	return renderTemplate(c, data)
}
