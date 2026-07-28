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
	CreatedAt string `json:"created_at"`
}

var (
	wallMutex sync.RWMutex
	wallPath  = filepath.Join("data", "wallpapers.json")
)

func initWallStore() {
	_ = os.MkdirAll("data", 0755)
	if _, err := os.Stat(wallPath); os.IsNotExist(err) {
		_ = os.WriteFile(wallPath, []byte("[]"), 0644)
	}
}

// GetWallpapers returns all saved wallpapers from database
func GetWallpapers(c *fiber.Ctx) error {
	wallMutex.RLock()
	defer wallMutex.RUnlock()

	initWallStore()
	data, err := os.ReadFile(wallPath)
	if err != nil {
		return c.JSON([]WallpaperItem{})
	}

	var walls []WallpaperItem
	_ = json.Unmarshal(data, &walls)
	return c.JSON(walls)
}

// SaveWallpaper saves a new uploaded wallpaper into the database
func SaveWallpaper(c *fiber.Ctx) error {
	var body struct {
		Name    string `json:"name"`
		DataURL string `json:"data_url"`
	}
	if err := c.BodyParser(&body); err != nil || body.DataURL == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid payload")
	}

	wallMutex.Lock()
	defer wallMutex.Unlock()

	initWallStore()
	data, err := os.ReadFile(wallPath)
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
		CreatedAt: time.Now().Format("02/01/2006 15:04"),
	}

	// Unshift so latest is first
	walls = append([]WallpaperItem{newItem}, walls...)

	savedData, _ := json.MarshalIndent(walls, "", "  ")
	_ = os.WriteFile(wallPath, savedData, 0644)

	return c.JSON(newItem)
}

// DeleteWallpaper removes a wallpaper by ID
func DeleteWallpaper(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("ID required")
	}

	wallMutex.Lock()
	defer wallMutex.Unlock()

	initWallStore()
	data, err := os.ReadFile(wallPath)
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
	_ = os.WriteFile(wallPath, savedData, 0644)

	return c.JSON(fiber.Map{"status": "deleted", "id": id})
}
