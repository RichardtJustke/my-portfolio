package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRecord struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	CreatedAt    string `json:"created_at"`
}

type SessionRecord struct {
	Token     string `json:"token"`
	Username  string `json:"username"`
	ExpiresAt int64  `json:"expires_at"`
}

var (
	userMutex    sync.RWMutex
	usersPath    = filepath.Join("data", "users.json")
	sessionsPath = filepath.Join("data", "sessions.json")
)

func initAuthStore() {
	_ = os.MkdirAll("data", 0755)
	if _, err := os.Stat(usersPath); os.IsNotExist(err) {
		_ = os.WriteFile(usersPath, []byte("[]"), 0644)
	}
	if _, err := os.Stat(sessionsPath); os.IsNotExist(err) {
		_ = os.WriteFile(sessionsPath, []byte("[]"), 0644)
	}
}

func hashPassword(password string) string {
	salt := "newsite_secure_salt_2026"
	hash := sha256.Sum256([]byte(password + salt))
	return hex.EncodeToString(hash[:])
}

// RegisterHandler registers a new user with hashed credentials in MongoDB / JSON store
func RegisterHandler(c *fiber.Ctx) error {
	var body UserCredentials
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Payload inválido"})
	}

	body.Username = strings.TrimSpace(body.Username)
	if body.Username == "" || body.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Usuário e senha são obrigatórios"})
	}

	userMutex.Lock()
	defer userMutex.Unlock()

	initAuthStore()
	data, _ := os.ReadFile(usersPath)
	var users []UserRecord
	_ = json.Unmarshal(data, &users)

	for _, u := range users {
		if strings.EqualFold(u.Username, body.Username) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Usuário já existe no banco"})
		}
	}

	newUser := UserRecord{
		ID:           fmt.Sprintf("usr_%d", time.Now().UnixNano()),
		Username:     body.Username,
		PasswordHash: hashPassword(body.Password),
		CreatedAt:    time.Now().Format("02/01/2006 15:04"),
	}

	users = append(users, newUser)
	savedData, _ := json.MarshalIndent(users, "", "  ")
	_ = os.WriteFile(usersPath, savedData, 0644)

	// Create session token automatically
	token := createSession(body.Username)
	c.Cookie(&fiber.Cookie{
		Name:     "hub_session",
		Value:    token,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"success":  true,
		"message":  "Conta criada com sucesso!",
		"username": body.Username,
	})
}

// LoginHandler authenticates user against MongoDB / JSON store credentials
func LoginHandler(c *fiber.Ctx) error {
	var body UserCredentials
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Informe usuário e senha"})
	}

	username := strings.TrimSpace(body.Username)
	if username == "" || body.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Informe usuário e senha"})
	}

	userMutex.RLock()
	initAuthStore()
	data, _ := os.ReadFile(usersPath)
	var users []UserRecord
	_ = json.Unmarshal(data, &users)
	userMutex.RUnlock()

	targetHash := hashPassword(body.Password)
	var matchedUser *UserRecord

	for _, u := range users {
		if strings.EqualFold(u.Username, username) && u.PasswordHash == targetHash {
			matchedUser = &u
			break
		}
	}

	if matchedUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Credenciais incorretas"})
	}

	token := createSession(matchedUser.Username)
	c.Cookie(&fiber.Cookie{
		Name:     "hub_session",
		Value:    token,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"success":  true,
		"username": matchedUser.Username,
	})
}

// LogoutHandler clears authentication session cookie
func LogoutHandler(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "hub_session",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{"success": true})
}

// MeHandler returns currently authenticated user status
func MeHandler(c *fiber.Ctx) error {
	token := c.Cookies("hub_session")
	if token == "" {
		return c.JSON(fiber.Map{"authenticated": false})
	}

	username := validateSession(token)
	if username == "" {
		return c.JSON(fiber.Map{"authenticated": false})
	}

	return c.JSON(fiber.Map{
		"authenticated": true,
		"username":      username,
	})
}

// GetAuthenticatedUser returns username if session is active and valid, else empty string
func GetAuthenticatedUser(c *fiber.Ctx) string {
	token := c.Cookies("hub_session")
	if token == "" {
		return ""
	}
	return validateSession(token)
}

func createSession(username string) string {
	token := fmt.Sprintf("sess_%d_%s", time.Now().UnixNano(), username)
	session := SessionRecord{
		Token:     token,
		Username:  username,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	initAuthStore()
	data, _ := os.ReadFile(sessionsPath)
	var sessions []SessionRecord
	_ = json.Unmarshal(data, &sessions)
	sessions = append(sessions, session)

	savedData, _ := json.MarshalIndent(sessions, "", "  ")
	_ = os.WriteFile(sessionsPath, savedData, 0644)
	return token
}

func validateSession(token string) string {
	initAuthStore()
	data, err := os.ReadFile(sessionsPath)
	if err != nil {
		return ""
	}
	var sessions []SessionRecord
	_ = json.Unmarshal(data, &sessions)

	now := time.Now().Unix()
	for _, s := range sessions {
		if s.Token == token && s.ExpiresAt > now {
			return s.Username
		}
	}
	return ""
}
