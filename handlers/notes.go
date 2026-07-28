package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type NoteItem struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UpdatedAt string `json:"updated_at"`
}

var notesMutex sync.RWMutex

func getUserNotesPath(username string) string {
	_ = os.MkdirAll("data", 0755)
	if username == "" {
		return filepath.Join("data", "notes_guest.json")
	}
	return filepath.Join("data", fmt.Sprintf("notes_%s.json", username))
}

// GetNotes returns persistent scratchpad notes for the authenticated user
func GetNotes(c *fiber.Ctx) error {
	username := GetAuthenticatedUser(c)

	notesMutex.RLock()
	defer notesMutex.RUnlock()

	userPath := getUserNotesPath(username)
	data, err := os.ReadFile(userPath)
	if err != nil || len(data) == 0 {
		initial := []NoteItem{
			{
				ID:        "tab-1",
				Title:     "Tarefas",
				Content:   "# Minhas Tarefas\n- [ ] Configurar atalhos CLI\n- [ ] Testar API do Clima\n- [x] Ajustar dock transparente",
				UpdatedAt: "hoje",
			},
			{
				ID:        "tab-2",
				Title:     "Comandos CLI",
				Content:   "```bash\ngo run main.go\ngit status\ncurl http://127.0.0.1:3000/clock\n```",
				UpdatedAt: "hoje",
			},
		}
		return c.JSON(initial)
	}

	var notes []NoteItem
	_ = json.Unmarshal(data, &notes)
	return c.JSON(notes)
}

// SaveNotes persists scratchpad notes ONLY for authenticated users
func SaveNotes(c *fiber.Ctx) error {
	username := GetAuthenticatedUser(c)
	if username == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "🔒 Faça login para salvar suas notas no banco de dados!",
		})
	}

	var notes []NoteItem
	if err := c.BodyParser(&notes); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	notesMutex.Lock()
	defer notesMutex.Unlock()

	userPath := getUserNotesPath(username)
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if err := os.WriteFile(userPath, data, 0644); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"status": "ok", "count": len(notes), "user": username})
}
