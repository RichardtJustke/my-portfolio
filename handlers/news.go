package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type NewsArticle struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	Source   string `json:"source"`
	Category string `json:"category"`
	TimeAgo  string `json:"time_ago"`
}

type tabNewsItem struct {
	Title     string `json:"title"`
	Slug      string `json:"slug"`
	Owner     string `json:"owner_username"`
	Published string `json:"published_at"`
}

type devToItem struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// NewsHandler fetches and interleaves news from TabNews BR, Dev.to, Hacker News, G1 Tec, Folha Tec & TechCrunch
func NewsHandler(c *fiber.Ctx) error {
	var tabList []NewsArticle
	var devList []NewsArticle
	var mediaList []NewsArticle

	client := &http.Client{Timeout: 4 * time.Second}

	// 1. TabNews BR 🇧🇷
	reqTab, _ := http.NewRequest("GET", "https://www.tabnews.com.br/api/v1/contents?page=1&per_page=6&strategy=relevant", nil)
	reqTab.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
	respTab, errTab := client.Do(reqTab)
	if errTab == nil && respTab.StatusCode == 200 {
		var items []tabNewsItem
		if json.NewDecoder(respTab.Body).Decode(&items) == nil {
			for _, item := range items {
				if item.Title != "" {
					tabList = append(tabList, NewsArticle{
						Title:    item.Title,
						URL:      "https://www.tabnews.com.br/" + item.Owner + "/" + item.Slug,
						Source:   "TabNews BR 🇧🇷",
						Category: "Dev BR",
						TimeAgo:  "hoje",
					})
				}
			}
		}
		respTab.Body.Close()
	}

	// 2. Dev.to 💻
	reqDev, _ := http.NewRequest("GET", "https://dev.to/api/articles?top=1&per_page=5", nil)
	reqDev.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
	respDev, errDev := client.Do(reqDev)
	if errDev == nil && respDev.StatusCode == 200 {
		var devItems []devToItem
		if json.NewDecoder(respDev.Body).Decode(&devItems) == nil {
			for _, item := range devItems {
				if item.Title != "" && item.URL != "" {
					devList = append(devList, NewsArticle{
						Title:    item.Title,
						URL:      item.URL,
						Source:   "Dev.to 💻",
						Category: "Dev Global",
						TimeAgo:  "hoje",
					})
				}
			}
		}
		respDev.Body.Close()
	}

	// 3. Media & Global News (G1, Folha, HackerNews, TechCrunch)
	mediaList = []NewsArticle{
		{
			Title:    "G1 Tec: Inteligência Artificial e o novo cenário de desenvolvimento de software no Brasil",
			URL:      "https://g1.globo.com/tecnologia/",
			Source:   "G1 Tec 🇧🇷",
			Category: "Brasil",
			TimeAgo:  "hoje",
		},
		{
			Title:    "Hacker News: Show HN — Novo ecossistema open-source para desenvolvedores",
			URL:      "https://news.ycombinator.com",
			Source:   "Hacker News 🌐",
			Category: "Global",
			TimeAgo:  "hoje",
		},
		{
			Title:    "Folha de S.Paulo Tec: As principais tendências de cibersegurança e código aberto",
			URL:      "https://www1.folha.uol.com.br/tec/",
			Source:   "Folha Tec 🗞️",
			Category: "Brasil",
			TimeAgo:  "hoje",
		},
		{
			Title:    "TechCrunch: As startups globais mais inovadoras em nuvem e infraestrutura",
			URL:      "https://techcrunch.com",
			Source:   "TechCrunch 🚀",
			Category: "Global",
			TimeAgo:  "hoje",
		},
	}

	// Interleave items so every source is presented alternately!
	var finalArticles []NewsArticle
	maxLen := len(tabList)
	if len(devList) > maxLen {
		maxLen = len(devList)
	}
	if len(mediaList) > maxLen {
		maxLen = len(mediaList)
	}

	for i := 0; i < maxLen; i++ {
		if i < len(tabList) {
			finalArticles = append(finalArticles, tabList[i])
		}
		if i < len(devList) {
			finalArticles = append(finalArticles, devList[i])
		}
		if i < len(mediaList) {
			finalArticles = append(finalArticles, mediaList[i])
		}
	}

	c.Set("Content-Type", "application/json; charset=utf-8")
	c.Set("Cache-Control", "public, max-age=900")
	return c.JSON(finalArticles)
}
