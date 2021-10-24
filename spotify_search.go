package main

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

func spotifySearch(postTitle string) string {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SSB_SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SSB_SPOTIFY_CLIENT_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
	results, err := client.Search(ctx, postTitle, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}

	if results.Tracks != nil {
		return results.Tracks.Tracks[0].ExternalURLs["spotify"]
	}
	return ""
}
