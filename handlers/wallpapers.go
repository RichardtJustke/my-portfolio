package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type WallpaperItem struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	DataURL   string `json:"data_url"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}

var wallMutex sync.RWMutex

func getUserWallPath(username string) string {
	_ = os.MkdirAll("data", 0755)
	if username == "" {
		return filepath.Join("data", "wallpapers_public.json")
	}
	return filepath.Join("data", fmt.Sprintf("wallpapers_%s.json", username))
}

// GetWallpapers returns saved wallpapers for the current authenticated user
func GetWallpapers(c *fiber.Ctx) error {
	username := GetAuthenticatedUser(c)

	wallMutex.RLock()
	defer wallMutex.RUnlock()

	userPath := getUserWallPath(username)
	data, err := os.ReadFile(userPath)
	if err != nil || len(data) == 0 {
		return c.JSON([]WallpaperItem{})
	}

	var walls []WallpaperItem
	_ = json.Unmarshal(data, &walls)
	return c.JSON(walls)
}

// SaveWallpaper saves a new uploaded wallpaper into the database ONLY for logged-in users
func SaveWallpaper(c *fiber.Ctx) error {
	username := GetAuthenticatedUser(c)
	if username == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "🔒 Faça login para salvar wallpapers no banco de dados!",
		})
	}

	var body struct {
		Name    string `json:"name"`
		DataURL string `json:"data_url"`
	}
	if err := c.BodyParser(&body); err != nil || body.DataURL == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid payload")
	}

	wallMutex.Lock()
	defer wallMutex.Unlock()

	userPath := getUserWallPath(username)
	data, err := os.ReadFile(userPath)
	var walls []WallpaperItem
	if err == nil {
		_ = json.Unmarshal(data, &walls)
	}

	name := body.Name
	if name == "" {
		name = fmt.Sprintf("Wallpaper %d", len(walls)+1)
	}

	newItem := WallpaperItem{
		ID:        fmt.Sprintf("wall_%d", time.Now().UnixNano()),
		Name:      name,
		DataURL:   body.DataURL,
		Username:  username,
		CreatedAt: time.Now().Format("02/01/2006 15:04"),
	}

	walls = append([]WallpaperItem{newItem}, walls...)

	savedData, _ := json.MarshalIndent(walls, "", "  ")
	_ = os.WriteFile(userPath, savedData, 0644)

	return c.JSON(newItem)
}

// DeleteWallpaper removes a wallpaper by ID for the logged-in user
func DeleteWallpaper(c *fiber.Ctx) error {
	username := GetAuthenticatedUser(c)
	if username == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("🔒 Faça login necessário")
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("ID required")
	}

	wallMutex.Lock()
	defer wallMutex.Unlock()

	userPath := getUserWallPath(username)
	data, err := os.ReadFile(userPath)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Not found")
	}

	var walls []WallpaperItem
	_ = json.Unmarshal(data, &walls)

	var filtered []WallpaperItem
	for _, w := range walls {
		if w.ID != id {
			filtered = append(filtered, w)
		}
	}

	savedData, _ := json.MarshalIndent(filtered, "", "  ")
	_ = os.WriteFile(userPath, savedData, 0644)

	return c.JSON(fiber.Map{"status": "deleted", "id": id})
}
