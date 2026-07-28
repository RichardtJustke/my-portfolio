package handlers

import (
	"encoding/json"
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

var (
	notesMutex sync.RWMutex
	dataPath   = filepath.Join("data", "notes.json")
)

func initStore() {
	_ = os.MkdirAll("data", 0755)
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		initial := []NoteItem{
			{
				ID:        "tab-1",
				Title:     "Tarefas",
				Content:   "- [ ] Configurar atalhos CLI\n- [ ] Testar API do Clima\n- [x] Ajustar dock transparente",
				UpdatedAt: "hoje",
			},
			{
				ID:        "tab-2",
				Title:     "Comandos CLI",
				Content:   "go run main.go\ngit status\ncurl http://127.0.0.1:3000/clock",
				UpdatedAt: "hoje",
			},
		}
		data, _ := json.MarshalIndent(initial, "", "  ")
		_ = os.WriteFile(dataPath, data, 0644)
	}
}

// GetNotes returns all persistent scratchpad notes
func GetNotes(c *fiber.Ctx) error {
	notesMutex.RLock()
	defer notesMutex.RUnlock()

	initStore()
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return c.JSON([]NoteItem{})
	}

	var notes []NoteItem
	_ = json.Unmarshal(data, &notes)
	return c.JSON(notes)
}

// SaveNotes persists all scratchpad notes
func SaveNotes(c *fiber.Ctx) error {
	var notes []NoteItem
	if err := c.BodyParser(&notes); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	notesMutex.Lock()
	defer notesMutex.Unlock()

	initStore()
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if err := os.WriteFile(dataPath, data, 0644); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"status": "ok", "count": len(notes)})
}
