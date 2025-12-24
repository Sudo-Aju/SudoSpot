package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

const (
	redirectURL = "http://127.0.0.1:8080/callback"
	tokenFile = "token.json"
)


func Authenticate(clientID, clientSecret, tokenPath string) (*spotify.Client, error) {
	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("client ID and Secret are required")
	}

	auth := spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURL),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadCurrentlyPlaying,
			spotifyauth.ScopeUserReadPlaybackState,
			spotifyauth.ScopeUserModifyPlaybackState,
		),
		spotifyauth.WithClientID(clientID),
		spotifyauth.WithClientSecret(clientSecret),
	)
	
	if token, err := loadToken(tokenPath); err == nil {
		client := spotify.New(auth.Client(context.Background(), token))
		_, err := client.PlayerCurrentlyPlaying(context.Background())
		if err == nil {
			return client, nil
		}
		fmt.Println("Saved token expired, re-authenticating...")
	}

	ch := make(chan *spotify.Client)
	state := "my-cli-app-state"
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request){
		token, err := auth.Token(r.Context(), state, r)
		if err != nil {
			http.Error(w, "couldn't get token", http.StatusForbidden)
			log.Printf("Error getting token: %v", err)
			return
		}
		saveToken(token, tokenPath)

		client := spotify.New(auth.Client(r.Context(), token))
		fmt.Fprintf(w, "login Completed! You can close this window and return to the SudoSpot")
		ch <-client
	})

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:")
	fmt.Printf("\n%s\n\n", url)

	client := <-ch
	_ = server.Shutdown(context.Background())

	return client, nil

}

func saveToken(tok *oauth2.Token, path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Printf("Warning: failed to save token: %v", err)
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(tok)
}

func loadToken(path string) (*oauth2.Token, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tok := &oauth2.Token{}
	if err := json.NewDecoder(file).Decode(tok); err != nil {
		return nil, err
	}
	return tok, nil
}