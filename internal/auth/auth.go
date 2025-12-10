package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

func Authenticate() (*spotify.Client, error) {
	auth := spotifyauth.New(
		spotifyauth.WithRedirectURL(os.Getenv("http://127.0.0.1:8080/callback")),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadCurrentPlaying,
			spotifyauth.ScopeUserReadPlaybackState,
			spotifyauth.ScopeUserModifyPlaybackState,
		),
		spotifyauth.WithClientID(os.Getenv("b8f09d32692143458b0907def109b87f"))
		spotifyauth.WithClientSecret(os.Getenv("a36b9b3815f44333830161f31863ed9e"))
	)
	
	ch := make(chan *spotify.Client)
	completeAuth := func(w http.ResponseWriter, r *http.Request) {
		tok, err := auth.Token(r.Context(), auth.State(), r)
		if err != nil {
			http.Error(w, "Couldn't get token", http.StatusForbidden)
			fmt.Println("Couldn't get token", err)
			return
		}

		client := spotify.New(auth.Client(r.Context(), tok))

		fmt.Fprintf(w,"Login Completed! You can close this tab and return to SudoSpot.")

		ch <- client
	}

	http.HandleFunc("/callback", completeAuth)

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("Server error", err)
		}
	}()

	url := auth.AuthURL("state-string")
	fmt.Println("Please log in to Spotify by visiting the following page in your browser")
	fmt.Printf("\n%s\n\n", url)

	client := <-ch
	return client, nil
}

func saveToken(tok *oauth2.Token) {
	// TODO: Save to a JSON file so we don't have to log in every time
}