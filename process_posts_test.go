package main

import (
	"fmt"
	"testing"

	"github.com/turnage/graw/reddit"
)

func Test_checkURL(t *testing.T) {
	tests := []struct {
		name string
		args *reddit.Post
		want bool
	}{
		{"youtube.com", &reddit.Post{URL: "https://www.youtube.com/watch?v=aDlZckOUHiw"}, true},
		{"reddit.com", &reddit.Post{URL: "https://www.reddit.com/user/streamingservicebot"}, false},
		{"open.spotify.com", &reddit.Post{URL: "https://open.spotify.com/track/4JXmVLQ6cYfakyjRvIAWDW"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkURL(tt.args); got != tt.want {
				t.Errorf("checkURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processStreamingService(t *testing.T) {
	tests := []struct {
		name    string
		service string
		args    *reddit.Post
		want    string
	}{
		{"youtu.be", "YouTube", &reddit.Post{Title: "John Askew - Chime"}, fmt.Sprintf("- [YouTube](%s)\n", "https://youtu.be/aDlZckOUHiw")},
		{"open.spotify.com", "Spotify", &reddit.Post{Title: "John Askew - Chime"}, fmt.Sprintf("- [Spotify](%s)\n", "https://open.spotify.com/track/4JXmVLQ6cYfakyjRvIAWDW")},
		{"None", "None", &reddit.Post{Title: "John Askew - Chime"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processStreamingService(tt.service, tt.args); got != tt.want {
				t.Errorf("processStreamingService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateReply(t *testing.T) {
	tests := []struct {
		name string
		args *reddit.Post
		want string
	}{
		{"All", &reddit.Post{Title: "John Askew - Chime"}, `I found this track on other streaming services!:

- [YouTube](https://youtu.be/aDlZckOUHiw)
- [Spotify](https://open.spotify.com/track/4JXmVLQ6cYfakyjRvIAWDW)


----

^(I am a bot, and this action was performed automatically.) [^(Source)](https://github.com/ScottBrenner/streamingservicebot)`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateReply(tt.args); got != tt.want {
				t.Errorf("generateReply() = %v, want %v", got, tt.want)
			}
		})
	}
}
