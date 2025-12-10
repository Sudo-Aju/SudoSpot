package main

import(
	"fmt"
	"log"
	"os"

	"github.com/Sudo-Aju/sudospot/internal/auth"
	"github.com/Sudo-Aju/sudospot/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)



func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	clientID := os.Getenv("SPOTIFY_ID")
	clientSecret := os.Getenv("SPOTIFY_SECRET")
	
	if clientID == "" || clientSecret == "" {
		log.Fatal("please set SPOTIFY_ID and SPOTIFY_SECRET environment variables.")
	}

	client, err := auth.Authenticate(clientID, clientSecret)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	p := tea.NewProgram(ui.NewModel(client))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}