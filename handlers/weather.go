package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type WeatherResponse struct {
	City        string  `json:"city"`
	Temp        float64 `json:"temp"`
	FeelsLike   float64 `json:"feels_like"`
	Description string  `json:"description"`
	Icon        string  `json:"icon"`
	Humidity    int     `json:"humidity"`
}

type openWeatherMapDTO struct {
	Name string `json:"name"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
}

// WeatherHandler proxies OpenWeatherMap API using City ID and API Key
func WeatherHandler(c *fiber.Ctx) error {
	cityID := c.Query("city_id", "3469058") // Default: Brasília
	apiKey := c.Query("api_key", "")

	if apiKey == "" {
		// Return friendly mock/cached response if API Key is not set yet
		return c.JSON(WeatherResponse{
			City:        "Brasília",
			Temp:        24.5,
			FeelsLike:   25.0,
			Description: "céu limpo",
			Icon:        "01d",
			Humidity:    60,
		})
	}

	targetURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?id=%s&appid=%s&units=metric&lang=pt_br", cityID, apiKey)

	client := &http.Client{Timeout: 6 * time.Second}
	resp, err := client.Get(targetURL)
	if err != nil || resp.StatusCode != 200 {
		// Fallback to Brasília defaults if API error or 401 unactivated key
		return c.JSON(WeatherResponse{
			City:        "Brasília",
			Temp:        24.0,
			FeelsLike:   24.5,
			Description: "parcialmente nublado",
			Icon:        "02d",
			Humidity:    58,
		})
	}
	defer resp.Body.Close()

	var dto openWeatherMapDTO
	if err := json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	desc := "parcialmente nublado"
	icon := "02d"
	if len(dto.Weather) > 0 {
		desc = dto.Weather[0].Description
		icon = dto.Weather[0].Icon
	}

	return c.JSON(WeatherResponse{
		City:        dto.Name,
		Temp:        dto.Main.Temp,
		FeelsLike:   dto.Main.FeelsLike,
		Description: desc,
		Icon:        icon,
		Humidity:    dto.Main.Humidity,
	})
}
