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
	"github.com/joho/godotenv"
)



func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	clientID := os.Getenv("SPOTIFY_ID")
	clientSecret := os.Getenv("SPOTIFY_SECRET")
	
	if clientID == "" || clientSecret == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Spotify credentials not found in environment.")

		if clientID == "" {
			fmt.Print("Enter Spotify Client ID: ")
			input, _ := reader.ReadString('\n')
			clientID = strings.TrimSpace(input)
		}

		if clientSecret == "" {
			fmt.Print("Enter Spotify Client Secret: ")
			input, _ := reader.ReadString('\n')
			clientSecret = strings.TrimSpace(input)
		}
	}

	if clientID == "" || clientSecret == "" {
		log.Fatal("please set SPOTIFY_ID and SPOTIFY_SECRET environment variables or enter them when prompted.")
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