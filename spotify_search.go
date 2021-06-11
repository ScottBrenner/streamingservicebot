package main

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify"
)

func spotifySearch(postTitle string) string {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SSB_SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SSB_SPOTIFY_CLIENT_SECRET"),
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)
	results, err := client.Search(postTitle, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}

	if results.Tracks != nil {
		return results.Tracks.Tracks[0].ExternalURLs["spotify"]
	}
	return ""
}
