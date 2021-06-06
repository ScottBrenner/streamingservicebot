package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

const maxResults = 5
const urlPrefix = "https://youtu.be/%s"

func youtubeSearch(postTitle string) string {
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("SSB_YOUTUBE_KEY")},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List([]string{"id,snippet"}).
		Q(postTitle).
		MaxResults(maxResults)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}

	var result string

	// Iterate through each item and find auto-generated upload.
	for _, item := range response.Items {
		if strings.Contains(item.Snippet.Description, "Auto-generated") {
			result = fmt.Sprintf(urlPrefix, item.Id.VideoId)
		}
	}
	return result
}
