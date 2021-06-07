package main

import "testing"

func Test_youtubeSearch(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"John Askew - Chime", "John Askew - Chime", "https://youtu.be/aDlZckOUHiw"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := youtubeSearch(tt.args); got != tt.want {
				t.Errorf("youtubeSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
