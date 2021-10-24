package main

import "testing"

func Test_spotifySearch(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"John Askew - Chime", "John Askew - Chime", "https://open.spotify.com/track/6xWwbqSsos0LFDVrx134HC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spotifySearch(tt.args); got != tt.want {
				t.Errorf("spotifySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
