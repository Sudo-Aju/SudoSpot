package main

import(
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"


	"github.com/Sudo-Aju/sudospot/internal/auth"
	"github.com/Sudo-Aju/sudospot/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"encoding/json"
	"path/filepath"
)



func main() {

	// Define config structure
	type Config struct {
		ClientID     string `json:"spotify_id"`
		ClientSecret string `json:"spotify_secret"`
	}

	// Get user config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Could not determine user config directory:", err)
	}
	appConfigDir := filepath.Join(configDir, "sudospot")
	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		log.Fatal("Could not create config directory:", err)
	}

	configFile := filepath.Join(appConfigDir, "config.json")
	tokenPath := filepath.Join(appConfigDir, "token.json")

	// Load existing config
	var config Config
	if file, err := os.Open(configFile); err == nil {
		json.NewDecoder(file).Decode(&config)
		file.Close()
	}

	// Fallback to env vars
	if config.ClientID == "" {
		config.ClientID = os.Getenv("SPOTIFY_ID")
	}
	if config.ClientSecret == "" {
		config.ClientSecret = os.Getenv("SPOTIFY_SECRET")
	}

	// Prompt if still missing
	if config.ClientID == "" || config.ClientSecret == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Spotify credentials not found.")

		if config.ClientID == "" {
			fmt.Print("Enter Spotify Client ID: ")
			input, _ := reader.ReadString('\n')
			config.ClientID = strings.TrimSpace(input)
		}

		if config.ClientSecret == "" {
			fmt.Print("Enter Spotify Client Secret: ")
			input, _ := reader.ReadString('\n')
			config.ClientSecret = strings.TrimSpace(input)
		}

		// Save to config file
		file, err := os.Create(configFile)
		if err != nil {
			log.Printf("Warning: failed to save config: %v", err)
		} else {
			json.NewEncoder(file).Encode(config)
			file.Close()
			fmt.Printf("Credentials saved to %s\n", configFile)
		}
	}

	if config.ClientID == "" || config.ClientSecret == "" {
		log.Fatal("please set SPOTIFY_ID and SPOTIFY_SECRET environment variables or enter them when prompted.")
	}

	client, err := auth.Authenticate(config.ClientID, config.ClientSecret, tokenPath)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	p := tea.NewProgram(ui.NewModel(client))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}